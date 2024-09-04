package its_test

import (
	"time"

	"github.com/youta-t/its"
)

func ExampleBefore() {
	t1, err := time.Parse(
		time.RFC3339,
		"2023-10-11T12:13:14+00:00",
	)
	if err != nil {
		panic(err)
	}
	t2 := t1.Add(5 * time.Minute) // = 2023-10-11T12:18:14

	its.Before(t1).Match(t2).OrError(t)
	its.Before(t1).Match(t1).OrError(t)
	its.Before(t2).Match(t1).OrError(t)

	// Output:
	// ✘ (/* got */ 2023-10-11 12:18:14 +0000 +0000).Before(/* want */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./before_test.go:19
	//
	// ✘ (/* got */ 2023-10-11 12:13:14 +0000 +0000).Before(/* want */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./before_test.go:20
}
