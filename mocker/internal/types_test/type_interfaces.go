//go:generate go run github.com/youta-t/its/mocker
package types

import (
	"github.com/youta-t/its/mocker/internal/example/sub"
)

type I0 interface {
	M0()
	M1(int, string) (bool, error)
	M2(i int, s string) (ok bool, err error)
	M3(i int, s ...string) (ok bool, err error)
	M4(s ...string) bool
}

type I1[S, T sub.T, U X[T]] interface {
	M0(S, T) (U, error)
	M1(S)
	M2() U
}

type X[T any] struct{}

type I2[T ~string] interface {
	M0(T) T
}
