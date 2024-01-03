package ast

import (
	"bytes"
	"chimp/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type IntStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (is *IntStatement) statementNode()       {}
func (is *IntStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IntStatement) String() string {
	var out bytes.Buffer

	out.WriteString(is.TokenLiteral() + " ")
	out.WriteString(is.Name.String())
	out.WriteString(" = ")

	if is.Value != nil {
		out.WriteString(is.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type BoolStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (bs *BoolStatement) statementNode()       {}
func (bs *BoolStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BoolStatement) String() string {
	var out bytes.Buffer

	out.WriteString(bs.TokenLiteral() + " ")
	out.WriteString(bs.Name.String())
	out.WriteString(" = ")

	if bs.Value != nil {
		out.WriteString(bs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type StringStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ss *StringStatement) statementNode()       {}
func (ss *StringStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *StringStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ss.TokenLiteral() + " ")
	out.WriteString(ss.Name.String())
	out.WriteString(" = ")

	if ss.Value != nil {
		out.WriteString(ss.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}
