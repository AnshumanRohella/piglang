package statements

import (
	"bytes"
	"github.com/piglang/ast"
	"github.com/piglang/token"
)

//LetStatement representing the let statement.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value ast.Expression
}

func (ls *LetStatement) StatementNode() {}

//TokenLiteral implementation for Node
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
