package lexer

import (
	"fmt"

	"github.com/piglang/token"
)

type Lexer struct {
	input    string
	currIdx  int
	nextIdx  int
	currChar byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.eatWhiteSpace()
	switch l.currChar {
	case '=':
		tok = newToken(token.ASSIGN, l.currChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currChar)
	case '(':
		tok = newToken(token.LPAREN, l.currChar)
	case ')':
		tok = newToken(token.RPAREN, l.currChar)
	case ',':
		tok = newToken(token.COMMA, l.currChar)
	case '+':
		tok = newToken(token.PLUS, l.currChar)
	case '{':
		tok = newToken(token.LBRACE, l.currChar)
	case '}':
		tok = newToken(token.RBRACE, l.currChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.currChar) {
			tok.Literal = l.readIdentifier()
			fmt.Printf("Got %s ", tok.Literal)
			tok.Type = token.GetIdentifierType(tok.Literal)
			return tok
		} else if isDigit(l.currChar) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currChar)
		}
	}
	l.readChar()
	return tok
}
func (l *Lexer) eatWhiteSpace() {
	for l.currChar == ' ' || l.currChar == '\t' || l.currChar == '\n' || l.currChar == '\r' {
		l.readChar()
	}
}
func (l *Lexer) readIdentifier() string {
	pos := l.currIdx
	for isLetter(l.currChar) {
		l.readChar()
	}
	return l.input[pos:l.currIdx]
}

func (l *Lexer) readNumber() string {
	pos := l.currIdx
	for isDigit(l.currChar) {
		l.readChar()
	}
	return l.input[pos:l.currIdx]
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (l *Lexer) readChar() {
	if l.nextIdx >= len(l.input) {
		l.currChar = 0
	} else {
		l.currChar = l.input[l.nextIdx]
	}
	l.currIdx = l.nextIdx
	l.nextIdx += 1

}
