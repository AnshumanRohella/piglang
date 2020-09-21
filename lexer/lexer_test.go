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
			t.Fatalf("tests[%d] - tokentype wrong. expected: %q, got: %q",
				i, tc.expectedType, tk.Type)
		}
		if tk.Literal != tc.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected: %q, got: %q",
				i, tc.expectedLiteral, tk.Literal)
		}
	}

}
