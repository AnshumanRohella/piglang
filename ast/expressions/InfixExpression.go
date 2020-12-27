package expressions

import (
	"bytes"
	"github.com/piglang/ast"
	"github.com/piglang/token"
)

type InfixExpression struct {
	Token        token.Token
	LeftOperand  ast.Expression
	Operator     string
	RightOperand ast.Expression
}

func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.LeftOperand.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.RightOperand.String())
	out.WriteString(")")

	return out.String()
}
func (ie *InfixExpression) ExpressionNode() {}
