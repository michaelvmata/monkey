package token

// Type is a human reable value of the token type
type Type string

// Token houses the raw literal value and type
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL is an unrecognized token type
	ILLEGAL = "ILLEGAL"
	// EOF is the end of file
	EOF = "EOF"

	// IDENT is an identifier
	IDENT = "IDENT"
	// INT is a literal integer
	INT = "INT"

	// ASSIGN is an assignment operator
	ASSIGN = "="
	// PLUS is a plus operator
	PLUS = "+"

	// COMMA is a comma delimiter
	COMMA = ","
	// SEMICOLON is a comma delimiter
	SEMICOLON = ";"

	// LPAREN is self explanatory
	LPAREN = "("
	// RPAREN is self explanatory
	RPAREN = ")"
	// LBRACE is self explanatory
	LBRACE = "{"
	// RBRACE is self explanatory
	RBRACE = "}"

	// FUNCTION is a function keyword
	FUNCTION = "FUNCTION"
	// LET is a let keyword
	LET = "LET"
)
