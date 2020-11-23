package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hculpan/gopascal/lexer"
)

func TestHello1(t *testing.T) {
	parser := NewParser(strings.NewReader("program hello;"))
	program, errors := parser.Parse()
	if len(errors) > 0 {
		for _, v := range errors {
			fmt.Println(v)
		}
		t.Errorf("Found %d errors in parsing", len(errors))
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
	var token lexer.Token = parser.scan()
	VerifyToken(lexer.PROGRAM, token, t)
	token = parser.scan()
	VerifyTokenWithValue(lexer.IDENTIFIER, "hello", token, t)
	token = parser.scan()
	VerifyToken(lexer.SEMI_COLON, token, t)
	token = parser.scan()
	VerifyToken(lexer.EOF, token, t)
}

/*
* Tests scan/unscan
 */
func TestParser2(t *testing.T) {
	parser := NewParser(strings.NewReader("program hello;"))
	var token lexer.Token = parser.scan()
	VerifyToken(lexer.PROGRAM, token, t)
	parser.unscan()
	token = parser.scan()
	VerifyToken(lexer.PROGRAM, token, t)
	token = parser.scan()
	VerifyTokenWithValue(lexer.IDENTIFIER, "hello", token, t)
	parser.unscan()
	token = parser.scan()
	VerifyTokenWithValue(lexer.IDENTIFIER, "hello", token, t)
	token = parser.scan()
	VerifyToken(lexer.SEMI_COLON, token, t)
	token = parser.scan()
	VerifyToken(lexer.EOF, token, t)
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
