package its_test

import (
	"time"

	"github.com/youta-t/its"
)

func ExampleEqual() {
	t1, err := time.Parse(
		time.RFC3339,
		"2023-10-11T12:13:14+00:00",
	)
	if err != nil {
		panic(err)
	}
	t2 := t1.Add(5 * time.Minute) // = 2023-10-11T12:18:14

	its.Equal(t1).Match(t2).OrError(t)
	its.Equal(t1).Match(t1).OrError(t)
	its.Equal(t2).Match(t1).OrError(t)

	// Output:
	// ✘ (/* want */ 2023-10-11 12:13:14 +0000 +0000).Equal(/* got */ 2023-10-11 12:18:14 +0000 +0000)		--- @ ./equal_test.go:19
	//
	// ✘ (/* want */ 2023-10-11 12:18:14 +0000 +0000).Equal(/* got */ 2023-10-11 12:13:14 +0000 +0000)		--- @ ./equal_test.go:21
}
