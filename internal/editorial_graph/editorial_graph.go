package editorialgraph

import (
	"github.com/youta-t/its/internal/diff"
	"github.com/youta-t/its/itskit"
)

type position struct {
	spec  int
	value int
}

type diffTrace struct {
	p     position
	trace []diff.Diff
}

func New[A any](ss []itskit.Matcher[A], as []A) []diff.Diff {
	// based Myers Algorithm.

	head := []diffTrace{
		{
			p:     position{spec: -1, value: -1},
			trace: []diff.Diff{},
		},
	}
	visit := map[position]struct{}{}
	centinel := position{spec: len(ss) - 1, value: len(as) - 1}

	for {
		{
			newHead := []diffTrace{}
			for _, h := range head[:] {
				for h.p.spec < centinel.spec && h.p.value < centinel.value {
					p := position{spec: h.p.spec + 1, value: h.p.value + 1}
					if _, ok := visit[p]; ok {
						goto SKIP
					}

					m := ss[p.spec].Match(as[p.value])
					if !m.Ok() {
						break
					}
					visit[p] = struct{}{}
					h = diffTrace{p: p, trace: append(
						sliceclone(h.trace), diff.OkItem(m),
					)}
					if h.p == centinel {
						return h.trace
					}
				}
				newHead = append(newHead, h)
			SKIP:
			}
			head = newHead
		}

		newHead := []diffTrace{}
		for _, h := range head {
			if h.p.spec < centinel.spec {
				p := position{spec: h.p.spec + 1, value: h.p.value}
				if _, ok := visit[p]; !ok {
					visit[p] = struct{}{}

					hh := diffTrace{
						p: p,
						trace: append(
							sliceclone(h.trace),
							diff.MissingItem[A](ss[p.spec]),
						),
					}
					if hh.p == centinel {
						return hh.trace
					}
					newHead = append(newHead, hh)
				}
			}

			if h.p.value < centinel.value {
				p := position{spec: h.p.spec, value: h.p.value + 1}
				if _, ok := visit[p]; !ok {
					visit[p] = struct{}{}

					hh := diffTrace{
						p: p,
						trace: append(
							sliceclone(h.trace),
							diff.ExtraItem(as[p.value]),
						),
					}
					if hh.p == centinel {
						return hh.trace
					}
					newHead = append(newHead, hh)
				}
			}
		}

		head = newHead
	}
}

func sliceclone[T any](ts []T) []T {
	dest := make([]T, len(ts))
	copy(dest, ts)
	return dest
}
