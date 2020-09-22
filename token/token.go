package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywordMap = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

const (
	//ILLEGAL token
	ILLEGAL = "ILLEGAL"
	//EOF END-OF-FILE token
	EOF = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

func GetIdentifierType(identifier string) TokenType {
	if tok, ok := keywordMap[identifier]; ok {
		return tok
	}
	return IDENT
}
