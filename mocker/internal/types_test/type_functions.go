//go:generate go run github.com/youta-t/its/mocker
package types

type F1 func()
type F2 func() int
type F3 func() (int, string)
type F4 func() (i int, s string)
type F5 func(int, string)
type F6 func(int, ...string)
type F7 func(...string)
type F8 func(i int, s string)
type F9 func(i int, ss ...string)
type F10 func(ss ...string)
type F11 func(_ int, _ string)
type F12 func(_ int, _ ...string)
type F13 func(_ ...string)
type F14 func(int, string) (f float64)
type F15 func(i int, ss ...string) float64
type F16[T any] func(F16[T]) T
type F17[T struct{ Foo int }] func(T) T
type F18[T, U any] func(func(T, U)) func(T, U)
type F19[T ~int] func(F16[T]) T
type F20[T ~int | int8 | ~int16 | int32] func(T) T

type F21 func(p private) bool
type F22 func() private
type F23[T private] func() T
type f24 func()

func init() {
	var x f24 = nil
	_ = x
}
