package ast

import (
	"bytes"
	"github.com/michaelvmata/monkey/token"
	"strings"
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

// IntegerLiteral node
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}

// String returns the string representation of IntegerLiteral
func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

// TokenLiteral returns the token literal value for the integer
func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }

// PrefixExpression node
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (p *PrefixExpression) expressionNode() {}

// String returns the string representation of PrefixExpression
func (p *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")
	return out.String()
}

// TokenLiteral returns the token literal value for PrefixExpression
func (p *PrefixExpression) TokenLiteral() string { return p.Token.Literal }

// InfixExpression node
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}

func (i *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.Right.String())
	out.WriteString(")")
	return out.String()
}

// TokenLiteral returns the token literal value for InfixExpression
func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}

// Boolean node
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) String() string { return b.Token.Literal }

// TokenLiteral returns the token literal value for Boolean
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

// BlockStatement node
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (b *BlockStatement) statementNode() {}

// TokenLiteral returns the token literal value for BlockStatement
func (b *BlockStatement) TokenLiteral() string { return b.Token.Literal }

func (b *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range b.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// IfExpression node
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IfExpression) expressionNode() {}

// TokenLiteral returns the token literal value for IfExpression
func (i *IfExpression) TokenLiteral() string { return i.Token.Literal }

func (i *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(i.Condition.String())
	out.WriteString(" ")
	out.WriteString(i.Consequence.String())
	if i.Alternative != nil {
		out.WriteString("else")
		out.WriteString(i.Alternative.String())
	}
	return out.String()
}

// FunctionLiteral node
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *FunctionLiteral) expressionNode() {}

// TokenLiteral returns the token literal value for FunctionLiteral
func (f *FunctionLiteral) TokenLiteral() string { return f.Token.Literal }

func (f *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(f.Body.String())
	return out.String()
}

// CallExpression node
type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (c *CallExpression) expressionNode() {}

func (c *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range c.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(c.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ","))
	out.WriteString(")")
	return out.String()
}

// TokenLiteral returns the token literal value for CallExpression
func (c *CallExpression) TokenLiteral() string { return c.Token.Literal }
