package editorialgraph

import (
	"github.com/youta-t/its/internal/diff"
	"github.com/youta-t/its/itskit"
)

type position struct {
	posFrom int
	posTo   int
}

type diffTrace[T any] struct {
	p     position
	trace []diff.Diff[T]
}

func NewWithMatcher[A any](
	ss []itskit.Matcher[A], as []A,
) []diff.Diff[itskit.Match] {
	return New[itskit.Matcher[A], A, itskit.Match](
		ss, as,
		func(m itskit.Matcher[A], a A) (itskit.Match, bool) {
			match := m.Match(a)
			return match, match.Ok()
		},
		diff.MissingMatch[A],
		diff.ExtraMatch[A],
	)
}

func New[A, B, C any](
	from []A,
	to []B,
	pred func(A, B) (C, bool),
	toInsert func(A) C,
	toDelete func(B) C,
) []diff.Diff[C] {
	// based Myers Algorithm.
	if len(from) == 0 && len(to) == 0 {
		return []diff.Diff[C]{}
	}

	head := []diffTrace[C]{
		{
			p:     position{posFrom: -1, posTo: -1},
			trace: []diff.Diff[C]{},
		},
	}
	visit := map[position]struct{}{}
	centinel := position{posFrom: len(from) - 1, posTo: len(to) - 1}

	for {
		{
			newHead := []diffTrace[C]{}
			for _, h := range head[:] {
				for h.p.posFrom < centinel.posFrom && h.p.posTo < centinel.posTo {
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
			if h.p.posFrom < centinel.posFrom {
				p := position{posFrom: h.p.posFrom + 1, posTo: h.p.posTo}
				if _, ok := visit[p]; !ok {
					visit[p] = struct{}{}

					hh := diffTrace[C]{
						p: p,
						trace: append(
							sliceclone(h.trace),
							diff.MissingItem[C](toInsert(from[p.posFrom])),
						),
					}
					if hh.p == centinel {
						return hh.trace
					}
					newHead = append(newHead, hh)
				}
			}

			if h.p.posTo < centinel.posTo {
				p := position{posFrom: h.p.posFrom, posTo: h.p.posTo + 1}
				if _, ok := visit[p]; !ok {
					visit[p] = struct{}{}

					hh := diffTrace[C]{
						p: p,
						trace: append(
							sliceclone(h.trace),
							diff.ExtraItem(toDelete(to[p.posTo])),
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
