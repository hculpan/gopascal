package lexer

import "fmt"

type TokenType string

const (
	PROGRAM     = "PROGRAM"
	IDENTIFIER  = "IDENTIFER"
	BUILT_IN    = "BUILT_IN"
	SEMI_COLON  = ";"
	BEGIN       = "BEGIN"
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	STRING      = "STRING"
	END         = "END"
	PERIOD      = "."
	INTEGER     = "INTEGER"
	EOF         = "EOF"
	ILLEGAL     = "ILLEGAL"
)

type Position struct {
	Line   int
	Column int
}

type Token struct {
	Type  TokenType
	Value string
	Pos   Position
}

func NewToken(tokenType TokenType, value string, position Position) Token {
	return Token{
		Type:  tokenType,
		Value: value,
		Pos:   position,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("Type: %s, value: %s [%d:%d]", t.Type, t.Value, t.Pos.Line, t.Pos.Column)
}

func (t Token) IsIllegal() bool {
	return t.Type == ILLEGAL
}
