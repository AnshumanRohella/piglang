package lexer

import (
	"piglang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	lexer := New(input)

	for i, tc := range tests {
		tk := lexer.NextToken()
		if tk.Type != tc.expectedType {
			t.Fatalf("fucked at %d", i)
		}
		if tk.Literal != tc.expectedLiteral {
			t.Fatalf("fucked at %d", i)
		}
	}

}
