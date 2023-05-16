package its_test

// utilities for testing its-self

import "fmt"

// FakeT is fake of *testing.T
type FakeT struct{}

func (*FakeT) Error(values ...any) {
	fmt.Println(values...)
}

var t = new(FakeT)
