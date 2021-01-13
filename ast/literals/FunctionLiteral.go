package literals

import (
	"bytes"
	"github.com/piglang/ast/statements"
	"github.com/piglang/token"
	"strings"
)

type FunctionLiteral struct {
	Token token.Token
	Arguments []*statements.Identifier
	Body *statements.BlockStatement
}
func (fl *FunctionLiteral) ExpressionNode(){}
func (fl *FunctionLiteral) TokenLiteral() string {return fl.Token.Literal }
func (fl* FunctionLiteral) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, p := range fl.Arguments{
		args = append(args, p.String())
	}
	out.WriteString("fn")
	out.WriteString(" ( ")
	out.WriteString(strings.Join(args,", "))
	out.WriteString(" ) ")
	out.WriteString(fl.Body.String())

	return out.String()
}
