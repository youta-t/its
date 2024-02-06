package config

// Indent of messages
var INDENT = "    "

// Formatter for failed Match
var Failed = func(s string) string {
	return FailedSigil + s
}

// Sigil for failed Match.
//
// default: "✘ "
var FailedSigil = "✘ "

// Formatter for passed Match
var Pass = func(s string) string {
	return PassSigil + s
}

// Sigil for passed Match
//
// default; "✔ "
var PassSigil = "✔ "

// Formatter for failed Match in passed test
var FailedSuppressed = func(s string) string {
	sig := FailedSuppressedSigil
	return sig + s
}

// Sigil for failed but suppressed Match
//
// default: "* "
var FailedSuppressedSigil = "~ "
