package ast

import (
	"bytes"
	"github.com/michaelvmata/monkey/token"
)

// Node base
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement node
type Statement interface {
	Node
	statementNode()
}

// Expression node
type Expression interface {
	Node
	expressionNode()
}

// Program is a collection of statements
type Program struct {
	Statements []Statement
}

// String is the string representation of Program
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// TokenLiteral returns the token literal of th eroot node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Identifier node
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// String return the string representation of Identifier
func (i *Identifier) String() string { return i.Value }

// TokenLiteral returns the token literal value for the identifier
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// LetStatement node
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}

// String is the string representation of LetStatement
func (l *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(l.TokenLiteral() + " " + l.Name.String() + " = ")
	if l.Value != nil {
		out.WriteString(l.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// TokenLiteral returns the token literal value for the let statement
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }

// ReturnStatement node
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (r *ReturnStatement) statementNode() {}

// String returns the stirng representation of ReturnStatement
func (r *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(r.TokenLiteral() + " ")
	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// TokenLiteral returns the token literal value for the let statement
func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }

// ExpressionStatement node
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {}

// String returns the string representation of ExpressionStatement
func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}

// TokenLiteral returns the token literal value for the let statement
func (e *ExpressionStatement) TokenLiteral() string { return e.Token.Literal }
