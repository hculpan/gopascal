package lexer

import (
	"strings"
	"testing"
)

func VerifyToken(expectedType TokenType, token Token, t *testing.T) {
	if token.Type == ILLEGAL {
		t.Errorf("Illegal token '%s' [%d:%d]", token.Type, token.Pos.Line, token.Pos.Column)
	} else if expectedType != token.Type {
		t.Errorf("Expected %s token, got '%s' [%d:%d]", expectedType, token.Value, token.Pos.Line, token.Pos.Column)
	}
}

func VerifyTokenWithValue(expectedType TokenType, expectedValue string, token Token, t *testing.T) {
	if token.Type == ILLEGAL {
		t.Errorf("Illegal token '%s' [%d:%d]", token.Type, token.Pos.Line, token.Pos.Column)
	} else if expectedType != token.Type {
		t.Errorf("Expected %s token, got '%s' [%d:%d]", expectedType, token.Value, token.Pos.Line, token.Pos.Column)
	} else if expectedValue != token.Value {
		t.Errorf("Expected '%s', got '%s' [%d:%d]", expectedValue, token.Value, token.Pos.Line, token.Pos.Column)
	}
}

/*
* Tests basic characters ;(). as well as spaces
 */
func TestLex1(t *testing.T) {
	lexer := NewLexer(strings.NewReader("; ()   ."))
	tokens := lexer.Lex()
	if l := len(tokens); l != 4 {
		t.Error("Expected 4 tokens, got", len(tokens))
		return
	}
	VerifyToken(SEMI_COLON, tokens[0], t)
	VerifyToken(LEFT_PAREN, tokens[1], t)
	VerifyToken(RIGHT_PAREN, tokens[2], t)
	VerifyToken(PERIOD, tokens[3], t)
}

/*
* Tests program keyword, identifier
 */
func TestLex2(t *testing.T) {
	lexer := NewLexer(strings.NewReader("program hello    ;"))
	tokens := lexer.Lex()
	if l := len(tokens); l != 3 {
		t.Error("Expected 3 tokens, got", len(tokens))
		return
	}
	VerifyToken(PROGRAM, tokens[0], t)
	VerifyToken(IDENTIFIER, tokens[1], t)
	VerifyToken(SEMI_COLON, tokens[2], t)
}

/*
* Tests identifier with letters, digits, and underscore
 */
func TestLex3(t *testing.T) {
	lexer := NewLexer(strings.NewReader("program h12_ell_o    ;"))
	tokens := lexer.Lex()
	if l := len(tokens); l != 3 {
		t.Error("Expected 3 tokens, got", len(tokens))
		return
	}
	VerifyToken(PROGRAM, tokens[0], t)
	VerifyTokenWithValue(IDENTIFIER, "h12_ell_o", tokens[1], t)
	VerifyToken(SEMI_COLON, tokens[2], t)
}

/*
* Tests being/end keywords, multi-line, case-insensitivity
 */
func TestLex4(t *testing.T) {
	lexer := NewLexer(strings.NewReader(
		`Begin
		END.`))
	tokens := lexer.Lex()
	if l := len(tokens); l != 3 {
		t.Error("Expected 3 tokens, got", len(tokens))
		return
	}
	VerifyToken(BEGIN, tokens[0], t)
	VerifyToken(END, tokens[1], t)
	VerifyToken(PERIOD, tokens[2], t)
}

/*
* Tests being/end keywords, multi-line, case-insensitivity
 */
func TestLex5(t *testing.T) {
	lexer := NewLexer(strings.NewReader(
		`1 847 hell23              
		66`))
	tokens := lexer.Lex()
	if l := len(tokens); l != 4 {
		t.Error("Expected 3 tokens, got", len(tokens))
		return
	}
	VerifyTokenWithValue(INTEGER, "1", tokens[0], t)
	VerifyTokenWithValue(INTEGER, "847", tokens[1], t)
	VerifyTokenWithValue(IDENTIFIER, "hell23", tokens[2], t)
	VerifyTokenWithValue(INTEGER, "66", tokens[3], t)
}

/*
* Tests built-in functions and strings
 */
func TestLex6(t *testing.T) {
	lexer := NewLexer(strings.NewReader(
		`writeln();`))
	tokens := lexer.Lex()
	if l := len(tokens); l != 4 {
		t.Error("Expected 4 tokens, got", len(tokens))
		return
	}
	VerifyTokenWithValue(BUILT_IN, "writeln", tokens[0], t)
	VerifyToken(LEFT_PAREN, tokens[1], t)
	VerifyToken(RIGHT_PAREN, tokens[2], t)
	VerifyToken(SEMI_COLON, tokens[3], t)
}

/*
* Tests built-in functions and strings
 */
func TestLex7(t *testing.T) {
	lexer := NewLexer(strings.NewReader(
		`writeln('Hello, world!');`))
	tokens := lexer.Lex()
	if l := len(tokens); l != 5 {
		t.Error("Expected 5 tokens, got", len(tokens))
		return
	}
	VerifyTokenWithValue(BUILT_IN, "writeln", tokens[0], t)
	VerifyToken(LEFT_PAREN, tokens[1], t)
	VerifyTokenWithValue(STRING, "Hello, world!", tokens[2], t)
	VerifyToken(RIGHT_PAREN, tokens[3], t)
	VerifyToken(SEMI_COLON, tokens[4], t)
}

/*
* Tests whole Hello program
 */
func TestLex8(t *testing.T) {
	lexer := NewLexer(strings.NewReader(
		`program Hello;

		begin
			writeln('Hello, world!');
		end.`))
	tokens := lexer.Lex()
	if l := len(tokens); l != 11 {
		t.Error("Expected 11 tokens, got", len(tokens))
		return
	}
	VerifyToken(PROGRAM, tokens[0], t)
	VerifyTokenWithValue(IDENTIFIER, "Hello", tokens[1], t)
	VerifyToken(SEMI_COLON, tokens[2], t)
	VerifyToken(BEGIN, tokens[3], t)
	VerifyTokenWithValue(BUILT_IN, "writeln", tokens[4], t)
	VerifyToken(LEFT_PAREN, tokens[5], t)
	VerifyTokenWithValue(STRING, "Hello, world!", tokens[6], t)
	VerifyToken(RIGHT_PAREN, tokens[7], t)
	VerifyToken(SEMI_COLON, tokens[8], t)
	VerifyToken(END, tokens[9], t)
	VerifyToken(PERIOD, tokens[10], t)
}
