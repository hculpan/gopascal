package ast

import "github.com/hculpan/gopascal/lexer"

type Block struct {
	BeginToken lexer.Token
	EndToken   lexer.Token
	Statements []*ProcedureStatement
}

func (b *Block) StringLiteral() string {
	return b.BeginToken.Value
}

func NewBlock(bt lexer.Token) *Block {
	return &Block{
		BeginToken: bt,
	}
}
