package editorialgraph

import (
	"github.com/youta-t/its/internal/diff"
	"github.com/youta-t/its/itskit"
)

type position struct {
	posTo   int
	posFrom int
}

type diffTrace[T any] struct {
	p     position
	trace []diff.Diff[T]
}

func NewWithMatcher[A any](
	as []A, ss []itskit.Matcher[A],
) []diff.Diff[itskit.Match] {
	return New(
		as, ss,
		func(a A, m itskit.Matcher[A]) (itskit.Match, bool) {
			match := m.Match(a)
			return match, match.Ok()
		},
		diff.ExtraMatch[A],
		diff.MissingMatch[A],
	)
}

func New[A, B, C any](
	from []A,
	to []B,
	pred func(A, B) (C, bool),
	toDelete func(A) C,
	toInsert func(B) C,
) []diff.Diff[C] {
	// based Myers Algorithm.
	if len(to) == 0 && len(from) == 0 {
		return []diff.Diff[C]{}
	}

	head := []diffTrace[C]{
		{
			p:     position{posTo: -1, posFrom: -1},
			trace: []diff.Diff[C]{},
		},
	}
	visit := map[position]struct{}{}
	centinel := position{posTo: len(to) - 1, posFrom: len(from) - 1}

	for {
		{
			newHead := []diffTrace[C]{}
			for _, h := range head[:] {
				for h.p.posTo < centinel.posTo && h.p.posFrom < centinel.posFrom {
					p := position{posFrom: h.p.posFrom + 1, posTo: h.p.posTo + 1}
					if _, ok := visit[p]; ok {
						goto SKIP
					}

					vok, ok := pred(from[p.posFrom], to[p.posTo])
					if !ok {
						break
					}
					visit[p] = struct{}{}
					h = diffTrace[C]{p: p, trace: append(
						sliceclone(h.trace), diff.OkItem(vok),
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

		newHead := []diffTrace[C]{}
		for _, h := range head {
			if h.p.posTo < centinel.posTo {
				p := position{posTo: h.p.posTo + 1, posFrom: h.p.posFrom}
				if _, ok := visit[p]; !ok {
					visit[p] = struct{}{}

					hh := diffTrace[C]{
						p: p,
						trace: append(
							sliceclone(h.trace),
							diff.MissingItem(toInsert(to[p.posTo])),
						),
					}
					if hh.p == centinel {
						return hh.trace
					}
					newHead = append(newHead, hh)
				}
			}

			if h.p.posFrom < centinel.posFrom {
				p := position{posTo: h.p.posTo, posFrom: h.p.posFrom + 1}
				if _, ok := visit[p]; !ok {
					visit[p] = struct{}{}

					hh := diffTrace[C]{
						p: p,
						trace: append(
							sliceclone(h.trace),
							diff.ExtraItem(toDelete(from[p.posFrom])),
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
