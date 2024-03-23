package config

import (
	"cmp"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
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

type replacer struct {
	From string
	To   string
}

func (r replacer) Cmp(r2 replacer) int { return cmp.Compare(r.From, r2.From) }

// When set, this filepath will be replaced with FilepathReplaceWith value.
// If this value is empty string, replacing is not performed.
//
// Default is value of envvar ITS_REPLACE_FILEPATH.
// When you set a special value "!project" to ITS_REPLACE_FILEPATH,
// the replace path will be the project root detected by heuristics.
//
// The project root is the most shallow filepath
// having ".git", "go.work" or "go.mod" from working directory.
var FilepathReplace = []replacer{}

func ReplaceFilepath(p string) string {
	repl := []replacer{}
	repl = append(repl, FilepathReplace...)
	slices.SortFunc(repl, replacer.Cmp)

	for _, r := range repl {
		if strings.HasPrefix(p, r.From) {
			p = p[len(r.From):]
			if p[0] == os.PathSeparator {
				p = p[1:]
			}
			prefix := r.To
			if prefix[len(prefix)-1] == os.PathSeparator {
				prefix = prefix[:len(prefix)-1]
			}

			if prefix == "." {
				return strings.Join([]string{prefix, p}, string(os.PathSeparator))
			}
			return path.Join(prefix, p)
		}
	}

	return p
}

func init() {
	replace := os.Getenv("ITS_REPLACE_FILEPATH")
	switch replace {
	case "":
	case "!project":
		ReplaceProjectRoot()
	default:
		replaceWith, ok := os.LookupEnv("ITS_REPLACE_FILEPATH_WITH")
		if ok {
			FilepathReplace = append(FilepathReplace, replacer{
				From: replaceWith, To: replaceWith,
			})
		}
	}

	cloakGoRoot := os.Getenv("ITS_CLOAK_GOROOT")
	switch cloakGoRoot {
	case "":
	default:
		goroot := os.Getenv("GOROOT")
		if goroot != "" {
			FilepathReplace = append(FilepathReplace, replacer{
				goroot, "(GOROOT)",
			})
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
		wd_ := wd
		for {
			gomod := filepath.Join(wd_, "go.mod")
			if _, err := os.Stat(gomod); err == nil {
				FilepathReplace = append(FilepathReplace, replacer{
					From: wd_, To: ".",
				})
				return
			}

			parent := filepath.Dir(wd_)
			if wd_ == parent {
				break
			}
			wd_ = parent
		}
	}

	{
		wd_ := wd
		for {
			gowork := filepath.Join(wd_, "go.work")
			if _, err := os.Stat(gowork); err == nil {
				FilepathReplace = append(FilepathReplace, replacer{
					From: wd_, To: ".",
				})
				return
			}

			parent := filepath.Dir(wd_)
			if wd_ == parent {
				break
			}
			wd_ = parent
		}
	}
	{
		wd_ := wd
		for {
			gitroot := filepath.Join(wd_, ".git")
			if _, err := os.Stat(gitroot); err == nil {
				FilepathReplace = append(FilepathReplace, replacer{
					From: wd_, To: ".",
				})
				return
			}

			parent := filepath.Dir(wd_)
			if wd_ == parent {
				break
			}
			wd_ = parent
		}
	}
}
