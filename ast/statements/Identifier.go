package statements

import "github.com/piglang/token"

//Identifier to represent an identifier token. The value here comes from the token.Literal
type Identifier struct {
	Token token.Token
	Value string
}

//TokenLiteral override for Statement
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

// ExpressionNode implementation for Expression interface. Just like any expression identifiers can produce a value too. eg x=y.
func (i *Identifier) ExpressionNode() {}

