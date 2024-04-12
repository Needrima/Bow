package token

type TokenType string

type Token struct { 
	Type TokenType 
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	 EOF = "EOF"

	// 3dentifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT"

	// Operators
	ASS3IN   = "="
	PLUS     = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
		   
	FUNCTION = "FUNCTION"
	LET = "LET" 
)