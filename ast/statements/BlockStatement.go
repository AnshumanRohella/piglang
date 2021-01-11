package statements

import (
	"bytes"
	"github.com/piglang/ast"
	"github.com/piglang/token"
)

type BlockStatement struct {
	Token token.Token // token is { here
	Statements[] ast.Statement
}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) StatementNode()       {}
func (bs *BlockStatement) String() string{
	var out bytes.Buffer

	for _, st := range bs.Statements {
		out.WriteString(st.String())
	}
	return out.String()
}
