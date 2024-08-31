package parser

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	maps "github.com/youta-t/its/internal/maps"
)

func parsePackage(bc ParseContext, pkg importTarget) (*Package, error) {
	p := &Package{
		Path: pkg.ImportPath,
		Src:  pkg.Dir,
		Types: &TypeDeclarations{
			Structs:    maps.NewOrdered[string, *TypeDecl[*StructType]](),
			Interfaces: maps.NewOrdered[string, *TypeDecl[*InterfaceType]](),
			Funcs:      maps.NewOrdered[string, *TypeDecl[*FuncType]](),
			Maps:       maps.NewOrdered[string, *TypeDecl[*MapType]](),
			Slices:     maps.NewOrdered[string, *TypeDecl[*SliceType]](),
			Arrays:     maps.NewOrdered[string, *TypeDecl[*ArrayType]](),
			Names:      maps.NewOrdered[string, *TypeDecl[*NamedType]](),
			Unresolved: maps.NewOrdered[string, *TypeDecl[*unknwonType]](),
		},
	}

	bpkg, err := build.Import(pkg.ImportPath, ".", 0)
	if err != nil {
		return nil, err
	}
	p.DefaultName = bpkg.Name
	for _, gof := range bpkg.GoFiles {
		basename := filepath.Base(gof)
		f, err := parseFile(bc, pkg, basename)
		if err != nil {
			return nil, err
		}
		p.Types.Merge(f)
	}

	types := map[string]Type{}
	p.Types.Structs.Iter(func(s string, decl *TypeDecl[*StructType]) bool {
		types[s] = &NamedType{ImportPath: p.Path, Name: s}
		return true
	})
	p.Types.Interfaces.Iter(func(s string, decl *TypeDecl[*InterfaceType]) bool {
		types[s] = &NamedType{ImportPath: p.Path, Name: s}
		return true
	})
	p.Types.Funcs.Iter(func(s string, decl *TypeDecl[*FuncType]) bool {
		types[s] = &NamedType{ImportPath: p.Path, Name: s}
		return true
	})
	p.Types.Names.Iter(func(s string, decl *TypeDecl[*NamedType]) bool {
		types[s] = &NamedType{ImportPath: p.Path, Name: s}
		return true
	})
	p.Types.Maps.Iter(func(s string, decl *TypeDecl[*MapType]) bool {
		types[s] = &NamedType{ImportPath: p.Path, Name: s}
		return true
	})
	p.Types.Slices.Iter(func(s string, decl *TypeDecl[*SliceType]) bool {
		types[s] = &NamedType{ImportPath: p.Path, Name: s}
		return true
	})
	p.Types.Arrays.Iter(func(s string, decl *TypeDecl[*ArrayType]) bool {
		types[s] = &NamedType{ImportPath: p.Path, Name: s}
		return true
	})
	p.Types.Unresolved.Iter(func(s string, decl *TypeDecl[*unknwonType]) bool {
		types[s] = &NamedType{ImportPath: p.Path, Name: s}
		return true
	})

	{
		urs := p.Types.Unresolved.Slice()
		for i := range urs {
			ur := urs[i]
			det := ur.Body.detect(types)
			if unk, ok := det.(*unknwonType); ok {
				det = unk.detect(builtin)
			}
			if nt, ok := det.(*NamedType); ok {
				p.Types.Unresolved.Delete(ur.Name)
				p.Types.Names.Put(ur.Name, &TypeDecl[*NamedType]{
					DefinedIn: ur.DefinedIn, ImportPath: ur.ImportPath,
					Name: ur.Name, TypeParams: ur.TypeParams, Body: nt,
				})
			}
		}
	}

	p.Types.Structs.Iter(func(s string, decl *TypeDecl[*StructType]) bool {
		decl.resolve(types)
		decl.resolve(builtin)
		return true
	})
	p.Types.Interfaces.Iter(func(s string, decl *TypeDecl[*InterfaceType]) bool {
		decl.resolve(types)
		decl.resolve(builtin)
		return true
	})
	p.Types.Funcs.Iter(func(s string, decl *TypeDecl[*FuncType]) bool {
		decl.resolve(types)
		decl.resolve(builtin)
		return true
	})
	p.Types.Names.Iter(func(s string, decl *TypeDecl[*NamedType]) bool {
		decl.resolve(types)
		decl.resolve(builtin)
		return true
	})
	p.Types.Maps.Iter(func(s string, decl *TypeDecl[*MapType]) bool {
		decl.resolve(types)
		decl.resolve(builtin)
		return true
	})
	p.Types.Slices.Iter(func(s string, decl *TypeDecl[*SliceType]) bool {
		decl.resolve(types)
		decl.resolve(builtin)
		return true
	})
	p.Types.Arrays.Iter(func(s string, decl *TypeDecl[*ArrayType]) bool {
		decl.resolve(types)
		decl.resolve(builtin)
		return true
	})

	return p, nil
}

func parseFile(bc ParseContext, pkg importTarget, filename string) (*TypeDeclarations, error) {
	fset := token.NewFileSet()
	fullpath := filepath.Join(pkg.Dir, filename)
	afile, err := parser.ParseFile(fset, fullpath, nil, parser.Mode(0))
	if err != nil {
		return nil, err
	}

	imports := []ImportStatment{}
	types := map[string]Type{}
	for _, is := range afile.Imports {
		n := safeDeref(is.Name).Name
		p := strings.Trim(is.Path.Value, `"`)
		if strings.HasPrefix(p, ".") {
			p = path.Join(pkg.ImportPath, p)
		}

		switch n {
		case "_":
			// ignore
			continue
		case ".":
			dot, err := bc.Import(p)
			if err != nil {
				continue // skip unused transitive import
			}
			for _, decl := range dot.Types.Structs.Slice() {
				types[decl.Name] = &NamedType{ImportPath: p, Name: decl.Name}
			}
			for _, decl := range dot.Types.Interfaces.Slice() {
				types[decl.Name] = &NamedType{ImportPath: p, Name: decl.Name}
			}
			for _, decl := range dot.Types.Funcs.Slice() {
				types[decl.Name] = &NamedType{ImportPath: p, Name: decl.Name}
			}
			for _, decl := range dot.Types.Names.Slice() {
				types[decl.Name] = &NamedType{ImportPath: p, Name: decl.Name}
			}
			continue
		case "":
			imported, err := bc.Import(p)
			if err != nil {
				continue // skip unused transitive import
			}
			n = imported.DefaultName
		}

		imports = append(imports, ImportStatment{Name: n, Path: p})
	}

	decls := &TypeDeclarations{
		Structs:    maps.NewOrdered[string, *TypeDecl[*StructType]](),
		Interfaces: maps.NewOrdered[string, *TypeDecl[*InterfaceType]](),
		Funcs:      maps.NewOrdered[string, *TypeDecl[*FuncType]](),
		Names:      maps.NewOrdered[string, *TypeDecl[*NamedType]](),
		Maps:       maps.NewOrdered[string, *TypeDecl[*MapType]](),
		Slices:     maps.NewOrdered[string, *TypeDecl[*SliceType]](),
		Arrays:     maps.NewOrdered[string, *TypeDecl[*ArrayType]](),
		Unresolved: maps.NewOrdered[string, *TypeDecl[*unknwonType]](),
	}

	imp := pkg.ImportPath
	ast.Inspect(afile, func(n ast.Node) bool {
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
			name := safeDeref(typeSpec.Name).Name
			tps := parseTypeParam(imports, safeDeref(typeSpec.TypeParams).List)
			body := parseType(imports, typeSpec.Type)

			tpsmap := map[string]Type{}
			for i := range tps {
				tpsmap[tps[i].Name] = tps[i]
			}
			body.resolve(tpsmap)

			switch b := body.(type) {
			case *StructType:
				decls.Structs.Put(name, &TypeDecl[*StructType]{
					DefinedIn: fullpath, ImportPath: imp,
					Name: name, TypeParams: tps, Body: b,
				})
				types[name] = &NamedType{ImportPath: imp, Name: name}
			case *InterfaceType:
				decls.Interfaces.Put(name, &TypeDecl[*InterfaceType]{
					DefinedIn: fullpath, ImportPath: imp,
					Name: name, TypeParams: tps, Body: b,
				})
				types[name] = &NamedType{ImportPath: imp, Name: name}
			case *FuncType:
				decls.Funcs.Put(name, &TypeDecl[*FuncType]{
					DefinedIn: fullpath, ImportPath: imp,
					Name: name, TypeParams: tps, Body: b,
				})
				types[name] = &NamedType{ImportPath: imp, Name: name}
			case *NamedType:
				decls.Names.Put(name, &TypeDecl[*NamedType]{
					DefinedIn: fullpath, ImportPath: imp,
					Name: name, TypeParams: tps, Body: b,
				})
				types[name] = &NamedType{ImportPath: imp, Name: name}
			case *MapType:
				decls.Maps.Put(name, &TypeDecl[*MapType]{
					DefinedIn: fullpath, ImportPath: imp,
					Name: name, TypeParams: tps, Body: b,
				})
				types[name] = &NamedType{ImportPath: imp, Name: name}
			case *SliceType:
				decls.Slices.Put(name, &TypeDecl[*SliceType]{
					DefinedIn: fullpath, ImportPath: imp,
					Name: name, TypeParams: tps, Body: b,
				})
				types[name] = &NamedType{ImportPath: imp, Name: name}
			case *ArrayType:
				decls.Arrays.Put(name, &TypeDecl[*ArrayType]{
					DefinedIn: fullpath, ImportPath: imp,
					Name: name, TypeParams: tps, Body: b,
				})
				types[name] = &NamedType{ImportPath: imp, Name: name}
			case *unknwonType:
				decls.Unresolved.Put(name, &TypeDecl[*unknwonType]{
					DefinedIn: fullpath, ImportPath: imp,
					Name: name, TypeParams: tps, Body: b,
				})
				types[name] = &NamedType{ImportPath: imp, Name: name}
			}
		}
		return false
	})

	for _, decl := range decls.Structs.Slice() {
		decl.resolve(types)
	}
	for _, decl := range decls.Interfaces.Slice() {
		decl.resolve(types)
	}
	for _, decl := range decls.Funcs.Slice() {
		decl.resolve(types)
	}
	for _, decl := range decls.Names.Slice() {
		decl.resolve(types)
	}
	for _, decl := range decls.Maps.Slice() {
		decl.resolve(types)
	}
	for _, decl := range decls.Slices.Slice() {
		decl.resolve(types)
	}
	for _, decl := range decls.Arrays.Slice() {
		decl.resolve(types)
	}
	for _, decl := range decls.Unresolved.Slice() {
		det := decl.Body.detect(types)
		det.resolve(types)
		switch dt := det.(type) {
		case *NamedType:
			decls.Unresolved.Delete(decl.Name)
			decls.Names.Put(decl.Name, &TypeDecl[*NamedType]{
				DefinedIn: decl.DefinedIn, ImportPath: decl.ImportPath,
				Name: decl.Name, TypeParams: decl.TypeParams, Body: dt,
			})
		default: // *unknownType
		}
	}
	return decls, nil
}

type TypeDeclarations struct {
	Structs    maps.OrderedMap[string, *TypeDecl[*StructType]]
	Interfaces maps.OrderedMap[string, *TypeDecl[*InterfaceType]]
	Funcs      maps.OrderedMap[string, *TypeDecl[*FuncType]]
	Names      maps.OrderedMap[string, *TypeDecl[*NamedType]]
	Maps       maps.OrderedMap[string, *TypeDecl[*MapType]]
	Slices     maps.OrderedMap[string, *TypeDecl[*SliceType]]
	Arrays     maps.OrderedMap[string, *TypeDecl[*ArrayType]]
	Unresolved maps.OrderedMap[string, *TypeDecl[*unknwonType]]
}

func (tds *TypeDeclarations) Merge(other *TypeDeclarations) {
	other.Structs.Iter(func(s string, decl *TypeDecl[*StructType]) bool {
		tds.Structs.Put(s, decl)
		return true
	})
	other.Interfaces.Iter(func(s string, decl *TypeDecl[*InterfaceType]) bool {
		tds.Interfaces.Put(s, decl)
		return true
	})
	other.Funcs.Iter(func(s string, decl *TypeDecl[*FuncType]) bool {
		tds.Funcs.Put(s, decl)
		return true
	})
	other.Names.Iter(func(s string, decl *TypeDecl[*NamedType]) bool {
		tds.Names.Put(s, decl)
		return true
	})
	other.Maps.Iter(func(s string, decl *TypeDecl[*MapType]) bool {
		tds.Maps.Put(s, decl)
		return true
	})
	other.Slices.Iter(func(s string, decl *TypeDecl[*SliceType]) bool {
		tds.Slices.Put(s, decl)
		return true
	})
	other.Arrays.Iter(func(s string, decl *TypeDecl[*ArrayType]) bool {
		tds.Arrays.Put(s, decl)
		return true
	})
	other.Unresolved.Iter(func(s string, decl *TypeDecl[*unknwonType]) bool {
		tds.Unresolved.Put(s, decl)
		return true
	})
}

var builtin = map[string]Type{}

/*
 */
func init() {
	for _, typename := range []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "complex64", "complex128",
		"bool", "rune", "byte", "uintptr",
		"string", "error", "any",
		"comparable",
	} {
		builtin[typename] = &NamedType{Name: typename}
	}
}

func parseTypeParam(pkgs []ImportStatment, params []*ast.Field) []*TypeParam {
	tps := []*TypeParam{}
	tpmap := map[string]Type{}

	for _, tp := range params {
		constraint := parseType(pkgs, tp.Type)
		for i := range tp.Names {
			name := tp.Names[i].Name
			param := &TypeParam{Name: name, Constraint: constraint}
			tps = append(tps, param)
			tpmap[name] = param
		}
	}

	for _, tp := range tps {
		tp.resolve(tpmap)
	}

	return tps
}

func parseType(imports []ImportStatment, node ast.Expr) Type {
	switch t := node.(type) {
	case *ast.Ident: // local or built-in type
		return &unknwonType{Name: t.Name}

	case *ast.SelectorExpr: // imported type
		x, ok := t.X.(*ast.Ident)
		if !ok {
			return &ParseError{
				expected: "NAME.selector",
				node:     node,
			}
		}
		var imp string
		for _, is := range imports {
			if is.Name == x.Name {
				imp = is.Path
				break
			}
		}
		if imp == "" {
			return &ParseError{
				expected: "NAME.selector (not imported?)",
				node:     node,
			}
		}
		return &NamedType{ImportPath: imp, Name: t.Sel.Name}

	case *ast.IndexExpr: // generics
		hostType := parseType(imports, t.X)
		tp := parseType(imports, t.Index)
		switch ht := hostType.(type) {
		case *NamedType:
			ht.Params = []Type{tp}
			return ht
		case *unknwonType:
			ht.Params = []Type{tp}
			return ht
		default:
			return &ParseError{
				expected: "NAME[...]",
				node:     node,
			}
		}
	case *ast.IndexListExpr: // generics
		hostType := parseType(imports, t.X)
		params := []Type{}
		for i := range t.Indices {
			tp := parseType(imports, t.Indices[i])
			params = append(params, tp)
		}
		switch ht := hostType.(type) {
		case *NamedType:
			ht.Params = params
			return ht
		case *unknwonType:
			ht.Params = params
			return ht
		default:
			return &ParseError{
				expected: "NAME[..., ...]",
				node:     node,
			}
		}
	case *ast.StarExpr: // pointer
		elem := parseType(imports, t.X)
		return &PointerType{Elem: elem}

	case *ast.ArrayType: // array or slice
		elem := parseType(imports, t.Elt)
		if l := t.Len; l == nil {
			return &SliceType{Elem: elem}
		} else {
			id, ok := l.(*ast.BasicLit)
			if !ok {
				return &ParseError{
					expected: "array[INDEX] (should be literal)",
					node:     node,
				}
			}

			length, err := strconv.Atoi(id.Value)
			if err != nil {
				return &ParseError{
					expected: "array[INDEX] (should be numeric)",
					node:     node,
				}
			}
			return &ArrayType{Elem: elem, Len: length}
		}

	case *ast.MapType:
		key := parseType(imports, t.Key)
		val := parseType(imports, t.Value)
		return &MapType{Key: key, Elem: val}

	case *ast.ChanType: // channel
		elem := parseType(imports, t.Value)
		return &ChanType{Elem: elem, Dir: t.Dir}

	case *ast.FuncType: // func
		return parseFn(imports, t)

	case *ast.StructType: // struct literal
		return parseStruct(imports, t)

	case *ast.InterfaceType: // interface literal
		return parseInterface(imports, t)

	case *ast.UnaryExpr:
		tt := parseType(imports, t.X)
		if t.Op.String() != "~" {
			return &ParseError{
				expected: "~type (operator is not ~)",
				node:     node,
			}
		}
		return &TypeConstraint{
			Op:   t.Op.String(),
			Type: tt,
		}

	case *ast.BinaryExpr:
		tx := parseType(imports, t.X)
		ty := parseType(imports, t.Y)
		if t.Op.String() != "|" {
			return &ParseError{
				expected: "type | type (operator is not |)",
				node:     node,
			}
		}
		return &TypeUnion{
			Op: t.Op.String(),
			X:  tx,
			Y:  ty,
		}

	default:
		return &ParseError{
			expected: "unknown syntax",
			node:     node,
		}
	}
}

func parseStruct(imports []ImportStatment, snode *ast.StructType) *StructType {
	fields := []*Field{}
	for _, f := range snode.Fields.List {
		typ := parseType(imports, f.Type)

		if 0 < len(f.Names) {
			for _, n := range f.Names {
				if isPrivateName(n.Name) {
					continue
				}
				fields = append(fields, &Field{Name: n.Name, Type: typ})
			}
		} else {
			// embedded!
			t := typ
			if ptr, ok := t.(*PointerType); ok {
				t = ptr.Elem
			}
			var name string

			switch tt := t.(type) {
			case *NamedType:
				name = tt.Name
			case *unknwonType:
				name = tt.Name
			default:
				fields = append(
					fields,
					&Field{
						Name: "?",
						Type: &ParseError{
							expected: "NAME or NAME.SELECTOR",
							node:     f.Type,
						},
					},
				)
				continue
			}
			fields = append(fields, &Field{Name: name, Type: typ})
		}
	}
	return &StructType{Fields: fields}

}

func parseInterface(imports []ImportStatment, node *ast.InterfaceType) *InterfaceType {
	methods := []*Method{}
	embeddeds := []Type{}
	for _, f := range node.Methods.List {
		switch m := f.Type.(type) {
		case *ast.FuncType:
			typ := parseFn(imports, m)
			if len(f.Names) == 0 {
				// embedded
				methods = append(methods, &Method{Name: "", Func: typ})
			} else {
				for _, n := range f.Names {
					methods = append(methods, &Method{Name: n.Name, Func: typ})
				}
			}
		default:
			embeddeds = append(embeddeds, parseType(imports, m))
		}
	}
	return &InterfaceType{Methods: methods, Embedded: embeddeds}
}

func parseFn(imports []ImportStatment, fnode *ast.FuncType) *FuncType {
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
				e := parseType(imports, p.Elt)
				ellip = &FuncIOParam{
					ParamName: name,
					Type:      e,
					Variadic:  true,
				}
			default:
				e := parseType(imports, p)
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
			r := parseType(imports, result.Type)
			if names := result.Names; len(names) == 0 {
				returns = append(
					returns,
					&FuncIOParam{ParamName: "", Type: r},
				)
			} else {
				for _, name := range names {
					returns = append(
						returns,
						&FuncIOParam{ParamName: name.Name, Type: r},
					)
				}
			}
		}
	}

	return &FuncType{
		Args: params, VarArg: ellip, Returns: returns,
	}
}

func safeDeref[T any](val *T) T {
	if val == nil {
		return *new(T)
	}
	return *val
}
