package literals

import "github.com/piglang/token"

type BooleanLiteral struct {
	Token token.Token
	Value bool
}
func (li *BooleanLiteral) ExpressionNode(){}
func (li *BooleanLiteral) TokenLiteral() string {return li.Token.Literal }
func (li* BooleanLiteral) String() string { return li.Token.Literal }