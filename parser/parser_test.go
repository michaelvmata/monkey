package parser

import (
	"github.com/michaelvmata/monkey/lexer"
	"testing"
)

func Test(t *testing.T) {
	input := `5 + 5`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program != nil {
		t.Fatalf("ParseProgram() is not implemented yet")
	}
}
