package itskit

// itskit is a development kit to create your Matcher.
//
// # Basic: Create New Matcher
//
// There are 2 ways.
//
// 1. SimpleMatcher
//
// If your assertion can be a short function, and placeholders in Label is Got only,
// you can use SimpleMatcher utility.
//
// 	var Odd = SimpleMatcher(
// 		func(got int) bool { return got % 2 == 1 },  // true if pass, false if failed.
// 		"%d is odd",	// label template
// 		itskit.Got,		// label parameters (variadic). You can pass placeholders as like NewLabel.
// 	}
//
// Now you have matcher Odd, passing if got value is odd number.
//
// 2. Custom Mathcer
//
// If your assertion is not simple, you need to implement your Matcher.
//
// To implement, write 3 methods.
//
// - Match(T)itskit.Match: receive got value, return itskit.Match, represens the result of matching.
// You can create Match with itskit.OK (for passed), itskit.NG (for failed) or itskit.NewMatch.
// Match can be nested. Inner Match is indented when it is printed as (error) results.
//
// - Write(itskit.Writer)error: decrate string expression of this Matcher.
// Write expressions into given itskit.Writer.
//
// - String() string: Get string expression as string value.
// You can implement this just like below:
//
// 	func (m YourMatcher) String() string {
// 		return itskit.MatcherToString[GotType](m)
// 	}
//
