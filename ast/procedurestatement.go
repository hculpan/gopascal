package ast

import "github.com/hculpan/gopascal/lexer"

type ProcedureStatement struct {
	ProcedureName lexer.Token
}

func (p *ProcedureStatement) StringLiteral() string {
	return p.ProcedureName.Value
}

func NewProcStatement(p lexer.Token) *ProcedureStatement {
	return &ProcedureStatement{
		ProcedureName: p,
	}
}
