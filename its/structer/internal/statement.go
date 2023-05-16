package internal

import (
	"fmt"
	"go/ast"
	"io"
	"strings"
)

type File struct {
	PackageName string
	Imports     []*Import
	Structs     []*StructDecl
}

type Import struct {
	Name string
	Path string
}

type StructDecl struct {
	Name       string
	Package    *Import
	TypeParams []*TypeParam
	Body       *StructType
}

func (s *StructDecl) PlainName() string {
	return s.Name
}

func (s *StructDecl) GenericExpr(backtype bool) string {
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

func (s *StructDecl) Expr() string {
	pkg := ""
	if s.Package != nil {
		pkg = s.Package.Name + "."
	}
	return pkg + s.Name + s.GenericExpr(false)
}

func (s *StructDecl) AsType() Type {
	return &NamedType{
		Pkg:  s.Package,
		Name: s.Name,
	}
}

func (s *StructDecl) Require() []*Import {
	req := []*Import{s.Package}
	for i := range s.TypeParams {
		req = append(req, s.TypeParams[i].Back.Require()...)
	}
	req = append(req, s.Body.Require()...)
	return req
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

type Type interface {
	PlainName() string
	Expr() string
	Require() []*Import
}

type NamedType struct {
	Pkg  *Import
	Name string
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

type Field struct {
	Name string
	Type Type
}

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

type PointerType struct {
	Elem Type
}

func (p *PointerType) PlainName() string {
	return p.Elem.PlainName()
}

func (p *PointerType) Expr() string {
	return "*" + p.Elem.Expr()
}

func (p *PointerType) Require() []*Import {
	return p.Elem.Require()
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

type FuncType struct {
	Params        []Type
	EllipsisParam Type
	Results       []Type
}

func (*FuncType) PlainName() string {
	return ""
}

func (c *FuncType) Expr() string {
	return "func" + c.Signature()
}

func (c *FuncType) Signature() string {
	sb := new(strings.Builder)
	io.WriteString(sb, "(")
	params := []string{}
	for _, p := range c.Params {
		params = append(params, p.Expr())
	}
	if c.EllipsisParam != nil {
		params = append(params, "..."+c.EllipsisParam.Expr())
	}
	io.WriteString(sb, strings.Join(params, ", "))
	io.WriteString(sb, ")")
	if nret := len(c.Results); 1 < nret {
		io.WriteString(sb, " (")
	} else if 0 < nret {
		io.WriteString(sb, " ")
	}
	rets := []string{}
	for _, r := range c.Results {
		rets = append(rets, r.Expr())
	}
	io.WriteString(sb, strings.Join(rets, ", "))
	if 1 < len(c.Results) {
		io.WriteString(sb, ")")
	}

	return sb.String()
}

func (f *FuncType) Require() []*Import {
	req := []*Import{}
	for i := range f.Params {
		req = append(req, f.Params[i].Require()...)
	}
	if f.EllipsisParam != nil {
		req = append(req, f.EllipsisParam.Require()...)
	}
	for i := range f.Results {
		req = append(req, f.Results[i].Require()...)
	}
	return req
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
		io.WriteString(sb, m.Signature.Signature())
	}
	io.WriteString(sb, "\n}")
	return sb.String()
}

func (in *InterfaceType) Require() []*Import {
	req := []*Import{}
	for i := range in.Methods {
		req = append(req, in.Methods[i].Signature.Require()...)
	}
	return req
}

type Method struct {
	Name      string
	Signature *FuncType
}
