// Package ast provides the basic structure and methods
// for the AST the parser will be producing.
package ast

import "github.com/piglang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

//Program is the root node of every AST.
//Every valid piglang program is a list of statements.
type Program struct {
	Statements []Statement
}

//TokenLiteral function for Program returns the token literal for the first statement.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Statement Types here

//LetStatement representing the let statement.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

//TokenLiteral implementation for Node
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

//Identifier to repesent an identifier token.
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Indentifier can procude a value in some cases. Eg. x=y
func (i *Identifier) expressionNode() {}

//Return statement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) statementNode()       {}
