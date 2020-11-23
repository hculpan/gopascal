package ast

import "github.com/hculpan/gopascal/lexer"

type Program struct {
	ProgramToken lexer.Token
	ProgramName  lexer.Token
	Block        *Block
}

func (p *Program) StringLiteral() string {
	return p.ProgramToken.Value
}

func NewProgramNode(token lexer.Token, nameToken lexer.Token, block *Block) *Program {
	return &Program{
		ProgramToken: token,
		ProgramName:  nameToken,
		Block:        block,
	}
}
