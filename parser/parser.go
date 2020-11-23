package parser

import (
	"fmt"
	"io"

	"github.com/hculpan/gopascal/ast"
	"github.com/hculpan/gopascal/lexer"
)

type Parser struct {
	lexer        *lexer.Lexer
	buffer       lexer.Token
	lastToken    lexer.Token
	bufferFilled bool
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		lexer:        lexer.NewLexer(r),
		bufferFilled: false,
	}
}

func (p *Parser) Parse() (*ast.Program, []string) {
	/*	_, errors := p.lexer.Lex()

		if len(errors) > 0 {
			return nil, errors
		}
	*/
	return p.program()
}

func (p *Parser) scan() lexer.Token {
	var result lexer.Token
	if p.bufferFilled {
		result = p.buffer
		p.bufferFilled = false
	} else {
		result = p.lexer.NextToken()
		p.lastToken = result
	}
	return result
}

func (p *Parser) unscan() {
	if !p.bufferFilled {
		p.buffer = p.lastToken
		p.bufferFilled = true
	}
}

func verifyToken(expectedType lexer.TokenType, t lexer.Token) string {
	if expectedType != t.Type {
		return fmt.Sprintf("Expecting '%s', found '%s' at %d:%d", expectedType, t.Type, t.Pos.Line, t.Pos.Column)
	}

	return "Ok"
}

func IsOk(err string) bool {
	return err == "Ok"
}

func IsNotOk(err string) bool {
	return err != "Ok"
}

func (p *Parser) program() (*ast.Program, []string) {
	pt := p.scan()
	if pt.Type == lexer.EOF {
		return nil, []string{}
	}

	err := verifyToken(lexer.PROGRAM, pt)
	if IsNotOk(err) {
		return nil, []string{err}
	}
	pn := p.scan()
	err = verifyToken(lexer.IDENTIFIER, pn)
	if IsNotOk(err) {
		return nil, []string{err}
	}
	err = verifyToken(lexer.SEMI_COLON, p.scan())
	if IsNotOk(err) {
		return nil, []string{err}
	}

	block, errors := p.block()

	return ast.NewProgramNode(pt, pn, block), errors
}

func (p *Parser) block() (*ast.Block, []string) {
	pt := p.scan()
	if pt.Type == lexer.EOF {
		return nil, []string{}
	}

	err := verifyToken(lexer.BEGIN, pt)
	if IsNotOk(err) {
		return nil, []string{err}
	}

	block := ast.NewBlock(pt)

	pt = p.scan()
	err = verifyToken(lexer.END, pt)
	if IsNotOk(err) {
		return nil, []string{err}
	}

	return block, []string{}
}
