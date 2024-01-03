package lexer

import (
	"chimp/token"

	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(int x, int y) {
    x + y;
};
/*

*/
let result = add(five, ten);
`
	expTokens := []struct {
		expType    token.TokenType
		expLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.INT_KW, "int"},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.INT_KW, "int"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, "<eof>"},
	}

	l := New(input, "test.chp")

	for index, testTok := range expTokens {
		var tok token.Token = l.NextToken()

		if tok.Type != testTok.expType {
			t.Fatalf("tests[%d] - wrong tokentype. expected=%q, got=%q",
				index, testTok.expType, tok.Type)
		}

		if tok.Literal != testTok.expLiteral {
			t.Fatalf("tests[%d] - wrong tokenliteral. expected=%q, got=%q",
				index, testTok.expLiteral, tok.Literal)
		}
	}
}
