package ast

import (
	"github.com/michaelvmata/monkey/token"
)

// Node base
type Node interface {
	TokenLiteral() string
}

// Statement node
type Statement interface {
	Node
	statementNode()
}

// Expression node
type Expression interface {
	Node
	exressionNode()
}

// Program is a collection of statements
type Program struct {
	Statements []Statement
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

func (i *Identifier) statementNode() {}

// TokenLiteral returns the token literal value for the identifier
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// LetStatement node
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}

// TokenLiteral returns the token literal value for the let statement
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }

// ReturnStatement node
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (l *ReturnStatement) statementNode() {}

// TokenLiteral returns the token literal value for the let statement
func (l *ReturnStatement) TokenLiteral() string { return l.Token.Literal }
