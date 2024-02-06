package itskit_test

import (
	"fmt"
	"testing"

	"github.com/youta-t/its/itskit"
)

func TestWant(t *testing.T) {
	value := 42
	testee := itskit.Want(value)

	got := fmt.Sprintf("%d", testee)
	want := "/* want */ 42"
	if got != want {
		t.Errorf("\ngot: %s\nwant: %s\n", got, want)
	}
}

func TestEmptyGot(t *testing.T) {
	got := fmt.Sprintf("%d", itskit.Got)
	want := "/* got */ ??"
	if got != want {
		t.Errorf("\ngot: %s\nwant: %s\n", got, want)
	}
}

func TestLabel(t *testing.T) {
	testee := itskit.NewLabel(
		"%d (%s) == %d (%s)",
		itskit.Want(42), "fizz", itskit.Got, "bazz",
	)
	got := testee.Fill(113)
	want := "/* want */ 42 (fizz) == /* got */ 113 (bazz)"
	if got != want {
		t.Errorf("\ngot: %s\nwant: %s\n", got, want)
	}
}
