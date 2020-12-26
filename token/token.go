package token

type TokenType string

//Token is the basic structure for each token in piglang. The TokenType represents the various tokens in the language and the Literal represents the exact token literal.
type Token struct {
	Type    TokenType
	Literal string
}

var keywordMap = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

const (
	//ILLEGAL token
	ILLEGAL = "ILLEGAL"
	//EOF END-OF-FILE token
	EOF = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	EQ       = "=="
	NOT_EQ   = "!="
)

func GetIdentifierType(identifier string) TokenType {
	if tok, ok := keywordMap[identifier]; ok {
		return tok
	}
	return IDENT
}
