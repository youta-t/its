package parser

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
	localPackage, err := detectPackage(dir, afile)
	if err != nil {
		return nil, err
	}

	pkgs := map[string]*Import{}
	for _, imp := range getImports(afile) {
		pkgs[imp.Name] = imp
	}
	log.Printf("packages: %+v", pkgs)

	decls, err := getDecls(localPackage, pkgs, afile)
	if err != nil {
		return nil, err
	}

	imports := []*Import{
		localPackage,
	}
	for k := range pkgs {
		imports = append(imports, pkgs[k])
	}

	return &File{
		PackageName: afile.Name.Name,
		Imports:     imports,
		Types:       decls,
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

func detectPackage(directory string, aFile *ast.File) (*Import, error) {
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
	var packageName string
	ast.Inspect(aFile, func(n ast.Node) bool {
		pkg, ok := n.(*ast.Package)
		if !ok {
			return true
		}
		packageName = pkg.Name
		return false
	})

	return &Import{
		Name: packageName,
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

type TypeDeclarations struct {
	Structs    []*TypeStructDecl
	Interfaces []*TypeInterfaceDecl
	Funcs      []*TypeFuncDecl
}

func getDecls(local *Import, imports map[string]*Import, n ast.Node) (*TypeDeclarations, error) {

	type structDef struct {
		name      string
		typeParam []*ast.Field
		body      *ast.StructType
	}
	structDefs := []structDef{}

	type interfaceDef struct {
		name      string
		typeParam []*ast.Field
		body      *ast.InterfaceType
	}
	interfaceDefs := []interfaceDef{}

	type funcDef struct {
		name      string
		typeParam []*ast.Field
		body      *ast.FuncType
	}
	fnDefs := []funcDef{}

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

			switch spec := typeSpec.Type.(type) {
			case *ast.StructType:
				if spec.Incomplete {
					continue
				}
				structDefs = append(structDefs, structDef{
					name:      safeDeref(typeSpec.Name).Name,
					typeParam: safeDeref(typeSpec.TypeParams).List,
					body:      spec,
				})
			case *ast.InterfaceType:
				if spec.Incomplete {
					continue
				}
				interfaceDefs = append(interfaceDefs, interfaceDef{
					name:      safeDeref(typeSpec.Name).Name,
					typeParam: safeDeref(typeSpec.TypeParams).List,
					body:      spec,
				})
			case *ast.FuncType:
				fnDefs = append(fnDefs, funcDef{
					name:      safeDeref(typeSpec.Name).Name,
					typeParam: safeDeref(typeSpec.TypeParams).List,
					body:      spec,
				})
			}

		}
		return false
	})

	structs := []*TypeStructDecl{}
	for _, d := range structDefs {
		tps, err := parseTypeParam(local, imports, d.typeParam)
		if err != nil {
			return nil, err
		}

		body, err := parseStruct(local, imports, tps, d.body)
		if err != nil {
			return nil, err
		}

		structs = append(structs, &TypeStructDecl{
			Name:       d.name,
			Package:    local,
			TypeParams: tps,
			Body:       body,
		})
	}

	interfaces := []*TypeInterfaceDecl{}
	for _, d := range interfaceDefs {
		tps, err := parseTypeParam(local, imports, d.typeParam)
		if err != nil {
			return nil, err
		}
		body, err := parseInterface(local, imports, tps, d.body)
		if err != nil {
			return nil, err
		}
		interfaces = append(interfaces, &TypeInterfaceDecl{
			Name:       d.name,
			Package:    local,
			TypeParams: tps,
			Body:       body,
		})
	}

	funcs := []*TypeFuncDecl{}
	for _, d := range fnDefs {
		tps, err := parseTypeParam(local, imports, d.typeParam)
		if err != nil {
			return nil, err
		}
		body, err := parseFn(local, imports, tps, d.body)
		if err != nil {
			return nil, err
		}
		funcs = append(funcs, &TypeFuncDecl{
			Name:       d.name,
			Package:    local,
			TypeParams: tps,
			Body:       body,
		})
	}

	return &TypeDeclarations{
		Structs:    structs,
		Interfaces: interfaces,
		Funcs:      funcs,
	}, nil
}

func parseTypeParam(local *Import, pkgs map[string]*Import, params []*ast.Field) ([]*TypeParam, error) {
	tps := []*TypeParam{}
	for _, tp := range params {
		constraint, err := parseType(local, pkgs, nil, tp.Type)
		if err != nil {
			return nil, err
		}
		for i := range tp.Names {
			tps = append(tps, &TypeParam{
				Name: tp.Names[i].Name,
				Back: constraint,
			})
		}
	}

	for i := range tps {
		tps[i].Back.injectTypeParam(local, tps)
	}

	return tps, nil
}

func resolveBareNameType(local *Import, tps []*TypeParam, name string) Type {
	for _, tp := range tps {
		if tp.Name == name {
			return tp
		}
	}
	switch name {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "complex64", "complex128",
		"bool", "rune", "byte", "uintptr",
		"string", "error", "any":
		return &BuiltinType{Name: name}
	default:
		return &NamedType{Pkg: local, Name: name}
	}
}

func parseType(local *Import, imports map[string]*Import, tps []*TypeParam, node ast.Expr) (Type, error) {
	switch t := node.(type) {
	case *ast.Ident: // local or built-in type
		if tps == nil { // type params are unresolved
			return &pseudoType{Name: t.Name}, nil
		}
		return resolveBareNameType(local, tps, t.Name), nil

	case *ast.SelectorExpr: // imported type
		x, ok := t.X.(*ast.Ident)
		if !ok {
			return nil, errors.New("selector is not normal name")
		}
		return &NamedType{Pkg: imports[x.Name], Name: t.Sel.Name}, nil

	case *ast.IndexExpr: // generics
		hostType, err := parseType(local, imports, tps, t.X)
		if err != nil {
			return nil, err
		}
		tp, err := parseType(local, imports, tps, t.Index)
		if err != nil {
			return nil, err
		}
		return &GenericType{Host: hostType, Params: []Type{tp}}, nil
	case *ast.IndexListExpr: // generics
		hostType, err := parseType(local, imports, tps, t.X)
		if err != nil {
			return nil, err
		}
		params := []Type{}
		for i := range t.Indices {
			tp, err := parseType(local, imports, tps, t.Indices[i])
			if err != nil {
				return nil, err
			}
			params = append(params, tp)
		}
		return &GenericType{Host: hostType, Params: params}, nil
	case *ast.StarExpr: // pointer
		elem, err := parseType(local, imports, tps, t.X)
		if err != nil {
			return nil, err
		}
		return &PointerType{Elem: elem}, nil

	case *ast.ArrayType: // array or slice
		elem, err := parseType(local, imports, tps, t.Elt)
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
		key, err := parseType(local, imports, tps, t.Key)
		if err != nil {
			return nil, err
		}
		val, err := parseType(local, imports, tps, t.Value)
		if err != nil {
			return nil, err
		}
		return &MapType{Key: key, Elem: val}, nil

	case *ast.ChanType: // channel
		elem, err := parseType(local, imports, tps, t.Value)
		if err != nil {
			return nil, err
		}
		return &ChanType{Elem: elem, Dir: t.Dir}, nil

	case *ast.FuncType: // func
		return parseFn(local, imports, tps, t)

	case *ast.StructType: // struct literal
		return parseStruct(local, imports, tps, t)

	case *ast.InterfaceType: // interface literal
		return parseInterface(local, imports, tps, t)
	case *ast.UnaryExpr:
		tt, err := parseType(local, imports, tps, t.X)
		if err != nil {
			return nil, err
		}
		if t.Op.String() != "~" {
			return nil, fmt.Errorf("unknown type unaryop: %s", t.Op.String())
		}
		return &TypeConstraint{
			Op:   t.Op.String(),
			Type: tt,
		}, nil
	case *ast.BinaryExpr:
		tx, err := parseType(local, imports, tps, t.X)
		if err != nil {
			return nil, err
		}
		ty, err := parseType(local, imports, tps, t.Y)
		if err != nil {
			return nil, err
		}
		if t.Op.String() != "|" {
			return nil, fmt.Errorf("unknown type binaryop: %s", t.Op.String())
		}
		return &TypeUnion{
			Op: t.Op.String(),
			X:  tx,
			Y:  ty,
		}, nil

	}
	return nil, fmt.Errorf("unknown type reference: %T", node)
}

func parseStruct(local *Import, imports map[string]*Import, typeParams []*TypeParam, snode *ast.StructType) (*StructType, error) {
	fields := []*Field{}
	for _, f := range snode.Fields.List {
		typ, err := parseType(local, imports, typeParams, f.Type)
		if err != nil {
			return nil, err
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

func parseInterface(local *Import, imports map[string]*Import, typeParams []*TypeParam, node *ast.InterfaceType) (*InterfaceType, error) {
	methods := []*Method{}
	for _, f := range node.Methods.List {
		if m, ok := f.Type.(*ast.FuncType); ok {
			typ, err := parseFn(local, imports, typeParams, m)
			if err != nil {
				return nil, err
			}
			if len(f.Names) == 0 {
				// embedded
				methods = append(methods, &Method{Name: "", Func: typ})
			} else {
				for _, n := range f.Names {
					methods = append(methods, &Method{Name: n.Name, Func: typ})
				}
			}
		}
	}
	return &InterfaceType{Methods: methods}, nil
}

func parseFn(local *Import, imports map[string]*Import, tps []*TypeParam, fnode *ast.FuncType) (*FuncType, error) {
	params := []*FuncIOParam{}
	var ellip *FuncIOParam
	for _, param := range fnode.Params.List {
		names := []string{}
		for _, n := range param.Names {
			names = append(names, n.Name)
		}
		if len(names) == 0 {
			names = append(names, "")
		}
		for _, name := range names {
			switch p := param.Type.(type) {
			case *ast.Ellipsis:
				e, err := parseType(local, imports, tps, p.Elt)
				if err != nil {
					return nil, err
				}
				for _, tp := range tps {
					if tp.Name == e.PlainName() {
						e = tp
						break
					}
				}
				ellip = &FuncIOParam{
					ParamName: name,
					Type:      e,
					Variadic:  true,
				}
			default:
				e, err := parseType(local, imports, tps, p)
				if err != nil {
					return nil, err
				}
				for _, tp := range tps {
					if tp.Name == e.PlainName() {
						e = tp
						break
					}
				}
				params = append(params, &FuncIOParam{
					ParamName: name,
					Type:      e,
				})
			}
		}
	}

	returns := []*FuncIOParam{}
	if fnode.Results != nil {
		for _, result := range fnode.Results.List {
			r, err := parseType(local, imports, tps, result.Type)
			if err != nil {
				return nil, err
			}
			for _, tp := range tps {
				if tp.Name == r.PlainName() {
					r = tp
					break
				}
			}
			if names := result.Names; len(names) == 0 {
				returns = append(returns, &FuncIOParam{
					ParamName: "",
					Type:      r,
				})
			} else {
				for _, name := range names {
					returns = append(returns, &FuncIOParam{
						ParamName: name.Name,
						Type:      r,
					})
				}
			}
		}
	}

	return &FuncType{
		Args: params, VarArg: ellip, Returns: returns,
	}, nil
}

func safeDeref[T any](val *T) T {
	if val == nil {
		return *new(T)
	}
	return *val
}
