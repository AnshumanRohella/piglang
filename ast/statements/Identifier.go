package statements

import "github.com/piglang/token"

//Identifier to repesent an identifier token.
type Identifier struct {
	Token token.Token
	Value string
}

//TokenLiteral override for Statement
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

// Identifier can precede a value in some cases. Eg. x=y
func (i *Identifier) ExpressionNode() {}

