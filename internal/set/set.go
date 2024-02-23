package set

import (
	"github.com/youta-t/its/internal/diff"
	"github.com/youta-t/its/internal/list"
	"github.com/youta-t/its/itskit"
)

func CompareWithMatcher[T any](values []T, specs []itskit.Matcher[T]) []diff.Diff[itskit.Match] {
	return Compare[T, itskit.Matcher[T], itskit.Match](
		values, specs,
		func(t T, m itskit.Matcher[T]) (itskit.Match, bool) {
			match := m.Match(t)
			return match, match.Ok()
		},
		diff.ExtraMatch[T],
		diff.MissingMatch[T],
	)
}

func Compare[A, B, C any](
	values []A, specs []B,
	pred func(A, B) (C, bool),
	toExtra func(A) C,
	toMissing func(B) C,
) []diff.Diff[C] {
	// Spec may match 0 or more items in values, thus can match with many.
	// Even if in such case, we would like to do better pairing,
	// so, we need "maximum matching in a bipartite graph".
	//
	// To solve that, we use Ford-Fulkerson method.

	// step 1. make a adjacency matrix, as flowing from value to spec.
	//
	// We only think connections between value-spec and spec-sink.
	//
	// There are no value->value edge or spec->spec edge,
	// since it is a bipartite graph.
	// Thus, we should manage value->spec part and spec-> value part.
	//
	// There are no needs to think about edges between
	// source and value or spec and sink.
	// Because, in augmenting *path*, we can traverse from source/to sink only once.
	// So, value->source or sink->spec edge cannot be used.

	// value -> spec matrix:
	//
	// This graph is bipartite,
	// and there are no edges like value -> value or spec -> spec,
	// our adjcency matrices are [len(value)][len(spec)]int.
	edges := map[[2]int]C{}
	value_spec := make2D[int](len(values), len(specs))
	for i_value := range values {
		v := values[i_value]
		for i_spec := range specs {
			s := specs[i_spec]
			c, ok := pred(v, s)
			if !ok {
				continue
			}
			value_spec[i_value][i_spec] = 1
			edges[[2]int{i_value, i_spec}] = c
		}
	}

	// spec -> value matrix, spec & sink matrix:
	//
	// To make residual graph, prepare reversed edges for value -> spec.
	//
	// At the beggining, all edges are zero.
	spec_value := make2D[int](len(specs), len(values))

	// if there are key, the spec have capacity outgoing to the sink.
	spec_sink := map[int]struct{}{}
	for i := range specs {
		spec_sink[i] = struct{}{}
	}

	// now, we have adjcency matrix of the residual network.

	// step 2. walk augmenting path.
	//
	// Each edge's capacity is just 1, so augmenting pathes can be upto len(values).
	// Testing for each value node and try to reach sink node.
VALUE:
	for i_value := range value_spec {
		i_spec := -1

		// recording path. the path is:
		//   v(alue)_on_path[0]->s(pec)_on_path[0]->v_on_path[1]->s_on_path[1]-> ...
		v_on_path := list.New([]int{i_value})
		s_on_path := list.New([]int{})

		v_visit := map[int]struct{}{i_value: {}}
		s_visit := map[int]struct{}{}

		for {
			// /// graph traverser is at value node i_value ///
			next_i_spec := -1
			for i_s, capacity := range value_spec[i_value] {
				// find not visited node with outgoing capacity.
				if capacity <= 0 {
					continue
				}
				if _, ok := s_visit[i_s]; ok {
					continue
				}
				next_i_spec = i_s
				break
			}

			if next_i_spec < 0 {
				i_value, _ = v_on_path.PopRight() // backtrack

				// undo the last spec->value move
				if s_last, ok := s_on_path.Nth(-1); !ok {
					continue VALUE // no backtrack! No augmenting path from this beggining value node
				} else {
					i_spec = s_last
				}
			} else {
				i_spec = next_i_spec

				// move to next spec node
				s_on_path.Append(i_spec)
				s_visit[i_spec] = struct{}{}
			}
			value_spec[i_value][i_spec] -= 1
			spec_value[i_spec][i_value] += 1

			// /// graph traverser is at spec node i_spec ///
			if _, ok := spec_sink[i_spec]; ok {
				// augmenting path is found!
				delete(spec_sink, i_spec)
				continue VALUE
			}

			next_i_value := -1
			for i_v, capacity := range spec_value[i_spec] {
				if capacity <= 0 {
					continue
				}
				if _, ok := v_visit[i_v]; ok {
					continue
				}

				// traverse to value
				next_i_value = i_v
				v_on_path.Append(i_value)
				break
			}
			if next_i_value <= 0 {
				i_spec, _ = s_on_path.PopRight() // backtrack

				// undo the last value->spec move
				if v_last, ok := v_on_path.Nth(-1); !ok {
					continue VALUE
				} else {
					i_value = v_last
				}
			} else {
				i_value = next_i_value
				v_on_path.Append(i_value)
				v_visit[i_value] = struct{}{}
			}
			spec_value[i_spec][i_value] -= 1
			value_spec[i_value][i_spec] += 1
		}
	}

	// step3. we know maximum matching,
	// Matched pair are edge with 0 < capacity in spec->value.
	v_node := map[int]struct{}{}
	for i_value := range values {
		v_node[i_value] = struct{}{}
	}
	s_node := map[int]struct{}{}
	for i_spec := range specs {
		s_node[i_spec] = struct{}{}
	}

	ret := []diff.Diff[C]{}
	// matched nodes
	for i_spec := range spec_value {
		for i_value, capacity := range spec_value[i_spec] {
			if 0 < capacity {
				ret = append(ret, diff.OkItem(edges[[2]int{i_value, i_spec}]))
				delete(v_node, i_value)
				delete(s_node, i_spec)
				break
			}
		}
	}
	// unmatched values
	for i_value := range v_node {
		ret = append(ret, diff.ExtraItem(toExtra(values[i_value])))
	}
	for i_spec := range s_node {
		ret = append(ret, diff.MissingItem(toMissing(specs[i_spec])))
	}
	return ret
}

func make2D[T any](x int, y int) [][]T {
	ret := make([][]T, x)
	for i := range ret {
		ret[i] = make([]T, y)
	}
	return ret
}
