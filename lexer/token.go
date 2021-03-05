package lexer

import "fmt"

type TokenType string

const (
	PROGRAM     TokenType = "PROGRAM"
	IDENTIFIER  TokenType = "IDENTIFER"
	SEMI_COLON  TokenType = ";"
	BEGIN       TokenType = "BEGIN"
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"
	STRING      TokenType = "STRING"
	END         TokenType = "END"
	PERIOD      TokenType = "."
	INTEGER     TokenType = "INTEGER"
	EOF         TokenType = "EOF"
	ILLEGAL     TokenType = "ILLEGAL"
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
