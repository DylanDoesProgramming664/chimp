package lexer

import (
	"chimp/token"
	"fmt"
)

type Lexer struct {
	input    []rune
	Filename string
	pos      int
	Line     int
	char     rune
	Errors   []string
}

func New(input string, filename string) *Lexer {
	l := &Lexer{input: []rune(input), Filename: filename, Line: 1}
	l.char = l.input[0]

	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipSpace()

	if isDigit(l.char) {
		return newToken(token.INT, l.readNum())
	}

	if isAlpha(l.char) {
		return l.readIdent()
	}

	//two character statements
	twoCharStr := string(l.char) + string(l.nextChar())
	switch twoCharStr {
	case "<=":
		l.readChar()
		l.readChar()
		return newToken(token.LTEQ, twoCharStr)
	case "==":
		l.readChar()
		l.readChar()
		return newToken(token.EQ, twoCharStr)
	case ">=":
		l.readChar()
		l.readChar()
		return newToken(token.GTEQ, twoCharStr)
	case "!=":
		l.readChar()
		l.readChar()
		return newToken(token.NOTEQ, twoCharStr)
	case "&&":
		l.readChar()
		l.readChar()
		return newToken(token.BOOLAND, twoCharStr)
	case "||":
		l.readChar()
		l.readChar()
		return newToken(token.BOOLOR, twoCharStr)
	case "^^":
		l.readChar()
		l.readChar()
		return newToken(token.BOOLXOR, twoCharStr)
	case ">>":
		l.readChar()
		l.readChar()
		return newToken(token.RBITSHIFT, twoCharStr)
	case "<<":
		l.readChar()
		l.readChar()
		return newToken(token.LBITSHIFT, twoCharStr)
	case "??":
		l.readChar()
		l.readChar()
		return newToken(token.COALESCE, twoCharStr)
	case "//":
		l.skipLineComment()
		return l.NextToken()
	case "/*":
		l.skipBlockComment()
		return l.NextToken()

	}

	var tok token.Token

	charStr := string(l.char)
	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, charStr)
	case '+':
		tok = newToken(token.PLUS, charStr)
	case '-':
		tok = newToken(token.MINUS, charStr)
	case '*':
		tok = newToken(token.STAR, charStr)
	case '/':
		tok = newToken(token.SLASH, charStr)
	case ',':
		tok = newToken(token.COMMA, charStr)
	case '(':
		tok = newToken(token.LPAREN, charStr)
	case ')':
		tok = newToken(token.RPAREN, charStr)
	case '{':
		tok = newToken(token.LBRACE, charStr)
	case '}':
		tok = newToken(token.RBRACE, charStr)
	case ';':
		tok = newToken(token.SEMICOLON, charStr)
	case '<':
		tok = newToken(token.LT, charStr)
	case '>':
		tok = newToken(token.GT, charStr)
	case '&':
		tok = newToken(token.BITAND, charStr)
	case '|':
		tok = newToken(token.BITOR, charStr)
	case '^':
		tok = newToken(token.BITXOR, charStr)
	case '~':
		tok = newToken(token.BITNOT, charStr)
	case '!':
		tok = newToken(token.BANG, charStr)
	case '#':
		tok = newToken(token.HASH, charStr)
	case '?':
		tok = newToken(token.QUESTION, charStr)
	case 0:
		tok = newToken(token.EOF, "<eof>")
	default:
		tok = newToken(token.ILLEGAL, charStr)
	}

	l.readChar()

	return tok
}

func (l *Lexer) skipSpace() {
	for isSpace(l.char) {
		if l.char == '\n' {
			l.Line++
		}
		l.readChar()
	}
}

func isSpace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\r' || char == '\n'
}

func (l *Lexer) readChar() {
	l.char = l.nextChar()
	l.pos++
}

func (l *Lexer) nextChar() rune {
	return l.nextNthChar(1)
}

func (l *Lexer) nextNthChar(n int) rune {
	if l.pos+n >= len(l.input) {
		return 0
	}

	return l.input[l.pos+n]
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func (l *Lexer) readNum() string {
	numStr := ""
	for isDigit(l.char) {
		numStr += string(l.char)
		l.readChar()
	}
	return numStr
}

func isAlpha(char rune) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}

func (l *Lexer) readIdent() token.Token {
	ident := string(l.char)
	l.readChar()
	for isAlnum(l.char) || l.char == '_' {
		ident += string(l.char)
		l.readChar()
	}
	tokType := token.MatchIdent(ident)
	return newToken(tokType, ident)
}

func isAlnum(char rune) bool {
	return isDigit(char) || isAlpha(char)
}

func (l *Lexer) skipLineComment() {
	for l.char != '\n' && l.char != 0 {
		l.readChar()
	}
}

func (l *Lexer) skipBlockComment() {
	start_Line := l.Line
	l.readChar()
	l.readChar()
	for {
		if l.char == '*' && l.nextChar() == '/' {
			l.readChar()
			l.readChar()
			break
		}
		if l.char == '/' && l.nextChar() == '*' {
			l.skipBlockComment()
		}
		if l.char == 0 {
			l.Errors = append(l.Errors, fmt.Sprintf("Syntax Error:%d: Unterminated block bomment before end of file.\n", start_Line))
			break
		}
		l.readChar()
	}
}

func newToken(Type token.TokenType, Literal string) token.Token {
	return token.Token{Type: Type, Literal: Literal}
}
