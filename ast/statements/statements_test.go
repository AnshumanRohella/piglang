package statements


import (
	"github.com/piglang/ast"
	"github.com/piglang/token"
	"testing"
)


func TestString(t *testing.T){
	program := &ast.Program{
		Statements: []ast.Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "testVar"},
					Value: "testVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "testVar2"},
					Value: "testVar2",
				},
			},
		},
	}
	if program.String() != "let testVar = testVar2;"{
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}