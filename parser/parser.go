package parser

import (
	"fmt"
	"io"

	"github.com/hculpan/gopascal/ast"
	"github.com/hculpan/gopascal/lexer"
)

type Parser struct {
	lexer     *lexer.Lexer
	currToken lexer.Token
	nextToken lexer.Token
	Errors    []string
}

func (p *Parser) advanceTokens() {
	p.currToken = p.nextToken
	if lexer.EOF != p.CurrentTokeType() {
		p.nextToken = p.lexer.NextToken()
	} else {
		p.nextToken = p.currToken
	}
}

func (p *Parser) HasErrors() bool {
	return len(p.Errors) > 0
}

func NewParser(r io.Reader) *Parser {
	result := &Parser{
		lexer: lexer.NewLexer(r),
	}
	result.currToken = result.lexer.NextToken()
	result.nextToken = result.lexer.NextToken()
	return result
}

func (p *Parser) Parse() *ast.Program {
	p.Errors = []string{}
	return p.program()
}

func (p *Parser) CurrentTokeType() lexer.TokenType {
	return p.currToken.Type
}

func (p *Parser) NextTokeType() lexer.TokenType {
	return p.nextToken.Type
}

func (p *Parser) addError(msg string, v ...interface{}) {
	p.Errors = append(p.Errors, fmt.Sprintf(msg, v...))
}

func (p *Parser) swallowToken(t lexer.TokenType) bool {
	p.advanceTokens()
	return p.verifyToken(t, p.currToken)
}

func (p *Parser) verifyToken(expectedType lexer.TokenType, t lexer.Token) bool {
	if expectedType != t.Type {
		p.addError("Expecting '%s', found '%s' at %d:%d", expectedType, t.Type, t.Pos.Line, t.Pos.Column)
		return false
	}

	return true
}

func (p *Parser) program() *ast.Program {
	if p.CurrentTokeType() == lexer.EOF {
		return nil
	}

	if !p.verifyToken(lexer.PROGRAM, p.currToken) {
		return nil
	}
	pt := p.currToken

	p.advanceTokens()

	if !p.verifyToken(lexer.IDENTIFIER, p.currToken) {
		return nil
	}

	pn := p.currToken

	if !p.swallowToken(lexer.SEMI_COLON) {
		return nil
	}

	block := p.block()

	return ast.NewProgramNode(pt, pn, block)
}

func (p *Parser) block() *ast.Block {
	p.advanceTokens()

	if p.CurrentTokeType() == lexer.EOF {
		return nil
	}

	if !p.verifyToken(lexer.BEGIN, p.currToken) {
		return nil
	}

	p.advanceTokens()

	block := ast.NewBlock(p.currToken)

	block.Statements = p.statementSequence()

	if !p.verifyToken(lexer.END, p.currToken) {
		return nil
	}

	return block
}

func (p *Parser) statementSequence() []*ast.ProcedureStatement {
	result := []*ast.ProcedureStatement{}

	for p.CurrentTokeType() != lexer.END {
		result = append(result, p.statement())
	}

	return result
}

func (p *Parser) statement() *ast.ProcedureStatement {
	if !p.verifyToken(lexer.IDENTIFIER, p.currToken) {
		return nil
	}

	result := ast.ProcedureStatement{
		ProcedureName: p.currToken,
	}

	if !p.swallowToken(lexer.LEFT_PAREN) {
		return nil
	}

	for {
		if lexer.RIGHT_PAREN == p.NextTokeType() { // done reading arguments
			break
		}
	}

	return &result
}
