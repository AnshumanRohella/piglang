package lexer

import "piglang/token"

type Lexer struct {
	input    string
	currIdx  int
	nextIdx  int
	currChar byte
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() token.Token {

	return token.Token{Type: "", Literal: ""}
}
