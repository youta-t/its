package main

import (
	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/youta-t/its/structer/internal"
	"github.com/youta-t/its/structer/internal/try"
)

/*

structer -s Struct -as-package github.com/... ./dest

*/

type names map[string]struct{}

func (n names) Set(v string) error {
	n[v] = struct{}{}
	return nil
}

func (n names) String() string {
	if l := len(n); l == 0 {
		return ""
	}

	nslice := []string{}
	for name := range n {
		nslice = append(nslice, name)
	}
	sort.Strings(nslice)

	return strings.Join(nslice, " ")
}

func main() {
	logger := log.Default()
	invokedFrom := os.Getenv("GOFILE")

	flag.Usage = func() {
		name := os.Args[0]
		shortname := filepath.Base(name)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", name)
		fmt.Fprintf(
			flag.CommandLine.Output(),
			`%s is a matcher generator for any struct.
This is designed to be used as go:generate.

It generates a file with same name as a file having go:generate directive.
The new file, has "Matcher" and "Spec" types, is placed in "./gen" directory (by default).
`,
			shortname,
		)
		flag.PrintDefaults()
	}

	targetStructs := names{}
	flag.Var(targetStructs, "struct", `Struct names to generate Matcher. Repeatable. By default, all structs are target.`)
	flag.Var(targetStructs, "s", "alias of -struct")

	sourceAsPackage := flag.Bool("as-package", false, "handle -source as package path")
	flag.BoolVar(sourceAsPackage, "p", false, "alias of -as-package")

	psource := flag.String("source", invokedFrom, "recognise source as package path. If not set, use environmental variable GOFILE.")
	pdest := flag.String("dest", "./gen", "directory where new file to be created at")

	flag.Parse()

	source := *psource
	dest := *pdest

	if source == "" {
		log.Fatalf("-source is required")
		flag.Usage()
		return
	}
	if dest == "" {
		log.Fatalf("-dest is required")
		flag.Usage()
		return
	}

	sources := []string{}

	if !*sourceAsPackage {
		sources = append(sources, source)
	} else {
		logger.Printf("parse as package: %s", source)
		p := try.To(build.Default.Import(source, ".", 0)).OrFatal(logger)
		for _, gof := range p.GoFiles {
			sources = append(sources, filepath.Join(p.Dir, gof))
		}
		logger.Println("source files are found!:", sources, ", in dir:", p.Dir)
	}

	for _, s := range sources {
		f := try.To(internal.Parse(s)).OrFatal(logger)
		newFile := generatingFile{
			PackageName: path.Base(dest),
		}
		requiredImports := map[string]*internal.Import{
			"its":    {Name: "its", Path: "github.com/youta-t/its"},
			"itskit": {Name: "itskit", Path: "github.com/youta-t/its/itskit"},
			"itsio":  {Name: "itsio", Path: "github.com/youta-t/its/itskit/itsio"},
			"config": {Name: "config", Path: "github.com/youta-t/its/config"},
		}
		usedImports := map[string]*internal.Import{}

		for i := range f.Structs {
			s := f.Structs[i]

			if _, ok := targetStructs[s.Name]; len(targetStructs) != 0 && !ok {
				continue
			}

			newFile.Structs = append(newFile.Structs, s)

			types := s.Require()
			for i := range types {
				t := types[i]
				usedImports[t.Name] = t
			}
		}

		if len(newFile.Structs) == 0 {
			continue
		}

		for i := range requiredImports {
			newFile.Imports = append(newFile.Imports, requiredImports[i])
		}

		for i := range usedImports {
			imp := usedImports[i]
			if imp.Name == "" {
				imp.Name = "testee"
			} else {
				imp.Name = "u_" + imp.Name
			}
			newFile.Imports = append(newFile.Imports, imp)
		}

		if err := os.MkdirAll(dest, os.ModeDir|os.ModePerm); err != nil {
			logger.Fatal(err)
		}
		t := template.New("")
		t = try.To(t.Parse(tpl)).OrFatal(logger)

		newFilename := filepath.Join(dest, filepath.Base(s))
		gen := try.To(os.OpenFile(
			newFilename, os.O_TRUNC|os.O_RDWR|os.O_CREATE, os.ModePerm,
		)).OrFatal(logger)
		defer gen.Close()
		if err := t.Execute(gen, newFile); err != nil {
			logger.Fatal(err)
		}
	}
	os.Exit(0)
}

type generatingFile struct {
	PackageName string
	Imports     []*internal.Import
	Structs     []*internal.StructDecl
}

const tpl = `// Code generated -- DO NOT EDIT

package {{ .PackageName }}
import (
	"strings"

	{{ range .Imports }}
	{{- .Name }} "{{ .Path }}"
	{{ end }}
)
{{ range .Structs }}
{{ $s := . }}
type {{ .Name }}Spec{{ .GenericExpr true }} struct {
	{{ range .Body.Fields -}}
	{{- .Name }} its.Matcher[{{ .Type.Expr }}]
	{{ end }}
}

type _{{ .Name }}Matcher{{ .GenericExpr true }} struct {
	fields []its.Matcher[{{ .Expr }}]
}

func Its{{ .Name }}{{ .GenericExpr true }}(want {{ .Name }}Spec{{ .GenericExpr false }}) its.Matcher[{{ .Expr }}] {
	sub := []its.Matcher[{{ .Expr }}]{}
	{{ range .Body.Fields }}
	{
		matcher := want.{{ .Name }}
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[{{ .Type.Expr }}]()
			} else {
				matcher = its.Always[{{ .Type.Expr }}]()
			}
		}
		sub = append(
			sub,
			itskit.Property[{{ $s.Expr }}, {{ .Type.Expr }}](
				".{{ .Name }}",
				func(got {{ $s.Expr }}) {{ .Type.Expr }} { return got.{{ .Name }} },
				matcher,
			),
		)
	}
	{{ end }}

	return _{{ .Name }}Matcher{{ .GenericExpr false }}{ fields: sub }
}

func (m _{{ .Name }}Matcher{{ .GenericExpr false }}) Match(got {{ .Expr }}) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(
		len(sub) == ok,
		itskit.NewLabel("type {{ .Name }}:").Fill(struct{}{}),
		sub...,
	)
}

func (m _{{ .Name }}Matcher{{ .GenericExpr false }}) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type {{ .Name }}:", m.fields)
}

func (m _{{ .Name }}Matcher{{ .GenericExpr false }}) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}
{{ end }}
`
