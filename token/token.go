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
	// MINUS is a minus operator
	MINUS = "-"
	// BANG is a bang operator
	BANG = "!"
	// ASTERIK is an asterik operator
	ASTERIK = "*"
	// SLASH is a slash operator
	SLASH = "/"
	// LT is a less than operator
	LT = "<"
	// GT is a greater than operator
	GT = ">"
	// EQ is an equality operator
	EQ = "=="
	// NEQ is a non equality operator
	NEQ = "!="

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
	// TRUE is a true keyword
	TRUE = "TRUE"
	// FALSE is a false keyword
	FALSE = "FALSE"
	// IF is a if keyword
	IF = "IF"
	// ELSE is a else keyword
	ELSE = "ELSE"
	// RETURN is a return keyword
	RETURN = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent lookups up keyword of identifier if exists
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
