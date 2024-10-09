package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	parser "github.com/youta-t/its/internal/parser"
	"github.com/youta-t/its/internal/try"
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
The new file, has "Matcher" and "Spec" types, is placed in "./gen_structer" directory (by default).

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
	pdest := flag.String("dest", "./gen_structer", "directory where new file to be created at")

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

	parserInstance := try.To(parser.New()).OrFatal(logger)

	var pkg *parser.Package
	targetFile := ""

	if *sourceAsPackage {
		pkg = try.To(parserInstance.Import(source)).OrFatal(logger)
	} else {
		source = try.To(filepath.Abs(source)).OrFatal(logger)
		dir := filepath.Dir(source)
		targetFile = source
		pkg = try.To(parserInstance.ImportDir(dir)).OrFatal(logger)
	}

	filenames := map[string]struct{}{}

	structs := map[string][]*parser.TypeDecl[*parser.StructType]{}
	{
		_structs := pkg.Types.Structs.Slice()
		for i := range _structs {
			s := _structs[i]
			if targetFile != "" && s.DefinedIn != targetFile {
				continue
			}
			filenames[s.DefinedIn] = struct{}{}
			structs[s.DefinedIn] = append(structs[s.DefinedIn], s)
		}
	}

	maps := map[string][]*parser.TypeDecl[*parser.MapType]{}
	{
		_maps := pkg.Types.Maps.Slice()
		for i := range _maps {
			m := _maps[i]
			if targetFile != "" && m.DefinedIn != targetFile {
				continue
			}
			filenames[m.DefinedIn] = struct{}{}
			maps[m.DefinedIn] = append(maps[m.DefinedIn], m)
		}
	}

	slices := map[string][]*parser.TypeDecl[*parser.SliceType]{}
	{
		_slices := pkg.Types.Slices.Slice()
		for i := range _slices {
			s := _slices[i]
			if targetFile != "" && s.DefinedIn != targetFile {
				continue
			}
			filenames[s.DefinedIn] = struct{}{}
			slices[s.DefinedIn] = append(slices[s.DefinedIn], s)
		}
	}

	arrays := map[string][]*parser.TypeDecl[*parser.ArrayType]{}
	{
		_arrays := pkg.Types.Arrays.Slice()
		for i := range _arrays {
			a := _arrays[i]
			if targetFile != "" && a.DefinedIn != targetFile {
				continue
			}
			filenames[a.DefinedIn] = struct{}{}
			arrays[a.DefinedIn] = append(arrays[a.DefinedIn], a)
		}
	}

	for sourcepath := range filenames {
		genFile := generatingFile{
			PackageName:    path.Base(dest),
			ConfigIsNeeded: false,
			Imports:        new(parser.Imports),
		}

		requires := map[string]struct{}{}
		requires[pkg.Path] = struct{}{}

		{
			defs := structs[sourcepath]
			for i := range defs {
				s := defs[i]
				genFile.ConfigIsNeeded = genFile.ConfigIsNeeded || 0 < len(s.Body.Fields)
				genFile.Structs = append(genFile.Structs, s)
				for _, req := range s.Require() {
					requires[req] = struct{}{}
				}
			}
		}
		{
			defs := maps[sourcepath]
			for i := range defs {
				m := defs[i]
				genFile.Maps = append(genFile.Maps, m)
				for _, req := range m.Require() {
					requires[req] = struct{}{}
				}
			}
		}
		{
			defs := slices[sourcepath]
			for i := range defs {
				m := defs[i]
				genFile.Slices = append(genFile.Slices, m)
				for _, req := range m.Require() {
					requires[req] = struct{}{}
				}
			}
		}
		{
			defs := arrays[sourcepath]
			for i := range defs {
				m := defs[i]
				genFile.Arrays = append(genFile.Arrays, m)
				for _, req := range m.Require() {
					requires[req] = struct{}{}
				}
			}
		}
		{
			rs := []string{}
			for r := range requires {
				rs = append(rs, r)
			}
			sort.Strings(rs)
			for _, r := range rs {
				genFile.Imports.Add(r)
			}
		}
		if len(genFile.Structs) == 0 {
			continue
		}

		fname := filepath.Base(sourcepath)
		if err := writeFile(filepath.Join(dest, fname), genFile); err != nil {
			logger.Fatal(err)
		}
	}

	os.Exit(0)
}

func writeFile(dest string, newFile generatingFile) error {
	d := filepath.Dir(dest)
	if err := os.MkdirAll(d, os.ModeDir|os.ModePerm); err != nil {
		return err
	}
	t := template.New("")
	t, err := t.Parse(tpl)
	if err != nil {
		return err
	}

	gen, err := os.OpenFile(
		dest, os.O_TRUNC|os.O_RDWR|os.O_CREATE, os.ModePerm,
	)
	if err != nil {
		return err
	}
	defer gen.Close()
	if err := t.Execute(gen, newFile); err != nil {
		return err
	}
	return nil
}

type generatingFile struct {
	PackageName    string
	ConfigIsNeeded bool
	Imports        *parser.Imports
	Structs        []*parser.TypeDecl[*parser.StructType]
	Maps           []*parser.TypeDecl[*parser.MapType]
	Slices         []*parser.TypeDecl[*parser.SliceType]
	Arrays         []*parser.TypeDecl[*parser.ArrayType]
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

const tpl = `// Code generated -- DO NOT EDIT
{{ $imports := .Imports }}
package {{ .PackageName }}
import (
	"strings"

	its "github.com/youta-t/its"
	itskit "github.com/youta-t/its/itskit"
	itsio "github.com/youta-t/its/itskit/itsio"
	{{ if .ConfigIsNeeded }}config "github.com/youta-t/its/config"{{ end }}

	{{ range .Imports.Slice }}
	{{- .Name }} "{{ .Path }}"
	{{ end }}
)
{{ range .Structs }}
{{ $s := . }}
type {{ .Name }}Spec{{ .GenericExpr $imports true }} struct {
	{{ range .Body.Fields }}{{ if .IsOpaque }}{{ continue }}{{ end }}
	{{ .Name }} its.Matcher[{{ .Type.Expr $imports }}]
	{{ end }}
}

type _{{ .Name }}Matcher{{ .GenericExpr $imports true }} struct {
	label  itskit.Label
	fields []its.Matcher[{{ .Expr $imports }}]
}

func Its{{ .Name }}{{ .GenericExpr $imports true }}(want {{ .Name }}Spec{{ .GenericExpr $imports false }}) its.Matcher[{{ .Expr $imports }}] {
	cancel := itskit.SkipStack()
	defer cancel()

	sub := []its.Matcher[{{ .Expr $imports }}]{}
	{{ range .Body.Fields }}{{ if .IsOpaque }}{{ continue }}{{ end }}
	{
		matcher := want.{{ .Name }}
		if matcher == nil {
			if config.StrictModeForStruct {
				matcher = its.Never[{{ .Type.Expr $imports }}]()
			} else {
				matcher = its.Always[{{ .Type.Expr $imports }}]()
			}
		}
		sub = append(
			sub,
			its.Property[{{ $s.Expr $imports }}, {{ .Type.Expr $imports }}](
				".{{ .Name }}",
				func(got {{ $s.Expr $imports }}) {{ .Type.Expr $imports }} { return got.{{ .Name }} },
				matcher,
			),
		)
	}
	{{ end }}

	return _{{ .Name }}Matcher{{ .GenericExpr $imports false }}{
		label: itskit.NewLabelWithLocation("type {{ .Name }}:"),
		fields: sub,
	}
}

func (m _{{ .Name }}Matcher{{ .GenericExpr $imports false }}) Match(got {{ .Expr $imports }}) itskit.Match {
	ok := 0
	sub := []itskit.Match{}
	for _, f := range m.fields {
		m := f.Match(got)
		if m.Ok() {
			ok += 1
		}
		sub = append(sub, m)
	}

	return itskit.NewMatch(len(sub) == ok, m.label.Fill(got), sub...)
}

func (m _{{ .Name }}Matcher{{ .GenericExpr $imports false }}) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type {{ .Name }}:", m.fields)
}

func (m _{{ .Name }}Matcher{{ .GenericExpr $imports false }}) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}
{{ end }}

{{ range .Maps }}
func Its{{ .Name }}{{ .GenericExpr $imports true }}(want its.Matcher[map[{{ .Body.Key.Expr $imports }}]{{ .Body.Elem.Expr $imports }}]) its.Matcher[ {{ .Expr $imports }} ] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _{{ .Name }}Matcher{{ .GenericExpr $imports false }}{matchers: want}
}

type _{{ .Name }}Matcher{{ .GenericExpr $imports true }} struct {
	matchers its.Matcher[map[{{ .Body.Key.Expr $imports }}]{{ .Body.Elem.Expr $imports }}]
}

func (m _{{ .Name }}Matcher{{ .GenericExpr $imports false }}) Match(got {{ .Expr $imports }}) itskit.Match {
	gotm := map[{{ .Body.Key.Expr $imports }}]{{ .Body.Elem.Expr $imports }}(got)
	return m.matchers.Match(gotm)
}

func (m _{{ .Name }}Matcher{{ .GenericExpr $imports false }}) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type {{ .Name }}:", []its.Matcher[map[{{ .Body.Key.Expr $imports }}]{{ .Body.Elem.Expr $imports }}]{m.matchers})
}

func (m _{{ .Name }}Matcher{{ .GenericExpr $imports false }}) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}
{{ end }}

{{ range .Slices }}
func Its{{ .Name }}{{ .GenericExpr $imports true }}(want its.Matcher[[]{{ .Body.Elem.Expr $imports }}]) its.Matcher[{{ .Expr $imports }}] {
	cancel := itskit.SkipStack()
	defer cancel()

	return _{{ .Name }}Matcher{{ .GenericExpr $imports false }}{ matcher: want }
}

type _{{ .Name }}Matcher{{ .GenericExpr $imports true }} struct {
	matcher its.Matcher[[]{{ .Body.Elem.Expr $imports }}]
}

func (m _{{ .Name }}Matcher{{ .GenericExpr $imports false }}) Match(got {{ .Expr $imports }}) itskit.Match {
	gots := []{{ .Body.Elem.Expr $imports }}(got)
	return m.matcher.Match(gots)
}

func (m _{{ .Name }}Matcher{{ .GenericExpr $imports false }}) Write(ww itsio.Writer) error {
	return itsio.WriteBlock(ww, "type {{ .Name }}:", []its.Matcher[[]{{ .Body.Elem.Expr $imports }}]{m.matcher})
}

func (m _{{ .Name }}Matcher{{ .GenericExpr $imports false }}) String() string {
	sb := new(strings.Builder)
	w := itsio.Wrap(sb)
	m.Write(w)
	return sb.String()
}
{{ end }}

`
