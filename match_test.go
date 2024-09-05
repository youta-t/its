package its_test

import (
	"regexp"

	"github.com/youta-t/its"
)

func ExampleMatch_ok() {
	pattern := regexp.MustCompile(`^[a-z]([a-z0-9.-]+[a-z])?$`)
	its.Match(pattern).Match([]byte("github.com")).OrError(t)
	// Output:
}

func ExampleMatch_ng() {
	pattern := regexp.MustCompile(`^[a-z]([a-z0-9.-]+[a-z])?$`)
	its.Match(pattern).Match([]byte("github.com/youta-t/its")).OrError(t)
	// Output:
	// ✘ (/* want */ ^[a-z]([a-z0-9.-]+[a-z])?$).Match(/* got */ [103 105 116 104 117 98 46 99 111 109 47 121 111 117 116 97 45 116 47 105 116 115])		--- @ ./match_test.go:17
}
