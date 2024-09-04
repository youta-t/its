package its_test

import (
	"errors"
	"fmt"

	"github.com/youta-t/its"
)

func ExampleError() {
	e1 := errors.New("error")
	e2 := fmt.Errorf("wrapped: %w", e1)

	its.Error(e1).Match(e2).OrError(t)
	its.Error(e1).Match(e1).OrError(t)
	// Output:
}

func ExampleError_ng() {
	e1 := errors.New("error")
	e2 := fmt.Errorf("wrapped: %w", e1)

	its.Error(e2).Match(e1).OrError(t)

	// Output:
	// ✘ errors.Is(/* got */ error, /* want */ wrapped: error)		--- @ ./error_test.go:23
}
