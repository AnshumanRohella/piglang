package lexer

import (
	"testing"

	"github.com/piglang/token"
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

func TestNextTokenExtended(t *testing.T) {
	input := `let five = 5;
	let four = 4;
	
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, four);
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "four"},
		{token.ASSIGN, "="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "four"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
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
