package ast

import "github.com/hculpan/gopascal/lexer"

type Block struct {
	BeginToken lexer.Token
	EndToken   lexer.Token
	Statements []*Statement
}

func (b *Block) StringLiteral() string {
	return b.BeginToken.Value
}
