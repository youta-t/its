package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
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
	Imports     *parser.Imports
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

	interfaces := map[string][]*parser.TypeInterfaceDecl{}
	functions := map[string][]*parser.TypeFuncDecl{}
	filenames := map[string]struct{}{}

	{
		funcs := pkg.Types.Funcs.Slice()
		for i := range funcs {
			def := funcs[i]
			if targetFile != "" && def.DefinedIn != targetFile {
				continue
			}
			if _, ok := targetTypeName[def.Name]; len(targetTypeName) != 0 && !ok {
				continue
			}
			filenames[def.DefinedIn] = struct{}{}
			functions[def.DefinedIn] = append(functions[def.DefinedIn], def)
		}
	}

	{
		intfs := pkg.Types.Interfaces.Slice()
		for i := range intfs {
			def := intfs[i]

			if targetFile != "" && def.DefinedIn != targetFile {
				continue
			}
			if _, ok := targetTypeName[def.Name]; len(targetTypeName) != 0 && !ok {
				continue
			}

			def.Body = try.To(def.Body.Inlined(parserInstance)).OrFatal(logger)

			filenames[def.DefinedIn] = struct{}{}
			interfaces[def.DefinedIn] = append(interfaces[def.DefinedIn], def)

			for j := range def.Body.Methods {
				meth := def.Body.Methods[j]
				functions[def.DefinedIn] = append(functions[def.DefinedIn], &parser.TypeFuncDecl{
					DefinedIn:  def.DefinedIn,
					ImportPath: def.ImportPath,
					Name:       fmt.Sprintf("%s_%s", def.Name, meth.Name),
					TypeParams: def.TypeParams,
					Body:       meth.Func,
				})
			}
		}
	}

	for fname := range filenames {
		newFile := generatingFile{
			PackageName: path.Base(dest),
			Imports:     new(parser.Imports),
		}

		funcs := functions[fname]
		intfs := interfaces[fname]

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
				newFile.Imports.Add(t)
			}
		}

		for i := range intfs {
			s := intfs[i]

			if _, ok := targetTypeName[s.Name]; len(targetTypeName) != 0 && !ok {
				continue
			}
			if s.IsOpaque() || 0 < len(s.Body.Embedded) {
				continue
			}

			newFile.Interfaces = append(newFile.Interfaces, s)
			types := s.Require()
			newFile.Imports.Add(s.ImportPath)
			for i := range types {
				t := types[i]
				newFile.Imports.Add(t)
			}
		}

		if len(newFile.Interfaces)+len(newFile.Funcs) == 0 {
			continue
		}

		if err := writeFile(filepath.Join(dest, filepath.Base(fname)), newFile); err != nil {
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
	its "github.com/youta-t/its"
	itskit "github.com/youta-t/its/itskit"
	mockkit "github.com/youta-t/its/mocker/mockkit"
	{{ range .Imports.Slice }}
	{{- .Name }} "{{ .Path }}"
	{{ end }}
)
{{ $imports := .Imports }}
{{- range .Funcs -}}
{{- $func := . }}
type _{{ .Name }}CallSpec{{ .GenericExpr $imports true }} struct {
	{{- range $idx, $p := .Body.Args}}
	{{ $p.ParamNameOr (printf "arg%d" $idx) }} its.Matcher[{{ $p.Expr $imports }}]
	{{ end }}
	{{ if .Body.VarArg }}{{ .Body.VarArg.ParamNameOr "vararg" }} its.Matcher[[]{{ .Body.VarArg.Type.Expr $imports }}]{{ end }}
}

type _{{ .Name }}Call{{ .GenericExpr $imports true }} struct {
	name itskit.Label
	spec _{{ .Name }}CallSpec{{ .GenericExpr $imports false }}
}

func {{ .Name }}_Expects{{ .GenericExpr $imports true }}(
	{{- range $idx, $p := .Body.Args}}
	{{ $p.ParamNameOr (printf "arg%d" $idx) }} its.Matcher[{{ $p.Expr $imports }}],
	{{ end -}}
	{{ if .Body.VarArg }}{{ .Body.VarArg.ParamNameOr "vararg" }} its.Matcher[[]{{ .Body.VarArg.Type.Expr $imports }}],{{ end }}
) _{{ .Name }}Call{{ .GenericExpr $imports false }} {
	cancel := itskit.SkipStack()
	defer cancel()

	spec := _{{ .Name }}CallSpec{{ .GenericExpr $imports false }} {}
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
	return _{{ .Name }}Call{{ .GenericExpr $imports false }}{
		name: itskit.NewLabelWithLocation("func {{ .Name }}"),
		spec: spec,
	}
}

type _{{ .Name }}Behavior {{ .GenericExpr $imports true }} struct {
	name itskit.Label
	spec _{{ .Name }}CallSpec{{ .GenericExpr $imports false }}
	effect {{ .Body.Expr $imports }}
}

func (b *_{{ .Name }}Behavior{{ .GenericExpr $imports false }}) Fn(t mockkit.TestLike) {{ .Body.Expr $imports }} {
	return func (
		{{ range $idx, $p :=  .Body.Args }}
		{{ printf "arg%d" $idx }} {{ $p.Expr $imports }},
		{{ end }}
		{{ if .Body.VarArg }}vararg {{ .Body.VarArg.Expr $imports }},{{ end }}
	) {{ if .Body.Returns }}(
		{{ range .Body.Returns }}{{ .Expr $imports }},
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
				matcher = its.Never[{{ $p.Expr $imports }}]()
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
				matcher = its.Never[[]{{ .Body.VarArg.Type.Expr $imports }}]()
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

func (c _{{.Name}}Call{{ .GenericExpr $imports false }}) ThenReturn(
{{ range $idx, $p := .Body.Returns }}
	{{printf "ret%d" $idx}} {{  $p.Expr $imports }},
{{end}}
) mockkit.FuncBehavior[ func {{ .Body.Signature $imports false }}  ] {
	return c.ThenEffect(func(
		{{range .Body.Args}}
		{{ .Expr $imports }},
		{{ end }}
		{{ if .Body.VarArg }}{{ .Body.VarArg.Expr $imports }},{{ end }}
	){{ if .Body.Returns }}(
		{{ range .Body.Returns }}{{ .Expr $imports }},
		{{end}}
	){{ end }}{
		{{ if .Body.Returns }}
		return {{ range $idx, $p := .Body.Returns }}{{ if ne $idx 0}},  {{ end }}{{printf "ret%d" $idx}}{{end}}
		{{ end }}
	})
}

func (c _{{ .Name }}Call{{ .GenericExpr $imports false }}) ThenEffect(effect {{ .Body.Expr $imports }}) mockkit.FuncBehavior[ func {{ .Body.Signature $imports false }} ] {
	return &_{{ .Name }}Behavior{{ .GenericExpr $imports false }} {
		name: c.name,
		spec: c.spec,
		effect: effect,
	}
}

{{ end }}

{{ range .Interfaces }}
type _{{ .Name }}Impl{{ .GenericExpr $imports true }} struct {
	{{ range .Body.Methods }}
	{{ .Name }} {{ .Func.Expr $imports }}
	{{- end }}
}

func {{ .Name }}_Build{{ .GenericExpr $imports true }}(t mockkit.TestLike, spec {{ .Name }}_Spec{{ .GenericExpr $imports false }}) {{ .Expr $imports }} {
	impl := _{{ .Name }}Impl{{ .GenericExpr $imports false }}{}

	{{ range .Body.Methods }}
	if spec.{{ .Name }} != nil {
		impl.{{ .Name }} = spec.{{ .Name }}.Fn(t)
	}
	{{ end }}

	return _{{ .Name }}Mock{{ .GenericExpr $imports false }} { t: t, impl: impl }
}

type _{{ .Name }}Mock{{ .GenericExpr $imports true }} struct {
	t mockkit.TestLike
	impl _{{ .Name }}Impl{{ .GenericExpr $imports false }}
}

{{- $i := . -}}
{{- range .Body.Methods }}

func (m _{{ $i.Name }}Mock{{ $i.GenericExpr $imports false }}) {{ .Name }} ({{ range $idx, $param := .Func.Args }}
	{{ $param.ParamNameOr (printf "arg%d" $idx) }} {{ .Expr $imports }},{{ end -}}
	{{- if .Func.VarArg }}
	{{ .Func.VarArg.ParamNameOr "vararg" }} {{ .Func.VarArg.Expr $imports }},{{ end }}
){{ if .Func.Returns }} ({{ range $idx, $param := .Func.Returns }}
	{{ if.ParamName }}{{ .ParamName }} {{ end }}{{ .Expr $imports }},{{ end }}
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

type {{ $i.Name }}_Spec{{ $i.GenericExpr $imports true }} struct {
	{{ range .Body.Methods }}
	{{ .Name }} mockkit.FuncBehavior[func {{ .Func.Signature $imports false }}]
	{{ end }}
}
{{ end }}
`
