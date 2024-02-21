package main

import (
	"flag"
	"fmt"
	"go/build"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/youta-t/its/internal/parser"
	"github.com/youta-t/its/internal/try"
)

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

type generatingFile struct {
	PackageName string
	Imports     []*parser.Import
	Interfaces  []*parser.TypeInterfaceDecl
	Funcs       []*parser.TypeFuncDecl
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
			`%s is a mock generator for any interfaces or functions.
This is designed to be used as go:generate.

It generates a file with same name as a file having go:generate directive.

`,
			shortname,
		)
		flag.PrintDefaults()
	}

	targetTypeName := names{}
	flag.Var(targetTypeName, "type", `Type names to generate Mock. Repeatable. By default, all interfaces and type func are target.`)
	flag.Var(targetTypeName, "t", "alias of -type")

	sourceAsPackage := flag.Bool("as-package", false, "handle -source as package path")
	flag.BoolVar(sourceAsPackage, "p", false, "alias of -as-package")

	psource := flag.String("source", invokedFrom, "recognise source as package path. If not set, use environmental variable GOFILE.")
	pdest := flag.String("dest", "./gen_mock", "directory where new file to be created at")

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
	}

	for _, s := range sources {
		f := try.To(parser.Parse(s)).OrFatal(logger)
		newFile := generatingFile{
			PackageName: path.Base(dest),
		}
		requiredImports := map[string]*parser.Import{
			"its":    {Name: "its", Path: "github.com/youta-t/its"},
			"itskit": {Name: "itskit", Path: "github.com/youta-t/its/itskit"},
		}
		usedImports := map[string]*parser.Import{}

		for i := range f.Types.Interfaces {
			s := f.Types.Interfaces[i]

			if _, ok := targetTypeName[s.Name]; len(targetTypeName) != 0 && !ok {
				continue
			}

			newFile.Interfaces = append(newFile.Interfaces, s)

			types := s.Require()
			for i := range types {
				t := types[i]
				usedImports[t.Name] = t
			}
		}
		for i := range f.Types.Funcs {
			s := f.Types.Funcs[i]

			if _, ok := targetTypeName[s.Name]; len(targetTypeName) != 0 && !ok {
				continue
			}

			newFile.Funcs = append(newFile.Funcs, s)

			types := s.Require()
			for i := range types {
				t := types[i]
				usedImports[t.Name] = t
			}
		}

		if len(newFile.Interfaces)+len(newFile.Funcs) == 0 {
			continue
		}

		for i := range requiredImports {
			newFile.Imports = append(newFile.Imports, requiredImports[i])
		}
		slices.SortFunc(newFile.Imports, func(a, b *parser.Import) int {
			if a.Path < b.Path {
				return -1
			}
			if b.Path < a.Path {
				return 1
			}
			return 0
		})

		for i := range usedImports {
			imp := usedImports[i]
			if imp.Name == "" {
				imp.Name = "testee"
			} else {
				imp.Name = "u_" + imp.Name
			}
			newFile.Imports = append(newFile.Imports, imp)
		}

		funcs := newFile.Funcs

		for _, intf := range newFile.Interfaces {
			for _, m := range intf.Body.Methods {
				fd := &parser.TypeFuncDecl{
					Name:       fmt.Sprintf("%s_%s", intf.Name, m.Name),
					Package:    intf.Package,
					TypeParams: m.Func.TypeParams(),
					Body:       m.Func,
				}
				funcs = append(funcs, fd)
			}
		}
		newFile.Funcs = funcs

		err := writeFile(filepath.Join(dest, source), newFile)
		if err != nil {
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
	t.Funcs(template.FuncMap{
		"export": func(val string) string {
			initial := val[0]
			rest := val[1:]
			return strings.ToUpper(string(initial)) + rest
		},
		"genericExpr": func(back bool, t parser.Type) string {
			params := []string{}
			for _, p := range t.TypeParams() {
				expr := p.Name
				if back {
					expr += " " + p.Back.Expr()
				}
				params = append(params, expr)
			}

			if len(params) == 0 {
				return ""
			}
			return "[" + strings.Join(params, ", ") + "]"
		},
	})

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
package {{ .PackageName }}

import (
	{{ range .Imports }}
	{{- .Name }} "{{ .Path }}"
	{{ end }}
)

{{- range .Funcs -}}
{{- $func := . }}

type _{{ .Name }}ReturnFixture{{ .GenericExpr true }} struct {
	{{- range $idx, $p := .Body.Returns}}
	{{ $p.ParamNameOr (printf "ret%d" $idx) }} {{ $p.Expr }}
	{{ end }}
}

type _{{ .Name }}Return{{ .GenericExpr true }} struct {
	fixture _{{ .Name }}ReturnFixture{{ .GenericExpr false }}
}

func (rfx _{{ .Name }}Return{{ .GenericExpr false }}) Get() (
	{{- range $idx, $p := .Body.Returns}}
	{{ $p.Expr }},
	{{ end }}
) {
	return {{ range $idx, $p := .Body.Returns }}{{ if ne $idx 0}}, {{ end }}rfx.fixture.{{ $p.ParamNameOr (printf "ret%d" $idx) }}{{ end }}
}

type _{{ .Name }}CallSpec{{ .GenericExpr true }} struct {
	{{- range $idx, $p := .Body.Args}}
	{{ $p.ParamNameOr (printf "arg%d" $idx) }} its.Matcher[{{ $p.Expr }}]
	{{ end }}
	{{ if .Body.VarArg }}{{ .Body.VarArg.ParamNameOr "vararg" }} its.Matcher[[]{{ .Body.VarArg.Type.Expr }}]{{ end }}
}

type _{{ .Name }}Call{{ .GenericExpr true }} struct {
	name itskit.Label
	spec _{{ .Name }}CallSpec{{ .GenericExpr false }}
}

func New{{ .Name }}Call{{ .GenericExpr true }}(
	{{- range $idx, $p := .Body.Args}}
	{{ $p.ParamNameOr (printf "arg%d" $idx) }} its.Matcher[{{ $p.Expr }}],
	{{ end -}}
	{{ if .Body.VarArg }}{{ .Body.VarArg.ParamNameOr "vararg" }} its.Matcher[[]{{ .Body.VarArg.Type.Expr }}],{{ end }}
) _{{ .Name }}Call{{ .GenericExpr false }} {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _{{ .Name }}CallSpec{{ .GenericExpr false }} {}
	{{- range $idx, $p := .Body.Args}}
	spec.{{ $p.ParamNameOr (printf "arg%d" $idx) }} = itskit.Named(
		"{{ $p.ParamNameOr (printf "arg%d" $idx) }}",
		{{ $p.ParamNameOr (printf "arg%d" $idx) }},
	)
	{{ end }}
	{{ if .Body.VarArg }}spec.{{ .Body.VarArg.ParamNameOr "vararg" }} = itskit.Named(
		"{{ .Body.VarArg.ParamNameOr "vararg" }}",
		{{ .Body.VarArg.ParamNameOr "vararg" }},
	){{ end }}
	return _{{ .Name }}Call{{ .GenericExpr false }}{
		name: itskit.NewLabelWithLocation("func {{ .Name }}"),
		spec: spec,
	}
}

type {{ .Name }}Behaviour {{ .GenericExpr true }} struct {
	name itskit.Label
	spec _{{ .Name }}CallSpec{{ .GenericExpr false }}
	effect {{ .Body.Expr }}
}

func (b {{ .Name }}Behaviour{{ .GenericExpr false }}) Mock(t interface { Error(...any) }) {{ .Body.Expr }} {
	return func (
		{{ range $idx, $p :=  .Body.Args }}
		{{ printf "arg%d" $idx }} {{ $p.Expr }},
		{{ end }}
		{{ if .Body.VarArg }}vararg {{ .Body.VarArg.Expr }},{{ end }}
	) {{ if .Body.Returns }}(
		{{ range .Body.Returns }}{{ .Expr }},
		{{ end }}
	){{ end }} {
		if h, ok := t.(interface { Helper() }); ok {
			h.Helper()
		}
		ok := 0
		matches := []itskit.Match{}
		{{ range $idx, $p := .Body.Args}}
		{
			matcher := b.spec.{{ $p.ParamNameOr (printf "arg%d" $idx) }}
			if matcher == nil {
				matcher = its.Never[{{$p.Expr}}]()
			}
			m := matcher.Match({{ printf "arg%d" $idx }})
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}
		{{ end }}
		{{- if .Body.VarArg }}
		{
			matcher := b.spec.{{ .Body.VarArg.ParamNameOr "vararg" }}
			if matcher == nil {
				matcher = its.Never[[]{{ .Body.VarArg.Type.Expr }}]()
			}
			m := matcher.Match(vararg)
			if m.Ok() {
				ok += 1
			}
			matches = append(matches, m)
		}
		{{ end }}
		itskit.NewMatch(
			ok == len(matches),
			b.name.Fill(itskit.Missing),
			matches...,
		).OrError(t)
		{{ if .Body.Returns }}return {{ end }}b.effect(
			{{ range $idx, $p := .Body.Args}}
			{{ printf "arg%d" $idx }},
			{{ end }}
			{{ if .Body.VarArg }}
			vararg...,
			{{ end }}
		)
	}
}

func (c _{{.Name}}Call{{ .GenericExpr false }}) ThenReturn(
{{ range $idx, $p := .Body.Returns }}
	{{printf "ret%d" $idx}} {{  $p.Expr }},
{{end}}
) {{ .Name }}Behaviour{{ .GenericExpr false }} {
	return c.ThenEffect(func(
		{{range .Body.Args}}
		{{ .Expr }},
		{{ end }}
		{{ if .Body.VarArg }}{{ .Body.VarArg.Expr }},{{ end }}
	){{ if .Body.Returns }}(
		{{ range .Body.Returns }}{{ .Expr }},
		{{end}}
	){{ end }}{
		{{ if .Body.Returns }}
		return {{ range $idx, $p := .Body.Returns }}{{ if ne $idx 0}},  {{ end }}{{printf "ret%d" $idx}}{{end}}
		{{ end }}
	})
}

func (c _{{ .Name }}Call{{ .GenericExpr false }}) ThenEffect(effect {{ .Body.Expr }}) {{ .Name }}Behaviour{{ .GenericExpr false }} {
	return {{ .Name }}Behaviour{{ .GenericExpr false }} {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}

{{ end }}

{{ range .Interfaces }}
type {{ .Name }}Impl{{ .GenericExpr true }} struct {
	{{ range .Body.Methods }}
	{{ .Name }} {{ .Func.Expr }}
	{{- end }}
}

func NewMocked{{ .Name }}{{ .GenericExpr true }}(t interface { Fatal(...any) } ,impl {{ .Name }}Impl{{ .GenericExpr false }}) testee.{{ .Name }}{{ .GenericExpr false }} {
	return _{{ .Name }}Mock{{ .GenericExpr false }} { t: t, impl: impl }
}

type _{{ .Name }}Mock{{ .GenericExpr true }} struct {
	t interface { Fatal(...any) }
	impl {{ .Name }}Impl{{ .GenericExpr false }}
}

{{- $i := . -}}
{{- range .Body.Methods }}

func (m _{{ $i.Name }}Mock{{ $i.GenericExpr false }}) {{ .Name }} ({{ range $idx, $param := .Func.Args }}
	{{ $param.ParamNameOr (printf "arg%d" $idx) }} {{ .Expr }},{{ end -}}
	{{- if .Func.VarArg }}
	{{ .Func.VarArg.ParamNameOr "vararg" }} {{ .Func.VarArg.Expr }},{{ end }}
){{ if .Func.Returns }} ({{ range $idx, $param := .Func.Returns }}
	{{ if.ParamName }}{{ .ParamName }} {{ end }}{{ .Expr }},{{ end }}
){{ end }} {
	cancel := itskit.SkipStack()
	defer cancel()

	if m.impl.{{ .Name }} == nil {
		itskit.NG(
			itskit.NewLabelWithLocation("{{ $i.Name }}.{{ .Name }} is not mocked").String(),
		).OrFatal(m.t)
	}

	{{ if .Func.Returns }}return {{ end }}m.impl.{{ .Name }}({{ range $idx, $param := .Func.Args }}
		{{ $param.ParamNameOr (printf "arg%d" $idx) }},{{ end -}}
		{{- if .Func.VarArg }}
		{{ .Func.VarArg.ParamNameOr "vararg" }}...,{{ end }}
	)
}
{{ end }}
{{ end }}
`
