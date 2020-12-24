// Package ast provides the basic structure and methods
// for the AST the parser will be producing.
package ast

import (
	"bytes"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
