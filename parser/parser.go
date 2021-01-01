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

//Precedence map for token types
var tokenPrecedenceMap = map[token.TokenType]int {
	token.EQ: EQUALS,
	token.NOT_EQ: EQUALS,
	token.GT: LESSGREATER,
	token.LT: LESSGREATER,
	token.ASTERISK: PRODUCT,
	token.SLASH: PRODUCT,
	token.PLUS: SUM,
	token.MINUS: SUM,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := tokenPrecedenceMap[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}


func (p *Parser) currPrecedence() int {
	if p, ok := tokenPrecedenceMap[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}


type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(expression ast.Expression) ast.Expression
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
	p.registerPrefixParseFn(token.TRUE, p.parseBoolean)
	p.registerPrefixParseFn(token.FALSE, p.parseBoolean)
	p.registerPrefixParseFn(token.LPAREN, p.parseGroupedExpression)

	// register infix parser
	p.registerInfixParseFn(token.PLUS, p.parseInfixExpression)
	p.registerInfixParseFn(token.MINUS, p.parseInfixExpression)
	p.registerInfixParseFn(token.ASTERISK, p.parseInfixExpression)
	p.registerInfixParseFn(token.SLASH, p.parseInfixExpression)
	p.registerInfixParseFn(token.GT, p.parseInfixExpression)
	p.registerInfixParseFn(token.LT, p.parseInfixExpression)
	p.registerInfixParseFn(token.EQ, p.parseInfixExpression)
	p.registerInfixParseFn(token.NOT_EQ, p.parseInfixExpression)

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

	// process the tokens from left to right while the precedence is lower and the token is not a semi-colon.
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}

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

// Parser for infix expression. The function takes the left expression as an argument, gets the precedence for the current token and continues to build the right part.
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &expressions.InfixExpression{
		Token: p.curToken,
		LeftOperand: left,
		Operator: p.curToken.Literal,
	}

	precedence := p.currPrecedence()
	p.nextToken()
	exp.RightOperand = p.parseExpression(precedence)

	return exp

}

// Prefix parsing function for boolean
func (p *Parser) parseBoolean() ast.Expression {
	return &literals.BooleanLiteral{Token: p.curToken,Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) parseGroupedExpression() ast.Expression{
	// move to the next token after the brace
	p.nextToken()

	// start parsing like a normal expression
	exp := p.parseExpression(LOWEST)

	if !p.expectPeekAndProceed(token.RPAREN) {
		return nil
	}
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
