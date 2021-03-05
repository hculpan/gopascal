package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Lexer struct {
	pos    Position
	reader *bufio.Reader
	errors []string
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{Line: 1, Column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() ([]Token, []string) {
	result := []Token{}
	l.errors = []string{}

	token := l.NextToken()
	for token.Type != EOF {
		result = append(result, token)
		token = l.NextToken()
	}

	return result, l.errors
}

func (l *Lexer) NextToken() Token {
	// keep looping until we return a token
	for {
		l.pos.Column++

		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return NewToken(EOF, "", l.pos)
			}

			// at this point there isn't much we can do, and the compiler
			// should just return the raw error to the user
			panic(err)
		}

		switch r {
		case '\n':
			l.nextLine()
		case ';':
			return NewToken(SEMI_COLON, ";", l.pos)
		case '(':
			return NewToken(LEFT_PAREN, "(", l.pos)
		case ')':
			return NewToken(RIGHT_PAREN, ")", l.pos)
		case '.':
			return NewToken(PERIOD, ".", l.pos)
		case '\'':
			startPos := l.pos
			lit := l.lexString()
			return NewToken(STRING, lit, startPos)
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsLetter(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexIdentifier()
				return l.createToken(lit, startPos)
			} else if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexInteger()
				return NewToken(INTEGER, lit, startPos)
			} else {
				m := fmt.Sprintf("Invalid lexeme: '%s' at %d:%d", string(r), l.pos.Line, l.pos.Column)
				l.errors = append(l.errors, m)
				return NewToken(ILLEGAL, string(r), l.pos)
			}
		}
	}
}

func (l *Lexer) createToken(lit string, position Position) Token {
	switch strings.ToLower(lit) {
	case "program":
		return NewToken(PROGRAM, lit, position)
	case "begin":
		return NewToken(BEGIN, lit, position)
	case "end":
		return NewToken(END, lit, position)
	default:
		return NewToken(IDENTIFIER, lit, position)
	}
}

func (l *Lexer) lexString() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the int
				return lit
			}
		}

		l.pos.Column++
		if r != '\'' {
			lit = lit + string(r)
		} else {
			// scanned something not in the identifier
			return lit
		}
	}
}

func (l *Lexer) lexIdentifier() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the int
				return lit
			}
		}

		l.pos.Column++
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			lit = lit + string(r)
		} else {
			// scanned something not in the identifier
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexInteger() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the int
				return lit
			}
		}

		l.pos.Column++
		if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else {
			// scanned something not in the integer
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.Column--
}

func (l *Lexer) nextLine() {
	l.pos.Line++
	l.pos.Column = 0
}
