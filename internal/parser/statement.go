package parser

import (
	"cmp"
	"fmt"
	"go/ast"
	"io"
	"slices"
	"strings"
)

type File struct {
	PackageName string
	Imports     []*Import
	Types       *TypeDeclarations
}

type Import struct {
	Name string
	Path string
}

type TypeStructDecl struct {
	Name       string
	Package    *Import
	TypeParams []*TypeParam
	Body       *StructType
}

func (s *TypeStructDecl) PlainName() string {
	return s.Name
}

func (s *TypeStructDecl) GenericExpr(backtype bool) string {
	typeParams := ""
	{
		tps := []string{}
		for _, tp := range s.TypeParams {
			if backtype {
				tps = append(tps, tp.Name+" "+tp.Back.Expr())
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

func (s *TypeStructDecl) Expr() string {
	pkg := ""
	if s.Package != nil {
		pkg = s.Package.Name + "."
	}
	return pkg + s.Name + s.GenericExpr(false)
}

func (s *TypeStructDecl) AsType() Type {
	return &NamedType{
		Pkg:        s.Package,
		Name:       s.Name,
		isExported: s.Name[:1] == strings.ToUpper(s.Name[:1]),
	}
}

func (s *TypeStructDecl) Require() []*Import {
	req := []*Import{s.Package}
	for i := range s.TypeParams {
		req = append(req, s.TypeParams[i].Back.Require()...)
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
	return s.Body.IsOpaque() || s.Name[:1] != strings.ToUpper(s.Name[:1])
}

type TypeInterfaceDecl struct {
	Name       string
	Package    *Import
	TypeParams []*TypeParam
	Body       *InterfaceType
}

func (s *TypeInterfaceDecl) PlainName() string {
	return s.Name
}

func (s *TypeInterfaceDecl) GenericExpr(backtype bool) string {
	typeParams := ""
	{
		tps := []string{}
		for _, tp := range s.TypeParams {
			if backtype {
				tps = append(tps, tp.Name+" "+tp.Back.Expr())
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

func (s *TypeInterfaceDecl) Expr() string {
	pkg := ""
	if s.Package != nil {
		pkg = s.Package.Name + "."
	}
	return pkg + s.Name + s.GenericExpr(false)
}

func (s *TypeInterfaceDecl) AsType() Type {
	return &NamedType{
		Pkg:        s.Package,
		Name:       s.Name,
		isExported: s.Name[:1] == strings.ToUpper(s.Name[:1]),
	}
}

func (s *TypeInterfaceDecl) Require() []*Import {
	req := []*Import{s.Package}
	for i := range s.TypeParams {
		req = append(req, s.TypeParams[i].Back.Require()...)
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
	return in.Body.IsOpaque() || in.Name[:1] != strings.ToUpper(in.Name[:1])
}

type TypeFuncDecl struct {
	Name       string
	Package    *Import
	TypeParams []*TypeParam
	Body       *FuncType
}

func (fn *TypeFuncDecl) PlainName() string {
	return fn.Name
}

func (fn *TypeFuncDecl) GenericExpr(backtype bool) string {
	typeParams := ""
	{
		tps := []string{}
		for _, tp := range fn.TypeParams {
			if backtype {
				tps = append(tps, tp.Name+" "+tp.Back.Expr())
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

func (fn *TypeFuncDecl) ParamsGenericExpr(backtype bool) string {
	return fn.Body.ParamsGenericExpr(backtype)
}

func (fn *TypeFuncDecl) ReturnGenericExpr(backtype bool) string {
	return fn.Body.ReturnGenericExpr(backtype)
}

func (fn *TypeFuncDecl) Expr() string {
	pkg := ""
	if fn.Package != nil {
		pkg = fn.Package.Name + "."
	}
	return pkg + fn.Name + fn.GenericExpr(false)
}

func (fn *TypeFuncDecl) AsType() Type {
	return &NamedType{
		Pkg:        fn.Package,
		Name:       fn.Name,
		isExported: fn.Name[:1] == strings.ToUpper(fn.Name[:1]),
	}
}

func (fn *TypeFuncDecl) Require() []*Import {
	req := []*Import{fn.Package}
	for i := range fn.TypeParams {
		req = append(req, fn.TypeParams[i].Back.Require()...)
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
	return fn.Body.IsOpaque() || fn.Name[:1] != strings.ToUpper(fn.Name[:1])
}

type TypeParam struct {
	Name string
	Back Type
}

func (t *TypeParam) PlainName() string {
	return t.Name
}

func (t *TypeParam) Expr() string {
	return t.Name
}

func (t *TypeParam) Require() []*Import {
	return t.Back.Require()
}

func (t *TypeParam) TypeParams() []*TypeParam {
	tps := []*TypeParam{t}
	tps = append(tps, t.Back.TypeParams()...)
	slices.SortFunc(tps, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return tps
}

func (t *TypeParam) injectTypeParam(local *Import, tps []*TypeParam) {
	p, ok := t.Back.(*pseudoType)
	if !ok {
		t.Back.injectTypeParam(local, tps)
		return
	}

	t.Back = resolveBareNameType(local, tps, p.Name)
}

func (t *TypeParam) IsOpaque() bool {
	return t.Back.IsOpaque()
}

type Type interface {
	PlainName() string
	Expr() string
	Require() []*Import
	TypeParams() []*TypeParam
	IsOpaque() bool

	injectTypeParam(*Import, []*TypeParam)
}

type pseudoType struct {
	Name string
}

func (n *pseudoType) PlainName() string {
	return n.Name
}

func (n *pseudoType) Expr() string {
	return n.Name
}

func (*pseudoType) Require() []*Import {
	return []*Import{}
}

func (*pseudoType) injectTypeParam(*Import, []*TypeParam) {}
func (*pseudoType) TypeParams() []*TypeParam              { return []*TypeParam{} }

func (*pseudoType) IsOpaque() bool {
	return true
}

type NamedType struct {
	Pkg        *Import
	Name       string
	isExported bool
}

func (n NamedType) PlainName() string {
	return n.Name
}

func (n *NamedType) Expr() string {
	expr := n.Name
	if n.Pkg != nil {
		expr = n.Pkg.Name + "." + expr
	}

	return expr
}

func (n *NamedType) Require() []*Import {
	if n.Pkg != nil {
		return []*Import{n.Pkg}
	}
	return []*Import{}
}

func (*NamedType) injectTypeParam(*Import, []*TypeParam) {}
func (*NamedType) TypeParams() []*TypeParam              { return []*TypeParam{} }
func (nt *NamedType) IsOpaque() bool                     { return !nt.isExported }

type BuiltinType struct {
	Name string
}

func (bt *BuiltinType) PlainName() string {
	return bt.Name
}

func (bt *BuiltinType) Expr() string {
	return bt.Name
}

func (bt *BuiltinType) Require() []*Import {
	return []*Import{}
}

func (*BuiltinType) injectTypeParam(*Import, []*TypeParam) {}
func (*BuiltinType) TypeParams() []*TypeParam              { return []*TypeParam{} }
func (*BuiltinType) IsOpaque() bool                        { return false }

type StructType struct {
	Fields []*Field
}

func (s *StructType) PlainName() string {
	return ""
}

func (l *StructType) Expr() string {
	buf := new(strings.Builder)
	io.WriteString(buf, "struct{")
	for _, f := range l.Fields {
		io.WriteString(buf, "\n\t")
		if f.Name != "" {
			io.WriteString(buf, f.Name)
			io.WriteString(buf, " ")
		}
		io.WriteString(buf, f.Type.Expr())
	}
	io.WriteString(buf, "\n}")
	return buf.String()
}

func (s *StructType) Require() []*Import {
	req := []*Import{}
	for i := range s.Fields {
		req = append(req, s.Fields[i].Type.Require()...)
	}
	return req
}

func (s *StructType) injectTypeParam(local *Import, tps []*TypeParam) {
	for i := range s.Fields {
		fld := s.Fields[i]
		pt, ok := fld.Type.(*pseudoType)
		if !ok {
			fld.Type.injectTypeParam(local, tps)
			continue
		}
		fld.Type = resolveBareNameType(local, tps, pt.Name)
	}
}
func (s *StructType) TypeParams() []*TypeParam {
	tps := map[string]*TypeParam{}
	for _, fld := range s.Fields {
		_tps := fld.Type.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}
	ret := []*TypeParam{}
	for name := range tps {
		ret = append(ret, tps[name])
	}
	slices.SortFunc(ret, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ret
}

func (s *StructType) IsOpaque() bool {
	for _, f := range s.Fields {
		if !f.isExported || f.Type.IsOpaque() {
			return true
		}
	}

	return false
}

type Field struct {
	Name       string
	Type       Type
	isExported bool
}

func (f Field) IsOpaque() bool { return !f.isExported || f.Type.IsOpaque() }

type GenericType struct {
	Host   Type
	Params []Type
}

func (g *GenericType) PlainName() string {
	return g.Host.PlainName()
}

func (g *GenericType) Expr() string {
	param := []string{}
	for _, p := range g.Params {
		param = append(param, p.Expr())
	}
	return fmt.Sprintf("%s[%s]", g.Host.Expr(), strings.Join(param, ", "))
}

func (g *GenericType) Require() []*Import {
	req := []*Import{}
	req = append(req, g.Host.Require()...)
	for i := range g.Params {
		req = append(req, g.Params[i].Require()...)
	}
	return req
}

func (s *GenericType) TypeParams() []*TypeParam {
	tps := map[string]*TypeParam{}

	{
		_tps := s.Host.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}

	for _, para := range s.Params {
		_tps := para.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}
	ret := []*TypeParam{}
	for name := range tps {
		ret = append(ret, tps[name])
	}
	slices.SortFunc(ret, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ret
}

func (g *GenericType) injectTypeParam(local *Import, tps []*TypeParam) {
	p, ok := g.Host.(*pseudoType)
	if !ok {
		g.Host.injectTypeParam(local, tps)
	} else {
		g.Host = resolveBareNameType(local, tps, p.Name)
	}
	for i := range g.Params {
		gp := g.Params[i]
		pt, ok := gp.(*pseudoType)
		if !ok {
			gp.injectTypeParam(local, tps)
			continue
		}
		g.Params[i] = resolveBareNameType(local, tps, pt.Name)
	}
}

func (g *GenericType) IsOpaque() bool {
	if g.Host.IsOpaque() {
		return true
	}

	for _, tp := range g.Params {
		if tp.IsOpaque() {
			return true
		}
	}
	return false
}

type MapType struct {
	Key  Type
	Elem Type
}

func (*MapType) PlainName() string {
	return ""
}

func (m *MapType) Expr() string {
	return fmt.Sprintf(`map[%s]%s`, m.Key.Expr(), m.Elem.Expr())
}

func (m *MapType) Require() []*Import {
	req := []*Import{}
	req = append(req, m.Key.Require()...)
	req = append(req, m.Elem.Require()...)
	return req
}

func (s *MapType) TypeParams() []*TypeParam {
	tps := map[string]*TypeParam{}

	{
		_tps := s.Key.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}
	{
		_tps := s.Elem.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}

	ret := []*TypeParam{}
	for name := range tps {
		ret = append(ret, tps[name])
	}
	slices.SortFunc(ret, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ret
}

func (m *MapType) injectTypeParam(local *Import, tps []*TypeParam) {
	if p, ok := m.Key.(*pseudoType); !ok {
		m.Key.injectTypeParam(local, tps)
	} else {
		m.Key = resolveBareNameType(local, tps, p.Name)
	}

	if p, ok := m.Elem.(*pseudoType); !ok {
		m.Elem.injectTypeParam(local, tps)
	} else {
		m.Elem = resolveBareNameType(local, tps, p.Name)
	}
}

func (m *MapType) IsOpaque() bool {
	return m.Key.IsOpaque() || m.Elem.IsOpaque()
}

type SliceType struct {
	Elem Type
}

func (*SliceType) PlainName() string {
	return ""
}

func (s *SliceType) Expr() string {
	return "[]" + s.Elem.Expr()
}

func (s *SliceType) Require() []*Import {
	return s.Elem.Require()
}

func (s *SliceType) TypeParams() []*TypeParam {
	tps := map[string]*TypeParam{}

	{
		_tps := s.Elem.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}

	ret := []*TypeParam{}
	for name := range tps {
		ret = append(ret, tps[name])
	}
	slices.SortFunc(ret, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ret
}

func (s *SliceType) injectTypeParam(local *Import, tps []*TypeParam) {
	if p, ok := s.Elem.(*pseudoType); !ok {
		s.Elem.injectTypeParam(local, tps)
	} else {
		s.Elem = resolveBareNameType(local, tps, p.Name)
	}
}

func (s *SliceType) IsOpaque() bool {
	return s.Elem.IsOpaque()
}

type ArrayType struct {
	Len  int
	Elem Type
}

func (*ArrayType) PlainName() string {
	return ""
}

func (a *ArrayType) Expr() string {
	return fmt.Sprintf("[%d]%s", a.Len, a.Elem.Expr())
}

func (a *ArrayType) Require() []*Import {
	return a.Elem.Require()
}

func (a *ArrayType) TypeParams() []*TypeParam {
	tps := map[string]*TypeParam{}

	{
		_tps := a.Elem.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}

	ret := []*TypeParam{}
	for name := range tps {
		ret = append(ret, tps[name])
	}
	slices.SortFunc(ret, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ret
}

func (a *ArrayType) injectTypeParam(local *Import, tps []*TypeParam) {
	if p, ok := a.Elem.(*pseudoType); !ok {
		a.Elem.injectTypeParam(local, tps)
	} else {
		a.Elem = resolveBareNameType(local, tps, p.Name)
	}
}

func (a *ArrayType) IsOpaque() bool {
	return a.Elem.IsOpaque()
}

type PointerType struct {
	Elem Type
}

func (ptr *PointerType) PlainName() string {
	return ptr.Elem.PlainName()
}

func (ptr *PointerType) Expr() string {
	return "*" + ptr.Elem.Expr()
}

func (ptr *PointerType) Require() []*Import {
	return ptr.Elem.Require()
}

func (ptr *PointerType) TypeParams() []*TypeParam {
	tps := map[string]*TypeParam{}

	{
		_tps := ptr.Elem.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}

	ret := []*TypeParam{}
	for name := range tps {
		ret = append(ret, tps[name])
	}
	slices.SortFunc(ret, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ret
}

func (ptr *PointerType) injectTypeParam(local *Import, tps []*TypeParam) {
	if p, ok := ptr.Elem.(*pseudoType); !ok {
		ptr.Elem.injectTypeParam(local, tps)
	} else {
		ptr.Elem = resolveBareNameType(local, tps, p.Name)
	}
}

func (ptr *PointerType) IsOpaque() bool {
	return ptr.Elem.IsOpaque()
}

type ChanType struct {
	Dir  ast.ChanDir
	Elem Type
}

func (*ChanType) PlainName() string {
	return ""
}

func (ch *ChanType) Expr() string {
	switch ch.Dir {
	case ast.RECV:
		return "<-chan " + ch.Elem.Expr()
	case ast.SEND:
		return "chan<- " + ch.Elem.Expr()
	default:
		return "chan " + ch.Elem.Expr()
	}
}

func (ch *ChanType) Require() []*Import {
	return ch.Elem.Require()
}

func (ch *ChanType) TypeParams() []*TypeParam {
	tps := map[string]*TypeParam{}

	{
		_tps := ch.Elem.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}

	ret := []*TypeParam{}
	for name := range tps {
		ret = append(ret, tps[name])
	}
	slices.SortFunc(ret, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ret
}

func (ch *ChanType) injectTypeParam(local *Import, tps []*TypeParam) {
	if p, ok := ch.Elem.(*pseudoType); !ok {
		ch.Elem.injectTypeParam(local, tps)
	} else {
		ch.Elem = resolveBareNameType(local, tps, p.Name)
	}
}

func (ch *ChanType) IsOpaque() bool {
	return ch.Elem.IsOpaque()
}

type TypeConstraint struct {
	Op   string
	Type Type
}

func (tc *TypeConstraint) PlainName() string {
	return tc.Type.PlainName()
}

func (tc *TypeConstraint) Expr() string {
	return tc.Op + tc.Type.Expr()
}

func (tc *TypeConstraint) Require() []*Import {
	return tc.Type.Require()
}

func (tc *TypeConstraint) TypeParams() []*TypeParam {
	return tc.Type.TypeParams()
}

func (tx *TypeConstraint) injectTypeParam(local *Import, tps []*TypeParam) {
	if p, ok := tx.Type.(*pseudoType); ok {
		tx.Type = resolveBareNameType(local, tps, p.Name)
	} else {
		tx.Type.injectTypeParam(local, tps)
	}
}

func (tc *TypeConstraint) IsOpaque() bool {
	return tc.Type.IsOpaque()
}

type TypeUnion struct {
	Op string
	X  Type
	Y  Type
}

func (uni *TypeUnion) PlainName() string {
	return uni.Expr()
}

func (uni *TypeUnion) Expr() string {
	return fmt.Sprintf("%s | %s", uni.X.Expr(), uni.Y.Expr())
}

func (uni *TypeUnion) Require() []*Import {
	req := map[string]*Import{}
	for _, im := range uni.X.Require() {
		req[im.Path] = im
	}
	for _, im := range uni.Y.Require() {
		req[im.Path] = im
	}
	ret := []*Import{}
	for p := range req {
		ret = append(ret, req[p])
	}
	return ret
}

func (uni *TypeUnion) TypeParams() []*TypeParam {
	req := map[string]*TypeParam{}
	for _, im := range uni.X.TypeParams() {
		req[im.Name] = im
	}
	for _, im := range uni.Y.TypeParams() {
		req[im.Name] = im
	}
	ret := []*TypeParam{}
	for p := range req {
		ret = append(ret, req[p])
	}
	slices.SortFunc(ret, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ret
}

func (uni *TypeUnion) injectTypeParam(local *Import, tps []*TypeParam) {
	if p, ok := uni.X.(*pseudoType); ok {
		uni.X = resolveBareNameType(local, tps, p.Name)
	} else {
		uni.X.injectTypeParam(local, tps)
	}
	if p, ok := uni.Y.(*pseudoType); ok {
		uni.Y = resolveBareNameType(local, tps, p.Name)
	} else {
		uni.Y.injectTypeParam(local, tps)
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

func (fio *FuncIOParam) PlainName() string {
	return fio.Type.PlainName()
}

func (fio *FuncIOParam) Expr() string {
	t := fio.Type.Expr()
	if fio.Variadic {
		t = "..." + t
	}
	return t
}

func (fio *FuncIOParam) ExprWithName(defaultName string) string {
	name := fio.ParamNameOr(defaultName)
	return name + " " + fio.Expr()
}

func (fio *FuncIOParam) Require() []*Import {
	return fio.Type.Require()
}

func (a *FuncIOParam) TypeParams() []*TypeParam {
	tps := map[string]*TypeParam{}

	_tps := a.Type.TypeParams()
	for i := range _tps {
		tp := _tps[i]
		tps[tp.Name] = tp
	}

	ret := []*TypeParam{}
	for name := range tps {
		ret = append(ret, tps[name])
	}
	slices.SortFunc(ret, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ret
}

func (fio *FuncIOParam) injectTypeParam(local *Import, tps []*TypeParam) {
	if p, ok := fio.Type.(*pseudoType); !ok {
		fio.Type.injectTypeParam(local, tps)
	} else {
		fio.Type = resolveBareNameType(local, tps, p.Name)
	}
}

func (fio *FuncIOParam) IsOpaque() bool {
	return fio.Type.IsOpaque()
}

type FuncType struct {
	Args    []*FuncIOParam
	VarArg  *FuncIOParam
	Returns []*FuncIOParam
}

func (*FuncType) PlainName() string {
	return ""
}

func (c *FuncType) Expr() string {
	return "func" + c.Signature(false)
}

func (c *FuncType) Signature(nameSupply bool) string {
	sb := new(strings.Builder)
	io.WriteString(sb, "(")
	params := []string{}
	for i, p := range c.Args {
		params = append(params, p.ExprWithName(fmt.Sprintf("arg%d", i)))
	}
	if c.VarArg != nil {
		params = append(params, c.VarArg.ExprWithName("vararg"))
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
		rets = append(rets, r.ExprWithName(""))
	}
	io.WriteString(sb, strings.Join(rets, ", "))
	if needBrace {
		io.WriteString(sb, ")")
	}

	return sb.String()
}

func (f *FuncType) Require() []*Import {
	req := []*Import{}
	for i := range f.Args {
		req = append(req, f.Args[i].Require()...)
	}
	if f.VarArg != nil {
		req = append(req, f.VarArg.Require()...)
	}
	for i := range f.Returns {
		req = append(req, f.Returns[i].Require()...)
	}
	return req
}

func (fn *FuncType) ParamsGenericExpr(back bool) string {
	_tps := map[string]*TypeParam{}

	for _, para := range fn.Args {
		__tps := para.TypeParams()
		for i := range __tps {
			tp := __tps[i]
			_tps[tp.Name] = tp
		}
	}

	if fn.VarArg != nil {
		__tps := fn.VarArg.TypeParams()
		for i := range __tps {
			tp := __tps[i]
			_tps[tp.Name] = tp
		}
	}

	params := []*TypeParam{}
	for name := range _tps {
		params = append(params, _tps[name])
	}
	slices.SortFunc(params, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})

	tps := []string{}
	for _, p := range params {
		t := p.Expr()
		if back {
			t += " " + p.Back.Expr()
		}
		tps = append(tps, t)
	}
	if len(tps) == 0 {
		return ""
	}
	return "[" + strings.Join(tps, ", ") + "]"
}

func (fn *FuncType) ReturnGenericExpr(back bool) string {
	_tps := map[string]*TypeParam{}

	for _, para := range fn.Returns {
		__tps := para.TypeParams()
		for i := range __tps {
			tp := __tps[i]
			_tps[tp.Name] = tp
		}
	}

	params := []*TypeParam{}
	for name := range _tps {
		params = append(params, _tps[name])
	}
	slices.SortFunc(params, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})

	tps := []string{}
	for _, p := range params {
		t := p.Expr()
		if back {
			t += " " + p.Back.Expr()
		}
		tps = append(tps, t)
	}
	if len(tps) == 0 {
		return ""
	}
	return "[" + strings.Join(tps, ", ") + "]"
}

func (fn *FuncType) TypeParams() []*TypeParam {
	tps := map[string]*TypeParam{}

	for _, para := range fn.Args {
		_tps := para.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}

	if fn.VarArg != nil {
		_tps := fn.VarArg.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}

	for _, r := range fn.Returns {
		_tps := r.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}

	params := []*TypeParam{}
	for name := range tps {
		params = append(params, tps[name])
	}
	slices.SortFunc(params, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return params
}

func (f *FuncType) injectTypeParam(local *Import, tps []*TypeParam) {
	for _, p := range f.Args {
		p.injectTypeParam(local, tps)
	}
	if f.VarArg != nil {
		f.VarArg.injectTypeParam(local, tps)
	}
	for _, r := range f.Returns {
		r.injectTypeParam(local, tps)
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
	Methods []*Method
}

func (*InterfaceType) PlainName() string {
	return ""
}

func (i *InterfaceType) Expr() string {
	sb := new(strings.Builder)
	io.WriteString(sb, "interface{")
	for _, m := range i.Methods {
		io.WriteString(sb, "\n\t")
		io.WriteString(sb, m.Name)
		io.WriteString(sb, m.Func.Signature(false))
	}
	io.WriteString(sb, "\n}")
	return sb.String()
}

func (in *InterfaceType) Require() []*Import {
	req := []*Import{}
	for i := range in.Methods {
		req = append(req, in.Methods[i].Func.Require()...)
	}
	return req
}

func (in *InterfaceType) TypeParams() []*TypeParam {
	tps := map[string]*TypeParam{}

	for _, r := range in.Methods {
		_tps := r.Func.TypeParams()
		for i := range _tps {
			tp := _tps[i]
			tps[tp.Name] = tp
		}
	}

	ret := []*TypeParam{}
	for name := range tps {
		ret = append(ret, tps[name])
	}
	slices.SortFunc(ret, func(a, b *TypeParam) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return ret
}

func (in *InterfaceType) injectTypeParam(local *Import, tps []*TypeParam) {
	for _, meth := range in.Methods {
		meth.Func.injectTypeParam(local, tps)
	}
}

func (in *InterfaceType) IsOpaque() bool {
	for _, m := range in.Methods {
		if m.IsOpaque() {
			return true
		}
	}
	return false
}

type Method struct {
	Name       string
	Func       *FuncType
	isExported bool
}

func (m *Method) IsOpaque() bool {
	return !m.isExported || m.Func.IsOpaque()
}
