package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hculpan/gopascal/lexer"
)

func TestHello1(t *testing.T) {
	parser := NewParser(strings.NewReader("program hello;"))
	program := parser.Parse()
	if parser.HasErrors() {
		for _, v := range parser.Errors {
			fmt.Println(v)
		}
		t.Errorf("Found %d errors in parsing", len(parser.Errors))
	}
	if program == nil {
		t.Error("Invalid program node returned")
	}
}

/*
* Tests simple scanning
 */
func TestParser1(t *testing.T) {
	parser := NewParser(strings.NewReader("program hello;"))
	VerifyToken(lexer.PROGRAM, parser.currToken, t)
	VerifyTokenWithValue(lexer.IDENTIFIER, "hello", parser.nextToken, t)
	parser.advanceTokens()
	VerifyTokenWithValue(lexer.IDENTIFIER, "hello", parser.currToken, t)
	VerifyToken(lexer.SEMI_COLON, parser.nextToken, t)
	parser.advanceTokens()
	VerifyToken(lexer.SEMI_COLON, parser.currToken, t)
	VerifyToken(lexer.EOF, parser.nextToken, t)
	parser.advanceTokens()
	VerifyToken(lexer.EOF, parser.currToken, t)
	VerifyToken(lexer.EOF, parser.nextToken, t)
}

func VerifyToken(expectedType lexer.TokenType, token lexer.Token, t *testing.T) {
	if token.Type == lexer.ILLEGAL {
		t.Errorf("Illegal token '%s' [%d:%d]", token.Type, token.Pos.Line, token.Pos.Column)
	} else if expectedType != token.Type {
		t.Errorf("Expected %s token, got '%s' [%d:%d]", expectedType, token.Value, token.Pos.Line, token.Pos.Column)
	}
}

func VerifyTokenWithValue(expectedType lexer.TokenType, expectedValue string, token lexer.Token, t *testing.T) {
	if token.Type == lexer.ILLEGAL {
		t.Errorf("Illegal token '%s' [%d:%d]", token.Type, token.Pos.Line, token.Pos.Column)
	} else if expectedType != token.Type {
		t.Errorf("Expected %s token, got '%s' [%d:%d]", expectedType, token.Value, token.Pos.Line, token.Pos.Column)
	} else if expectedValue != token.Value {
		t.Errorf("Expected '%s', got '%s' [%d:%d]", expectedValue, token.Value, token.Pos.Line, token.Pos.Column)
	}
}
