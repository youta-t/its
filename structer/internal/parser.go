package internal

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/mod/modfile"
)

func Parse(filepath string) (*File, error) {
	fset := token.NewFileSet()
	afile, err := parser.ParseFile(fset, filepath, nil, parser.Mode(0))
	if err != nil {
		return nil, err
	}

	typenames := map[string]struct{}{}
	for _, tn := range flag.Args() {
		typenames[tn] = struct{}{}
	}

	dir := path.Dir(filepath)
	if dir == "" {
		dir = "."
	}
	targetPackage, err := detectPackage(dir)
	if err != nil {
		return nil, err
	}

	pkgs := map[string]*Import{
		"": targetPackage,
	}
	for _, imp := range getImports(afile) {
		pkgs[imp.Name] = imp
	}
	log.Printf("packages: %+v", pkgs)

	structs, err := getStructDecls(pkgs, afile)
	if err != nil {
		return nil, err
	}

	imports := []*Import{}
	for k := range pkgs {
		imports = append(imports, pkgs[k])
	}

	return &File{
		PackageName: afile.Name.Name,
		Imports:     imports,
		Structs:     structs,
	}, nil
}

func detectGomod(from string) (string, *modfile.File, error) {
	gomod := path.Join(from, "go.mod")
	buf, err := os.ReadFile(gomod)

	if err == nil {
		m, err := modfile.Parse(gomod, buf, nil)
		return from, m, err
	} else if !os.IsNotExist(err) {
		return from, nil, err
	}

	p := path.Dir(from)
	if p == from {
		return from, nil, err
	}
	return detectGomod(p)
}

func detectPackage(directory string) (*Import, error) {
	directory, err := filepath.Abs(directory)
	if err != nil {
		return nil, err
	}
	directory = strings.TrimSuffix(directory, "/")
	root, gomod, err := detectGomod(directory)
	if err != nil {
		return nil, err
	}
	pkgpath, err := filepath.Rel(root, directory)
	if err != nil {
		return nil, err
	}
	pkgroot := gomod.Module.Mod.Path

	return &Import{
		Name: "",
		Path: path.Join(pkgroot, pkgpath),
	}, nil
}

func getImports(n ast.Node) []*Import {
	ret := []*Import{}
	ast.Inspect(n, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}
		if decl.Tok != token.IMPORT {
			return false
		}
		for _, s := range decl.Specs {
			imp, ok := s.(*ast.ImportSpec)
			if !ok {
				continue
			}
			name := ""
			if n := imp.Name; n != nil {
				name = n.Name
			}
			importPath := imp.Path.Value
			importPath = strings.Trim(importPath, "`\"")
			if name == "" {
				name = path.Base(importPath)
			}
			ret = append(ret, &Import{
				Name: name, Path: importPath,
			})
		}
		return false
	})
	return ret
}

func getStructDecls(imports map[string]*Import, n ast.Node) ([]*StructDecl, error) {

	type defn struct {
		name      string
		typeParam []*ast.Field
		body      *ast.StructType
	}

	defs := []defn{}

	ast.Inspect(n, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl, *ast.BadDecl:
			return false
		}

		decl, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}
		if decl.Tok != token.TYPE {
			return false
		}
		for i := range decl.Specs {
			spec := decl.Specs[i]
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok || !typeSpec.Name.IsExported() {
				continue
			}
			snode, ok := typeSpec.Type.(*ast.StructType)
			if !ok || snode.Incomplete {
				continue
			}

			defs = append(defs, defn{
				name:      safeDeref(typeSpec.Name).Name,
				typeParam: safeDeref(typeSpec.TypeParams).List,
				body:      snode,
			})
		}
		return false
	})

	ret := []*StructDecl{}
	for _, d := range defs {
		tps := []*TypeParam{}
		for _, tp := range d.typeParam {
			back, err := parseType(imports, nil, tp.Type)
			if err != nil {
				return nil, err
			}
			for _, n := range tp.Names {
				tps = append(tps, &TypeParam{Name: n.Name, Back: back})
			}
		}

		body, err := parseStruct(imports, tps, d.body)
		if err != nil {
			return nil, err
		}

		ret = append(ret, &StructDecl{
			Name:       d.name,
			Package:    imports[""],
			TypeParams: tps,
			Body:       body,
		})
	}

	return ret, nil
}

func parseType(imports map[string]*Import, tps []*TypeParam, node ast.Expr) (Type, error) {
	switch t := node.(type) {
	case *ast.Ident: // local or built-in type
		switch t.Name {
		case "int", "int8", "int32", "int64",
			"uint", "uint8", "uint32", "uint64",
			"float32", "float64", "complex64", "complex128",
			"bool", "rune", "byte", "uintptr",
			"string", "error", "any":
			return &BuiltinType{Name: t.Name}, nil
		default:
			return &NamedType{Pkg: imports[""], Name: t.Name}, nil
		}

	case *ast.SelectorExpr: // imported type
		x, ok := t.X.(*ast.Ident)
		if !ok {
			return nil, errors.New("selector is not normal name")
		}
		return &NamedType{Pkg: imports[x.Name], Name: t.Sel.Name}, nil

	case *ast.IndexExpr: // generics
		hostType, err := parseType(imports, tps, t.X)
		if err != nil {
			return nil, err
		}
		tp, err := parseType(imports, tps, t.Index)
		if err != nil {
			return nil, err
		}
		return &GenericType{Host: hostType, Params: []Type{tp}}, nil
	case *ast.IndexListExpr: // generics
		hostType, err := parseType(imports, tps, t.X)
		if err != nil {
			return nil, err
		}
		params := []Type{}
		for i := range t.Indices {
			tp, err := parseType(imports, tps, t.Indices[i])
			if err != nil {
				return nil, err
			}
			params = append(params, tp)
		}
		return &GenericType{Host: hostType, Params: params}, nil
	case *ast.StarExpr: // pointer
		elem, err := parseType(imports, tps, t.X)
		if err != nil {
			return nil, err
		}
		return &PointerType{Elem: elem}, nil

	case *ast.ArrayType: // array or slice
		elem, err := parseType(imports, tps, t.Elt)
		if err != nil {
			return nil, err
		}
		if l := t.Len; l == nil {
			return &SliceType{Elem: elem}, nil
		} else {
			id, ok := l.(*ast.BasicLit)
			if !ok {
				return nil, errors.New("array length is not literal")
			}
			l, err := strconv.Atoi(id.Value)
			if err != nil {
				return nil, err
			}
			return &ArrayType{Elem: elem, Len: l}, nil
		}

	case *ast.MapType:
		key, err := parseType(imports, tps, t.Key)
		if err != nil {
			return nil, err
		}
		val, err := parseType(imports, tps, t.Value)
		if err != nil {
			return nil, err
		}
		return &MapType{Key: key, Elem: val}, nil

	case *ast.ChanType: // channel
		elem, err := parseType(imports, tps, t.Value)
		if err != nil {
			return nil, err
		}
		return &ChanType{Elem: elem, Dir: t.Dir}, nil

	case *ast.FuncType: // func
		return parseFn(imports, tps, t)

	case *ast.StructType: // struct literal
		return parseStruct(imports, tps, t)

	case *ast.InterfaceType: // interface literal
		methods := []*Method{}
		for _, f := range t.Methods.List {
			if m, ok := f.Type.(*ast.FuncType); ok {
				typ, err := parseFn(imports, tps, m)
				if err != nil {
					return nil, err
				}
				if len(f.Names) == 0 {
					// embedded
					methods = append(methods, &Method{Name: "", Signature: typ})
				} else {
					for _, n := range f.Names {
						methods = append(methods, &Method{Name: n.Name, Signature: typ})
					}
				}
			}
		}
		return &InterfaceType{Methods: methods}, nil
	}
	return nil, fmt.Errorf("unknown type reference: %T", node)
}

func parseStruct(imports map[string]*Import, typeParams []*TypeParam, snode *ast.StructType) (*StructType, error) {
	fields := []*Field{}
	for _, f := range snode.Fields.List {
		typ, err := parseType(imports, typeParams, f.Type)
		if err != nil {
			return nil, err
		}

		if nametype, ok := typ.(*NamedType); ok && nametype.Pkg.Name == "" {
			for n := range typeParams {
				tp := typeParams[n]
				if tp.Name == nametype.Name {
					typ = tp
				}
			}
		} else if builtin, ok := typ.(*BuiltinType); ok {
			for n := range typeParams {
				tp := typeParams[n]
				if tp.Name == builtin.Name {
					typ = tp
				}
			}
		}

		if len(f.Names) == 0 {
			// embedded!
			fields = append(fields, &Field{Name: typ.PlainName(), Type: typ})
		} else {
			for _, n := range f.Names {
				fields = append(fields, &Field{Name: n.Name, Type: typ})
			}
		}
	}
	return &StructType{Fields: fields}, nil

}

func parseFn(imports map[string]*Import, tps []*TypeParam, fnode *ast.FuncType) (*FuncType, error) {
	params := []Type{}
	var ellip Type
	for _, param := range fnode.Params.List {
		switch p := param.Type.(type) {
		case *ast.Ellipsis:
			e, err := parseType(imports, tps, p.Elt)
			if err != nil {
				return nil, err
			}
			ellip = e
		default:
			e, err := parseType(imports, tps, p)
			if err != nil {
				return nil, err
			}
			if names := param.Names; len(names) == 0 {
				params = append(params, e)
			} else {
				for range names {
					params = append(params, e)
				}
			}
		}
	}

	returns := []Type{}
	for _, result := range fnode.Results.List {
		r, err := parseType(imports, tps, result.Type)
		if err != nil {
			return nil, err
		}
		if names := result.Names; len(names) == 0 {
			returns = append(returns, r)
		} else {
			for range names {
				returns = append(returns, r)
			}
		}
	}

	return &FuncType{
		Params: params, EllipsisParam: ellip, Results: returns,
	}, nil
}

func safeDeref[T any](val *T) T {
	if val == nil {
		return *new(T)
	}
	return *val
}
