package its_test

// utilities for testing its-self

import (
	"fmt"
	"os"
	"testing"

	"github.com/youta-t/its/config"
)

// FakeT is fake of *testing.T
type FakeT struct{}

func (*FakeT) Error(values ...any) {
	fmt.Println(values...)
}

var t = new(FakeT)

func TestMain(m *testing.M) {
	config.ReplaceProjectRoot()
	os.Exit(m.Run())
}
