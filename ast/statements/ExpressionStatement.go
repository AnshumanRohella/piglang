package statements

import (
	"github.com/piglang/ast"
	"github.com/piglang/token"
)

//ExpressionStatement statement
type ExpressionStatement struct {
	Token      token.Token
	Expression ast.Expression
}

//TokenLiteral override for Statement
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) StatementNode()       {}
func (es *ExpressionStatement) String() string {

	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
