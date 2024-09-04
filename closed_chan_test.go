package its_test

import "github.com/youta-t/its"

func ExampleClosedChan_ok() {
	ch1 := make(chan int, 1)
	close(ch1)
	its.ClosedChan[chan int]().Match(ch1).OrError(t)
	// Output:
}

func ExampleClosedChan_ng() {
	ch2 := make(chan string, 1)
	its.ClosedChan[chan string]().Match(ch2).OrError(t)
	// Output:
	// ✘ chan string is not closed.		--- @ ./closed_chan_test.go:14
}
