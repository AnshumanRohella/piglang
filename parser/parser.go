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
	program := new(ast.Program)
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// real deal
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}

}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := new(ast.LetStatement)
	stmt.Token = p.curToken

	if !p.expectPeekAndProceed(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeekAndProceed(token.ASSIGN) {
		return nil
	}

	// limit to left side parsing for now.
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeekAndProceed(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}
