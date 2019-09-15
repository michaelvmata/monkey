package parser

import (
	"github.com/michaelvmata/monkey/ast"
	"github.com/michaelvmata/monkey/lexer"
	"github.com/michaelvmata/monkey/token"
)

// Parser is used to parse tokens into an AST
type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses token into an AST
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

// New returns new instance of Parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}
