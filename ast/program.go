package ast

import "github.com/hculpan/gopascal/lexer"

type Program struct {
	ProgramToken lexer.Token
	ProgramName  lexer.Token
	Statements   []*Statement
}

func (p *Program) StringLiteral() string {
	return p.ProgramToken.Value
}

func NewProgramNode(token lexer.Token, nameToken lexer.Token, statements []*Statement) *Program {
	return &Program{
		ProgramToken: token,
		ProgramName:  nameToken,
		Statements:   statements,
	}
}
