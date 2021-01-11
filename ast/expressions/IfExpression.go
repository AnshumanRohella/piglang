package expressions

import (
	"bytes"
	"github.com/piglang/ast"
	"github.com/piglang/ast/statements"
	"github.com/piglang/token"
)

type IfExpresion struct {
	Token token.Token
	Condition ast.Expression
	Consequence *statements.BlockStatement
	Alternative *statements.BlockStatement
}

func (ie *IfExpresion) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpresion) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}
func (ie *IfExpresion) ExpressionNode() {}
