package its_test

import (
	"time"

	"github.com/youta-t/its"
)

func ExampleAfter() {
	t1, err := time.Parse(
		time.RFC3339,
		"2023-10-11T12:13:14+00:00",
	)
	if err != nil {
		panic(err)
	}
	t2 := t1.Add(5 * time.Minute) // = 2023-10-11T12:18:14

	its.After(t1).Match(t2).OrError(t)
	its.After(t1).Match(t1).OrError(t)
	its.After(t2).Match(t1).OrError(t)

	// Output:
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).After(/* want */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./after_test.go:20
	//
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).After(/* want */ 2023-10-11 12:18:14 +0000 +0000)		--- @ ./after_test.go:21
}
