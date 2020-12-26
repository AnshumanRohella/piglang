package parser

import (
	"fmt"
	"github.com/piglang/ast"
	"github.com/piglang/ast/expressions"
	"github.com/piglang/ast/literals"
	"github.com/piglang/ast/statements"
	"github.com/piglang/lexer"
	"github.com/piglang/token"
	"strconv"
)

//Expression operator precedence
const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func() ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func (p *Parser) registerPrefixParseFn(t token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}

func (p *Parser) registerInfixParseFn(t token.TokenType, fn infixParseFn) {
	p.infixParseFns[t] = fn
}

func New(l *lexer.Lexer) *Parser {

	// This may cause problems. use &Parser.
	p := &Parser{l: l, errors: []string{}}

	// initialize current and peek token both.
	p.nextToken()
	p.nextToken()

	// initialize expression parsing functions map
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	// Register parsing functions.
	p.registerPrefixParseFn(token.IDENT, p.parseIdentifier)
	p.registerPrefixParseFn(token.INT, p.parseIntegerLiteral)
	p.registerPrefixParseFn(token.BANG,p.parsePrefixExpressions)
	p.registerPrefixParseFn(token.MINUS,p.parsePrefixExpressions)

	return p
}

// Expression Parsing functions.
// Parse Identifier
func (p *Parser) parseIdentifier() ast.Expression {
	return &statements.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// Parse Integer Literals by converting the string literal to integer.
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &literals.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 62)
	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) nextToken() {

	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()

}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) addPeekError(t token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead.", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
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

// Parse various statements.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatements() //Expressions are the default statements after all the cases.
	}

}

func (p *Parser) parseExpressionStatements() *statements.ExpressionStatement {

	stmt := &statements.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Checks whether we have a function registered for current token at prefix position.
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefixFn := p.prefixParseFns[p.curToken.Type]

	if prefixFn == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefixFn()

	return leftExp

}

func (p *Parser) noPrefixParseFnError(t token.TokenType){
	msg := fmt.Sprintf("no prefix parse function found for token %s", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseReturnStatement() *statements.ReturnStatement {
	stmt := &statements.ReturnStatement{Token: p.curToken}
	p.nextToken()

	//TODO: Skip expression until semicolon for now
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Parsers for acual prefix expressions. This will be called in case the prefix is one of the supported prefixes in PrefixExpression.go
func (p *Parser) parsePrefixExpressions() ast.Expression {
	exp := &expressions.PrefixExpression{
		Token: p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	exp.RightOperand = p.parseExpression(PREFIX)

	return exp
}

func (p *Parser) parseLetStatement() *statements.LetStatement {
	stmt := new(statements.LetStatement)
	stmt.Token = p.curToken

	if !p.expectPeekAndProceed(token.IDENT) {
		return nil
	}

	stmt.Name = &statements.Identifier{Token: p.curToken, Value: p.curToken.Literal}

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
	p.addPeekError(t)
	return false
}
