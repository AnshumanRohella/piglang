package literals

import (
	"github.com/piglang/token"
)
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (li *IntegerLiteral) ExpressionNode(){}
func (li *IntegerLiteral) TokenLiteral() string {return li.Token.Literal }
func (li* IntegerLiteral) String() string { return li.Token.Literal }