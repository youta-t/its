package parser

import (
	"cmp"
	"fmt"
	"go/ast"
	"io"
	"slices"
	"strings"
)

type Package struct {
	DefaultName string
	Path        string
	Src         string
	Types       *TypeDeclarations
}

type Namespace struct {
	TypeParam []*TypeParam
	Local     *Package
	DotImport []*Package
}

type Imports struct {
	imports map[string]string
}

func (imp *Imports) Add(importPath string) {
	if imp.imports == nil {
		imp.imports = map[string]string{}
	} else if _, ok := imp.imports[importPath]; ok {
		return
	}

	imp.imports[importPath] = fmt.Sprintf("pkg%d", len(imp.imports)+1)
}

func (imp *Imports) GetName(importPath string) string {
	return imp.imports[importPath]
}

func (imp *Imports) Slice() []ImportStatment {
	iss := []ImportStatment{}
	for impPath, name := range imp.imports {
		iss = append(iss, ImportStatment{Name: name, Path: impPath})
	}
	slices.SortFunc(
		iss,
		func(a, b ImportStatment) int { return cmp.Compare(a.Path, b.Path) },
	)
	return iss
}

type ImportStatment struct {
	Name string
	Path string
}

type TypeDecl interface {
	Name() *NamedType
	Instantiate(params []Type) Type
}

type TypeStructDecl struct {
	DefinedIn  string
	ImportPath string
	Name       string
	TypeParams []*TypeParam
	Body       *StructType
}

func (s *TypeStructDecl) InstantiateName(typeParams []Type) *NamedType {
	return &NamedType{
		ImportPath: s.ImportPath,
		Name:       s.Name,
		Params:     typeParams,
	}
}

func (s *TypeStructDecl) GenericExpr(imports *Imports, backtype bool) string {
	typeParams := ""
	{
		tps := []string{}
		for _, tp := range s.TypeParams {
			if backtype {
				tps = append(tps, tp.Name+" "+tp.Constraint.Expr(imports))
			} else {
				tps = append(tps, tp.Name)
			}
		}
		if 0 < len(tps) {
			typeParams = "[" + strings.Join(tps, ", ") + "]"
		}
	}
	return typeParams
}

func (s *TypeStructDecl) Expr(imports *Imports) string {
	expr := s.Name + s.GenericExpr(imports, false)
	pkg := imports.GetName(s.ImportPath)
	return pkg + "." + expr
}

func (s *TypeStructDecl) Require() []string {
	req := []string{}
	for i := range s.TypeParams {
		req = append(req, s.TypeParams[i].Constraint.Require()...)
	}
	req = append(req, s.Body.Require()...)
	return req
}

func (s *TypeStructDecl) IsOpaque() bool {
	for _, tp := range s.TypeParams {
		if tp.IsOpaque() {
			return true
		}
	}
	return s.Body.IsOpaque() || isPrivateName(s.Name)
}

func (decl *TypeStructDecl) resolve(ns map[string]Type) {
	nstp := map[string]Type{}
	for n := range ns {
		nstp[n] = ns[n]
	}
	for _, tp := range decl.TypeParams {
		nstp[tp.Name] = tp
	}

	for _, tp := range decl.TypeParams {
		tp.resolve(nstp)
	}
	decl.Body.resolve(nstp)
}

type TypeInterfaceDecl struct {
	DefinedIn  string
	ImportPath string
	Name       string
	TypeParams []*TypeParam
	Body       *InterfaceType
}

func (s *TypeInterfaceDecl) PlainName() string {
	return s.Name
}

func (s *TypeInterfaceDecl) GenericExpr(imports *Imports, backtype bool) string {
	typeParams := ""
	{
		tps := []string{}
		for _, tp := range s.TypeParams {
			if backtype {
				tps = append(tps, tp.Name+" "+tp.Constraint.Expr(imports))
			} else {
				tps = append(tps, tp.Name)
			}
		}
		if 0 < len(tps) {
			typeParams = "[" + strings.Join(tps, ", ") + "]"
		}
	}
	return typeParams
}

func (s *TypeInterfaceDecl) Expr(imports *Imports) string {
	pkg := imports.GetName(s.ImportPath)
	return pkg + "." + s.Name + s.GenericExpr(imports, false)
}

func (s *TypeInterfaceDecl) Require() []string {
	req := []string{}
	for i := range s.TypeParams {
		req = append(req, s.TypeParams[i].Constraint.Require()...)
	}
	req = append(req, s.Body.Require()...)
	return req
}

func (in *TypeInterfaceDecl) IsOpaque() bool {
	for _, tp := range in.TypeParams {
		if tp.IsOpaque() {
			return true
		}
	}
	return in.Body.IsOpaque() || isPrivateName(in.Name)
}

func (decl *TypeInterfaceDecl) resolve(ns map[string]Type) {
	nstp := map[string]Type{}
	for n := range ns {
		nstp[n] = ns[n]
	}
	for _, tp := range decl.TypeParams {
		nstp[tp.Name] = tp
	}

	for _, tp := range decl.TypeParams {
		tp.resolve(nstp)
	}
	decl.Body.resolve(nstp)
}

type TypeFuncDecl struct {
	DefinedIn  string
	ImportPath string
	Name       string
	TypeParams []*TypeParam
	Body       *FuncType
}

func (fn *TypeFuncDecl) PlainName() string {
	return fn.Name
}

func (fn *TypeFuncDecl) GenericExpr(imports *Imports, backtype bool) string {
	typeParams := ""
	{
		tps := []string{}
		for _, tp := range fn.TypeParams {
			if backtype {
				tps = append(tps, tp.Name+" "+tp.Constraint.Expr(imports))
			} else {
				tps = append(tps, tp.Name)
			}
		}
		if 0 < len(tps) {
			typeParams = "[" + strings.Join(tps, ", ") + "]"
		}
	}
	return typeParams
}

func (fn *TypeFuncDecl) Expr(imports *Imports) string {
	pkg := imports.GetName(fn.ImportPath)
	return pkg + "." + fn.Name + fn.GenericExpr(imports, false)
}

func (fn *TypeFuncDecl) Require() []string {
	req := []string{}
	for i := range fn.TypeParams {
		req = append(req, fn.TypeParams[i].Constraint.Require()...)
	}
	req = append(req, fn.Body.Require()...)
	return req
}

func (fn *TypeFuncDecl) IsOpaque() bool {
	for _, tp := range fn.TypeParams {
		if tp.IsOpaque() {
			return true
		}
	}
	return fn.Body.IsOpaque() || isPrivateName(fn.Name)
}

func (decl *TypeFuncDecl) resolve(ns map[string]Type) {
	nstp := map[string]Type{}
	for n := range ns {
		nstp[n] = ns[n]
	}
	for _, tp := range decl.TypeParams {
		nstp[tp.Name] = tp
	}

	for _, tp := range decl.TypeParams {
		tp.resolve(nstp)
	}
	decl.Body.resolve(nstp)
}

type TypeNameDecl struct {
	DefinedIn  string
	ImportPath string
	Name       string
	TypeParams []*TypeParam
	Body       *NamedType
}

func (n *TypeNameDecl) GenericExpr(imports *Imports, backtype bool) string {
	typeParams := ""
	{
		tps := []string{}
		for _, tp := range n.TypeParams {
			if backtype {
				tps = append(tps, tp.Name+" "+tp.Constraint.Expr(imports))
			} else {
				tps = append(tps, tp.Name)
			}
		}
		if 0 < len(tps) {
			typeParams = "[" + strings.Join(tps, ", ") + "]"
		}
	}
	return typeParams
}

func (n *TypeNameDecl) Expr(imports *Imports) string {
	pkg := imports.GetName(n.ImportPath)
	return pkg + "." + n.Name + n.GenericExpr(imports, false)
}

func (n *TypeNameDecl) Require() []string {
	req := []string{}
	for i := range n.TypeParams {
		req = append(req, n.TypeParams[i].Constraint.Require()...)
	}
	req = append(req, n.Body.Require()...)
	return req
}

func (decl *TypeNameDecl) resolve(ns map[string]Type) {
	nstp := map[string]Type{}
	for n := range ns {
		nstp[n] = ns[n]
	}
	for _, tp := range decl.TypeParams {
		nstp[tp.Name] = tp
	}

	for _, tp := range decl.TypeParams {
		tp.resolve(nstp)
	}
	decl.Body.resolve(nstp)
}

func (fn *TypeNameDecl) IsOpaque() bool {
	for _, tp := range fn.TypeParams {
		if tp.IsOpaque() {
			return true
		}
	}
	return fn.Body.IsOpaque() || isPrivateName(fn.Name)
}

type TypeUnresolvedDecl struct {
	DefinedIn  string
	ImportPath string
	Name       string
	TypeParams []*TypeParam
	Body       *unknwonType
}

func (n *TypeUnresolvedDecl) GenericExpr(imports *Imports, backtype bool) string {
	typeParams := ""
	{
		tps := []string{}
		for _, tp := range n.TypeParams {
			if backtype {
				tps = append(tps, tp.Name+" "+tp.Constraint.Expr(imports))
			} else {
				tps = append(tps, tp.Name)
			}
		}
		if 0 < len(tps) {
			typeParams = "[" + strings.Join(tps, ", ") + "]"
		}
	}
	return typeParams
}

func (n *TypeUnresolvedDecl) Expr(imports *Imports) string {
	pkg := imports.GetName(n.ImportPath)
	return pkg + "." + n.Name + n.GenericExpr(imports, false)
}

func (n *TypeUnresolvedDecl) Require() []string {
	req := []string{}
	for i := range n.TypeParams {
		req = append(req, n.TypeParams[i].Constraint.Require()...)
	}
	req = append(req, n.Body.Require()...)
	return req
}

func (decl *TypeUnresolvedDecl) resolve(ns map[string]Type) {
	nstp := map[string]Type{}
	for n := range ns {
		nstp[n] = ns[n]
	}
	for _, tp := range decl.TypeParams {
		nstp[tp.Name] = tp
	}

	for _, tp := range decl.TypeParams {
		tp.resolve(nstp)
	}
	decl.Body.resolve(nstp)
}

func (fn *TypeUnresolvedDecl) IsOpaque() bool {
	for _, tp := range fn.TypeParams {
		if tp.IsOpaque() {
			return true
		}
	}
	return fn.Body.IsOpaque() || isPrivateName(fn.Name)
}

type TypeParam struct {
	Name       string
	Constraint Type
}

func (t *TypeParam) Expr(*Imports) string { return t.Name }
func (t *TypeParam) Require() []string    { return t.Constraint.Require() }

func (t *TypeParam) resolve(namespace map[string]Type) {
	if p, ok := t.Constraint.(*unknwonType); ok {
		t.Constraint = p.detect(namespace)
	} else {
		t.Constraint.resolve(namespace)
		return
	}
}

func (t *TypeParam) IsOpaque() bool {
	return t.Constraint.IsOpaque()
}

type Type interface {
	Expr(*Imports) string
	Require() []string
	IsOpaque() bool

	resolve(map[string]Type)
}

type unknwonType struct {
	Name   string
	Params []Type
}

func (n *unknwonType) Expr(*Imports) string { return "/* unknown! */ " + n.Name }
func (*unknwonType) Require() []string      { return []string{} }
func (unk *unknwonType) resolve(ns map[string]Type) {
	for i := range unk.Params {
		pt := unk.Params[i]
		if p, ok := pt.(*unknwonType); ok {
			unk.Params[i] = p.detect(ns)
		} else {
			pt.resolve(ns)
		}
	}
}
func (*unknwonType) TypeParams() []*TypeParam { return []*TypeParam{} }
func (*unknwonType) IsOpaque() bool           { return true }

func (unk *unknwonType) detect(namespace map[string]Type) Type {
	name := unk.Name
	for n := range namespace {
		if n != name {
			continue
		}
		switch t := namespace[n].(type) {
		case *TypeParam:
			return t
		case *NamedType:
			return &NamedType{
				ImportPath: t.ImportPath,
				Name:       unk.Name,
				Params:     unk.Params,
			}
		}
	}

	return unk
}

type NamedType struct {
	ImportPath string
	Name       string
	Params     []Type
}

func (n NamedType) PlainName() string {
	return n.Name
}

func (n *NamedType) Expr(imports *Imports) string {
	expr := n.Name
	if n.ImportPath != "" {
		imp := imports.GetName(n.ImportPath)
		expr = imp + "." + expr
	}
	if 0 < len(n.Params) {
		params := []string{}
		for _, p := range n.Params {
			params = append(params, p.Expr(imports))
		}
		expr = expr + "[" + strings.Join(params, ", ") + "]"
	}

	return expr
}

func (n *NamedType) Require() []string {
	reqset := map[string]struct{}{}
	if n.ImportPath != "" {
		reqset[n.ImportPath] = struct{}{}
	}

	for i := range n.Params {
		requires := n.Params[i].Require()
		for j := range requires {
			reqset[requires[j]] = struct{}{}
		}
	}

	req := make([]string, 0, len(reqset))
	for p := range reqset {
		req = append(req, p)
	}

	return req
}

func (nt *NamedType) resolve(ns map[string]Type) {
	for i := range nt.Params {
		pt := nt.Params[i]
		if p, ok := pt.(*unknwonType); ok {
			nt.Params[i] = p.detect(ns)
		} else {
			pt.resolve(ns)
		}
	}
}
func (*NamedType) TypeParams() []*TypeParam { return []*TypeParam{} }
func (nt *NamedType) IsOpaque() bool {
	if nt.ImportPath == "" {
		return false // emdedded type
	}
	for _, p := range nt.Params {
		if p.IsOpaque() {
			return true
		}
	}
	return isPrivateName(nt.Name)
}

type StructType struct {
	Fields []*Field
}

func (s *StructType) PlainName() string {
	return ""
}

func (l *StructType) Expr(imports *Imports) string {
	buf := new(strings.Builder)
	io.WriteString(buf, "struct{")
	for _, f := range l.Fields {
		io.WriteString(buf, "\n\t")
		if f.Name != "" {
			io.WriteString(buf, f.Name)
			io.WriteString(buf, " ")
		}
		io.WriteString(buf, f.Type.Expr(imports))
	}
	io.WriteString(buf, "\n}")
	return buf.String()
}

func (s *StructType) Require() []string {
	req := []string{}
	for i := range s.Fields {
		req = append(req, s.Fields[i].Type.Require()...)
	}
	return req
}

func (s *StructType) resolve(ns map[string]Type) {
	for i := range s.Fields {
		fld := s.Fields[i]
		if unk, ok := fld.Type.(*unknwonType); ok {
			fld.Type = unk.detect(ns)
		}
		fld.Type.resolve(ns)
	}
}

func (s *StructType) IsOpaque() bool {
	for _, f := range s.Fields {
		if f.IsOpaque() {
			return true
		}
	}

	return false
}

type Field struct {
	Name string
	Type Type
}

func (f Field) IsOpaque() bool {
	_, isStruct := f.Type.(*StructType)
	return isPrivateName(f.Name) || (!isStruct && f.Type.IsOpaque())
}

type MapType struct {
	Key  Type
	Elem Type
}

func (*MapType) PlainName() string {
	return ""
}

func (m *MapType) Expr(imports *Imports) string {
	return fmt.Sprintf(`map[%s]%s`, m.Key.Expr(imports), m.Elem.Expr(imports))
}

func (m *MapType) Require() []string {
	req := []string{}
	req = append(req, m.Key.Require()...)
	req = append(req, m.Elem.Require()...)
	return req
}

func (m *MapType) resolve(ns map[string]Type) {
	if p, ok := m.Key.(*unknwonType); ok {
		m.Key = p.detect(ns)
	} else {
		m.Key.resolve(ns)
	}

	if p, ok := m.Elem.(*unknwonType); ok {
		m.Elem = p.detect(ns)
	} else {
		m.Elem.resolve(ns)
	}
}

func (m *MapType) IsOpaque() bool {
	return m.Key.IsOpaque() || m.Elem.IsOpaque()
}

type SliceType struct {
	Elem Type
}

func (s *SliceType) Expr(imports *Imports) string {
	return "[]" + s.Elem.Expr(imports)
}

func (s *SliceType) Require() []string { return s.Elem.Require() }

func (s *SliceType) resolve(ns map[string]Type) {
	if p, ok := s.Elem.(*unknwonType); ok {
		s.Elem = p.detect(ns)
	} else {
		s.Elem.resolve(ns)
	}
}

func (s *SliceType) IsOpaque() bool { return s.Elem.IsOpaque() }

type ArrayType struct {
	Len  int
	Elem Type
}

func (a *ArrayType) Expr(imports *Imports) string {
	return fmt.Sprintf("[%d]%s", a.Len, a.Elem.Expr(imports))
}

func (a *ArrayType) Require() []string { return a.Elem.Require() }

func (a *ArrayType) resolve(ns map[string]Type) {
	if unk, ok := a.Elem.(*unknwonType); ok {
		a.Elem = unk.detect(ns)
	} else {
		a.Elem.resolve(ns)
	}
}

func (a *ArrayType) IsOpaque() bool { return a.Elem.IsOpaque() }

type PointerType struct {
	Elem Type
}

func (ptr *PointerType) Expr(imports *Imports) string {
	return "*" + ptr.Elem.Expr(imports)
}

func (ptr *PointerType) Require() []string { return ptr.Elem.Require() }

func (ptr *PointerType) resolve(ns map[string]Type) {
	if unk, ok := ptr.Elem.(*unknwonType); ok {
		ptr.Elem = unk.detect(ns)
	} else {
		ptr.Elem.resolve(ns)
	}
}

func (ptr *PointerType) IsOpaque() bool { return ptr.Elem.IsOpaque() }

type ChanType struct {
	Dir  ast.ChanDir
	Elem Type
}

func (ch *ChanType) Expr(imports *Imports) string {
	switch ch.Dir {
	case ast.RECV:
		return "<-chan " + ch.Elem.Expr(imports)
	case ast.SEND:
		return "chan<- " + ch.Elem.Expr(imports)
	default:
		return "chan " + ch.Elem.Expr(imports)
	}
}

func (ch *ChanType) Require() []string { return ch.Elem.Require() }

func (ch *ChanType) resolve(ns map[string]Type) {
	if unk, ok := ch.Elem.(*unknwonType); ok {
		ch.Elem = unk.detect(ns)
	} else {
		ch.Elem.resolve(ns)
	}
}

func (ch *ChanType) IsOpaque() bool { return ch.Elem.IsOpaque() }

type TypeConstraint struct {
	Op   string
	Type Type
}

func (tc *TypeConstraint) Expr(imports *Imports) string {
	return tc.Op + tc.Type.Expr(imports)
}

func (tc *TypeConstraint) Require() []string { return tc.Type.Require() }

func (tx *TypeConstraint) resolve(ns map[string]Type) {
	if unk, ok := tx.Type.(*unknwonType); ok {
		tx.Type = unk.detect(ns)
	} else {
		tx.Type.resolve(ns)
	}
}

func (tc *TypeConstraint) IsOpaque() bool { return tc.Type.IsOpaque() }

type TypeUnion struct {
	Op string
	X  Type
	Y  Type
}

func (uni *TypeUnion) Expr(imports *Imports) string {
	return fmt.Sprintf("%s | %s", uni.X.Expr(imports), uni.Y.Expr(imports))
}

func (uni *TypeUnion) Require() []string {
	req := map[string]struct{}{}
	for _, im := range uni.X.Require() {
		req[im] = struct{}{}
	}
	for _, im := range uni.Y.Require() {
		req[im] = struct{}{}
	}
	ret := []string{}
	for p := range req {
		ret = append(ret, p)
	}
	return ret
}

func (uni *TypeUnion) resolve(ns map[string]Type) {
	if unk, ok := uni.X.(*unknwonType); ok {
		uni.X = unk.detect(ns)
	} else {
		uni.X.resolve(ns)
	}

	if unk, ok := uni.Y.(*unknwonType); ok {
		uni.Y = unk.detect(ns)
	} else {
		uni.Y.resolve(ns)
	}
}

func (uni *TypeUnion) IsOpaque() bool {
	return uni.X.IsOpaque() || uni.Y.IsOpaque()
}

type FuncIOParam struct {
	ParamName string
	Type      Type
	Variadic  bool
}

func (f *FuncIOParam) IsNamed() bool {
	return f.ParamName != "" && f.ParamName != "_"
}

func (f *FuncIOParam) ParamNameOr(defaultValue string) string {
	if f.IsNamed() {
		return f.ParamName
	}
	return defaultValue
}

func (fio *FuncIOParam) Expr(imports *Imports) string {
	t := fio.Type.Expr(imports)
	if fio.Variadic {
		t = "..." + t
	}
	return t
}

func (fio *FuncIOParam) ExprWithName(imports *Imports, defaultName string) string {
	name := fio.ParamNameOr(defaultName)
	return name + " " + fio.Expr(imports)
}

func (fio *FuncIOParam) Require() []string { return fio.Type.Require() }

func (fio *FuncIOParam) resolve(ns map[string]Type) {
	if unk, ok := fio.Type.(*unknwonType); ok {
		fio.Type = unk.detect(ns)
	} else {
		fio.Type.resolve(ns)
	}
}

func (fio *FuncIOParam) IsOpaque() bool { return fio.Type.IsOpaque() }

type FuncType struct {
	Args    []*FuncIOParam
	VarArg  *FuncIOParam
	Returns []*FuncIOParam
}

func (c *FuncType) Expr(imports *Imports) string {
	return "func" + c.Signature(imports, false)
}

func (c *FuncType) Signature(imports *Imports, nameSupply bool) string {
	sb := new(strings.Builder)
	io.WriteString(sb, "(")
	params := []string{}
	for i, p := range c.Args {
		params = append(params, p.ExprWithName(imports, fmt.Sprintf("arg%d", i)))
	}
	if c.VarArg != nil {
		params = append(params, c.VarArg.ExprWithName(imports, "vararg"))
	}
	io.WriteString(sb, strings.Join(params, ", "))
	io.WriteString(sb, ")")
	nret := len(c.Returns)
	needBrace := 2 <= nret || nret == 1 && c.Returns[0].IsNamed()
	if needBrace {
		io.WriteString(sb, " (")
	} else if 1 < nret {
		io.WriteString(sb, " ")
	}
	rets := []string{}
	for _, r := range c.Returns {
		rets = append(rets, r.ExprWithName(imports, ""))
	}
	io.WriteString(sb, strings.Join(rets, ", "))
	if needBrace {
		io.WriteString(sb, ")")
	}

	return sb.String()
}

func (f *FuncType) Require() []string {
	reqset := map[string]struct{}{}

	for _, a := range f.Args {
		for _, r := range a.Require() {
			reqset[r] = struct{}{}
		}
	}
	if f.VarArg != nil {
		for _, r := range f.VarArg.Require() {
			reqset[r] = struct{}{}
		}
	}
	for _, ret := range f.Returns {
		for _, r := range ret.Require() {
			reqset[r] = struct{}{}
		}
	}

	req := make([]string, 0, len(reqset))
	for r := range reqset {
		req = append(req, r)
	}

	return req
}

func (f *FuncType) resolve(ns map[string]Type) {
	for _, p := range f.Args {
		p.resolve(ns)
	}

	if f.VarArg != nil {
		f.VarArg.resolve(ns)
	}

	for _, r := range f.Returns {
		r.resolve(ns)
	}
}

func (f *FuncType) IsOpaque() bool {
	if f.VarArg != nil && f.VarArg.IsOpaque() {
		return true
	}

	for _, a := range f.Args {
		if a.IsOpaque() {
			return true
		}
	}

	for _, r := range f.Returns {
		if r.IsOpaque() {
			return true
		}
	}
	return false
}

type InterfaceType struct {
	Methods  []*Method
	Embedded []Type
}

func (*InterfaceType) PlainName() string {
	return ""
}

func (i *InterfaceType) Expr(imports *Imports) string {
	sb := new(strings.Builder)
	io.WriteString(sb, "interface{")
	for _, m := range i.Methods {
		io.WriteString(sb, "\n\t")
		io.WriteString(sb, m.Name)
		io.WriteString(sb, m.Func.Signature(imports, false))
	}
	for _, em := range i.Embedded {
		io.WriteString(sb, "\n\t")
		io.WriteString(sb, em.Expr(imports))
	}
	io.WriteString(sb, "\n}")
	return sb.String()
}

func (in *InterfaceType) Require() []string {
	reqset := map[string]struct{}{}

	for _, m := range in.Methods {
		for _, r := range m.Func.Require() {
			reqset[r] = struct{}{}
		}
	}
	for _, m := range in.Embedded {
		for _, r := range m.Require() {
			reqset[r] = struct{}{}
		}
	}
	req := make([]string, 0, len(reqset))
	for r := range reqset {
		req = append(req, r)
	}
	return req
}

func (in *InterfaceType) resolve(ns map[string]Type) {
	for i := range in.Embedded {
		emb := in.Embedded[i]
		if unk, ok := emb.(*unknwonType); ok {
			in.Embedded[i] = unk.detect(ns)
		} else {
			emb.resolve(ns)
		}
	}
	for _, meth := range in.Methods {
		meth.Func.resolve(ns)
	}
}

func (in *InterfaceType) IsOpaque() bool {
	for _, m := range in.Methods {
		if m.IsOpaque() {
			return true
		}
	}
	for _, em := range in.Embedded {
		if em.IsOpaque() {
			return true
		}
	}
	return false
}

func (in *InterfaceType) Inlined(bc ParseContext) (*InterfaceType, error) {
	inlined := &InterfaceType{
		Methods:  make([]*Method, 0, len(in.Methods)),
		Embedded: []Type{},
	}

	methods := map[string]*Method{}
	for i := range in.Methods {
		m := in.Methods[i]
		methods[m.Name] = m
	}

	resolving := make([]Type, len(in.Embedded))
	copy(resolving, in.Embedded)
	for 0 < len(resolving) {
		r := resolving[0]
		resolving = resolving[1:]

		switch typ := r.(type) {
		case *TypeConstraint, *TypeUnion:
			inlined.Embedded = append(inlined.Embedded, typ)
		case *NamedType:
			if typ.ImportPath == "" {
				// builtin
				inlined.Embedded = append(inlined.Embedded, typ)
				continue
			}

			pkg, err := bc.Import(typ.ImportPath)
			if err != nil {
				return nil, err
			}
			if fn, ok := pkg.Types.Funcs.Get(typ.Name); ok {
				inlined.Embedded = append(inlined.Embedded, fn.Body)
				continue
			}
			if s, ok := pkg.Types.Structs.Get(typ.Name); ok {
				inlined.Embedded = append(inlined.Embedded, s.Body)
				continue
			}
			if ifc, ok := pkg.Types.Interfaces.Get(typ.Name); ok {
				resolving = append(resolving, ifc.Body)
				continue
			}
			if n, ok := pkg.Types.Names.Get(typ.Name); ok {
				resolving = append(resolving, n.Body)
				continue
			}

		case *InterfaceType:
			typ_, err := typ.Inlined(bc)
			if err != nil {
				return nil, err
			}
			for i := range typ_.Methods {
				m := typ_.Methods[i]
				methods[m.Name] = m
			}
			inlined.Embedded = append(inlined.Embedded, typ_.Embedded...)
		}
	}

	for name := range methods {
		inlined.Methods = append(inlined.Methods, methods[name])
	}
	slices.SortFunc(inlined.Methods, func(a, b *Method) int { return cmp.Compare(a.Name, b.Name) })
	return inlined, nil
}

type Method struct {
	Name string
	Func *FuncType
}

func (m *Method) IsOpaque() bool { return isPrivateName(m.Name) || m.Func.IsOpaque() }

type ParseError struct {
	expected string
	node     ast.Expr
}

func (*ParseError) Expr(*Imports) string        { return "parse error!" }
func (pe *ParseError) Require() []string        { return []string{} }
func (pe *ParseError) TypeParams() []*TypeParam { return []*TypeParam{} }
func (pe *ParseError) IsOpaque() bool           { return true }
func (pe *ParseError) resolve(map[string]Type)  {}
func (pe *ParseError) Error() string {
	return fmt.Sprintf("cannot parse as %s: %s", pe.expected, pe.node)
}

func isPrivateName(s string) bool {
	initial := s[:1]
	return initial == "_" || initial != strings.ToUpper(initial)
}
