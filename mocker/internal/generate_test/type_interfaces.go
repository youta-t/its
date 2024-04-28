//go:generate go run github.com/youta-t/its/mocker
//go:generate gofmt -w gen_mock/type_interfaces.go
package generatetest

import (
	"io"

	"github.com/youta-t/its/mocker/internal/example/sub"
	. "github.com/youta-t/its/mocker/internal/generate_test/dot"
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

type I3 interface {
	P() private
}

type I4 interface {
	P(private)
}

type I5[T private] interface {
	P() T
}

type i6 interface {
}

type C1 interface {
	string | int
}
type C2 interface {
	~string
}

type S struct{}

type CS interface {
	S
}

type C3 interface {
	io.Reader
	io.Writer
}

type C4 interface {
	DotInterface

	interface {
		Method()
	}
	interface {
		io.Reader
		interface {
			AnotherMethod()
			io.Closer
		}
	}
}

type C5 interface {
	DotStruct

	M()
}

func init() {
	var x i6
	_ = x
}
