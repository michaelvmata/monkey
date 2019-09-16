package ast

import (
	"github.com/michaelvmata/monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "vegita"},
					Value: "vegita",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "PrinceOfAllSayians"},
					Value: "PrinceOfAllSayians",
				},
			},
		},
	}
	if program.String() != "let vegita = PrinceOfAllSayians;" {
		t.Errorf("program.String() is wrong. got=%q", program.String())
	}

}
