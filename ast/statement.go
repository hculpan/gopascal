package ast

import "github.com/hculpan/gopascal/lexer"

type Statement struct {
	FirstToken lexer.Token
}

func (s *Statement) StringLiteral() string {
	return s.FirstToken.Value
}

func NewStatement(t lexer.Token) *Statement {
	return &Statement{
		FirstToken: t,
	}
}
