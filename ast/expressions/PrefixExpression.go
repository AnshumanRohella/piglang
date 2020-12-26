package expressions

import (
	"bytes"
	"github.com/piglang/ast"
	"github.com/piglang/token"
)

type PrefixExpression struct {
	Token token.Token
	Operator string // Right now the supported prefix operators are ! and -
	RightOperand ast.Expression
}

func (pe *PrefixExpression) TokenLiteral() string {return pe.Token.Literal }
func (pe *PrefixExpression) ExpressionNode() { }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.RightOperand.String())
	out.WriteString(")")

	return out.String()
}