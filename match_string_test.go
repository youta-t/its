package its_test

import (
	"regexp"

	"github.com/youta-t/its"
)

func ExampleMatchString() {
	pattern := regexp.MustCompile(`^[a-z]([a-z0-9.-]+[a-z])?$`)
	its.MatchString(pattern).Match("github.com").OrError(t)
	// Output:
}

func ExampleMatchString_ng() {
	pattern := regexp.MustCompile(`^[a-z]([a-z0-9.-]+[a-z])?$`)
	its.MatchString(pattern).Match("github.com/youta-t/its").OrError(t)
	// Output:
	// ✘ (/* want */ ^[a-z]([a-z0-9.-]+[a-z])?$).MatchString(/* got */ "github.com/youta-t/its")		--- @ ./match_string_test.go:17
}
