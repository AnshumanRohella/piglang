package parser

import (
	"github.com/piglang/ast"
	"github.com/piglang/lexer"
	"github.com/piglang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {

	// This may cause problems. use &Parser.
	p := new(Parser)
	p.l = l

	// initialize current and peek token both.
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {

	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()

}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
