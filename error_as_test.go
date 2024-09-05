package its_test

import (
	"errors"

	"github.com/youta-t/its"
)

type CustomError string

func (c CustomError) Error() string {
	return string(c)
}

func ExampleErrorAs_ok() {
	customErr := CustomError("custom error")

	its.ErrorAs[CustomError]().Match(customErr).OrError(t)
	// Output:
}

func ExampleErrorAs_ng() {
	otherErr := errors.New("error by error.New")
	its.ErrorAs[CustomError]().Match(otherErr).OrError(t)

	// Output:
	// ✘ want := new(its_test.CustomError); errors.As(/* got */ error by error.New, want)		--- @ ./error_as_test.go:24
}
