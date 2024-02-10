package config

import (
	"os"
	"path/filepath"
)

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

// When set, this filepath will be replaced with FilepathReplaceWith value.
// If this value is empty string, replacing is not performed.
//
// Default is value of envvar ITS_REPLACE_FILEPATH.
// When you set a special value "!project" to ITS_REPLACE_FILEPATH,
// the replace path will be the project root detected by heuristics.
//
// The project root is the most shallow filepath
// having ".git", "go.work" or "go.mod" from working directory.
var FilepathReplace = ""

// Replacement by FilepathReplace.
//
// Default is ".".
//
// When ennvat ITS_REPLCE_FILEPATH_WITH is set, use that value as default.
var FilepathReplaceWith = "."

func init() {
	replace := os.Getenv("ITS_REPLACE_FILEPATH")
	switch replace {
	case "":
	case "!project":
		ReplaceProjectRoot()
	default:
		FilepathReplace = replace
		replaceWith, ok := os.LookupEnv("ITS_REPLACE_FILEPATH_WITH")
		if ok {
			FilepathReplaceWith = replaceWith
		}
	}
}

// Work as start test with envvar ITS_REPLACE_FILEPATH=!project .
func ReplaceProjectRoot() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	{
		for {
			gomod := filepath.Join(wd, "go.mod")
			if _, err := os.Stat(gomod); err == nil {
				FilepathReplace = wd
				break
			}

			parent := filepath.Dir(wd)
			if wd == parent {
				break
			}
			wd = parent
		}
	}

	{
		wd := FilepathReplace
		for {
			gomod := filepath.Join(wd, "go.work")
			if _, err := os.Stat(gomod); err == nil {
				FilepathReplace = wd
				break
			}

			parent := filepath.Dir(wd)
			if wd == parent {
				break
			}
			wd = parent
		}
	}
	{
		wd := FilepathReplace
		for {
			gomod := filepath.Join(wd, ".git")
			if _, err := os.Stat(gomod); err == nil {
				FilepathReplace = wd
				break
			}

			parent := filepath.Dir(wd)
			if wd == parent {
				break
			}
			wd = parent
		}
	}
}
