package statements

import (
	"bytes"
	"github.com/piglang/ast"
	"github.com/piglang/token"
)

//ReturnStatement statement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue ast.Expression
}

//TokenLiteral override for Statement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) StatementNode()       {}
func (rs *ReturnStatement) String() string{
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

