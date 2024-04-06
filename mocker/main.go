package main

import (
	"flag"
	"fmt"
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
	Imports     []*parser.ImportPath
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

	gomod := parser.New(".")

	var pkg *parser.Package
	targetFile := ""

	if *sourceAsPackage {
		pkg = try.To(gomod.Import(source, ".")).OrFatal(logger)
	} else {
		dir, t := filepath.Split(source)
		targetFile = t
		pkg = try.To(gomod.ImportDir(dir)).OrFatal(logger)
	}

	target := pkg.Types
	if targetFile != "" {
		f, ok := pkg.Types[targetFile]
		if !ok {
			os.Exit(0) // nothing to do.
		}

		target = map[string]*parser.File{targetFile: f}
	}

	for fname, content := range target {
		newFile := generatingFile{
			PackageName: path.Base(dest),
		}
		requiredImports := map[string]*parser.ImportPath{
			"its":     {Name: "its", Path: "github.com/youta-t/its"},
			"itskit":  {Name: "itskit", Path: "github.com/youta-t/its/itskit"},
			"mockkit": {Name: "mockkit", Path: "github.com/youta-t/its/mocker/mockkit"},
		}
		usedImports := map[string]*parser.ImportPath{}

		funcs := content.Types.Funcs.Slice()
		for i := range funcs {
			s := funcs[i]

			if _, ok := targetTypeName[s.Name]; len(targetTypeName) != 0 && !ok {
				continue
			}
			if s.IsOpaque() {
				continue
			}

			newFile.Funcs = append(newFile.Funcs, s)

			types := s.Require()
			for i := range types {
				t := types[i]
				usedImports[t.Name] = t
			}
		}

		interfaces := content.Types.Interfaces.Slice()
		for i := range interfaces {
			s := interfaces[i]

			if _, ok := targetTypeName[s.Name]; len(targetTypeName) != 0 && !ok {
				continue
			}
			if s.IsOpaque() || 0 < len(s.Body.Embedded) {
				continue
			}

			newFile.Interfaces = append(newFile.Interfaces, s)

			types := s.Require()
			for i := range types {
				t := types[i]
				usedImports[t.Name] = t
			}

			for _, m := range s.Body.Methods {
				fd := &parser.TypeFuncDecl{
					Name:       fmt.Sprintf("%s_%s", s.Name, m.Name),
					Package:    s.Package,
					TypeParams: m.Func.TypeParams(),
					Body:       m.Func,
				}
				funcs = append(funcs, fd)
			}
		}

		if len(newFile.Interfaces)+len(newFile.Funcs) == 0 {
			continue
		}

		for i := range requiredImports {
			newFile.Imports = append(newFile.Imports, requiredImports[i])
		}
		slices.SortFunc(newFile.Imports, func(a, b *parser.ImportPath) int {
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

		if err := writeFile(filepath.Join(dest, fname), newFile); err != nil {
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
					expr += " " + p.Constraint.Expr()
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

func inlineInterface(bc parser.BuildContext, intf *parser.InterfaceType) (*parser.InterfaceType, error) {
	embededds := []parser.Type{}
	methods := append([]*parser.Method{}, intf.Methods...)

	type namedTypeInstance struct {
		Type           *parser.NamedType
		GivenTypeParam []parser.Type
	}

	embeddedNames := []*parser.NamedType{}
	embeddedInterfaces := []*parser.InterfaceType{}
	for i := range intf.Embedded {
		switch t := intf.Embedded[i].(type) {
		case *parser.NamedType:
			embeddedNames = append(embeddedNames, t)
		case *parser.InterfaceType:
			ii, err := inlineInterface(bc, t)
			if err != nil {
				return nil, err
			}
			embeddedInterfaces = append(embeddedInterfaces, ii)
		default:
			embededds = append(embededds, t)
		}
	}

	for 0 < len(embeddedNames) {
		n := embeddedNames[0]
		embeddedNames = embeddedNames[1:]

		pkg, err := bc.Import(n.Pkg.Path, ".")
		if err != nil {
			return nil, err
		}

		for _, f := range pkg.Types {
			if found, ok := f.Types.Interfaces.Get(n.PlainName()); ok {
				b := found.Body

				b, err := inlineInterface(bc, found.Body)
				if err != nil {
					return nil, err
				}
				embeddedInterfaces = append(embeddedInterfaces, b)
				break
			}

			nn, ok := f.Types.Names.Get(n.PlainName())
			if !ok {
				continue
			}
			nt := &parser.NamedType{
				Pkg:    nn.Body.Pkg,
				Name:   nn.Body.Name,
				Params: n.Params,
			}
			embeddedNames = append(embeddedNames, nt)
		}
	}

	for i := range embeddedInterfaces {
		methods = append(methods, embeddedInterfaces[i].Methods...)
		embededds = append(embededds, embeddedInterfaces[i].Embedded...)
	}

	return &parser.InterfaceType{
		Embedded: embededds,
		Methods:  methods,
	}, nil
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
package {{ .PackageName }}

import (
	{{ range .Imports }}
	{{- .Name }} "{{ .Path }}"
	{{ end }}
)

{{- range .Funcs -}}
{{- $func := . }}
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

func {{ .Name }}_Expects{{ .GenericExpr true }}(
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

type _{{ .Name }}Behavior {{ .GenericExpr true }} struct {
	name itskit.Label
	spec _{{ .Name }}CallSpec{{ .GenericExpr false }}
	effect {{ .Body.Expr }}
}

func (b *_{{ .Name }}Behavior{{ .GenericExpr false }}) Fn(t mockkit.TestLike) {{ .Body.Expr }} {
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
) mockkit.FuncBehavior[ func {{ .Body.Signature false }}  ] {
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

func (c _{{ .Name }}Call{{ .GenericExpr false }}) ThenEffect(effect {{ .Body.Expr }}) mockkit.FuncBehavior[ func {{ .Body.Signature false }} ] {
	return &_{{ .Name }}Behavior{{ .GenericExpr false }} {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}

{{ end }}

{{ range .Interfaces }}
type _{{ .Name }}Impl{{ .GenericExpr true }} struct {
	{{ range .Body.Methods }}
	{{ .Name }} {{ .Func.Expr }}
	{{- end }}
}

func {{ .Name }}_Build{{ .GenericExpr true }}(t mockkit.TestLike, spec {{ .Name }}_Spec{{ .GenericExpr false }}) testee.{{ .Name }}{{ .GenericExpr false }} {
	impl := _{{ .Name }}Impl{{ .GenericExpr false }}{}

	{{ range .Body.Methods }}
	if spec.{{ .Name }} != nil {
		impl.{{ .Name }} = spec.{{ .Name }}.Fn(t)
	}
	{{ end }}

	return _{{ .Name }}Mock{{ .GenericExpr false }} { t: t, impl: impl }
}

type _{{ .Name }}Mock{{ .GenericExpr true }} struct {
	t mockkit.TestLike
	impl _{{ .Name }}Impl{{ .GenericExpr false }}
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

type {{ $i.Name }}_Spec{{ $i.GenericExpr true }} struct {
	{{ range .Body.Methods }}
	{{ .Name }} mockkit.FuncBehavior[func {{ .Func.Signature false }}]
	{{ end }}
}
{{ end }}
`
