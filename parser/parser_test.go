package parser

import (
	"chimp/ast"
	"chimp/lexer"
	"testing"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input, "lettest")
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil\n")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements. got=%d\n", len(program.Statements))
	}

	tests := []struct {
		expIdent string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for index, testStruct := range tests {
		stmt := program.Statements[index]
		if !testLetStatement(t, stmt, testStruct.expIdent, p) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string, p *Parser) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("%s:%d: s.TokenLiteral not 'let', got=%q\n",
			p.l.Filename, p.l.Line, s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("%s:%d: s not *ast.LetStatement, got=%T\n",
			p.l.Filename, p.l.Line, s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("%s:%d: letStmt.Name.Value not '%s'. got=%s\n",
			p.l.Filename, p.l.Line, name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("%s:%d: letStmt.Name.TokenLiteral not '%s'. got=%s\n",
			p.l.Filename, p.l.Line, name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestIntStatements(t *testing.T) {
	input := `
int x = 5;
int y = 10;
int foobar = 838383;
`

	l := lexer.New(input, "inttest")
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil\n")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements. got=%d\n", len(program.Statements))
	}

	tests := []struct {
		expIdent string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for index, testStruct := range tests {
		stmt := program.Statements[index]
		if !testIntStatement(t, stmt, testStruct.expIdent, p) {
			return
		}
	}
}

func testIntStatement(t *testing.T, s ast.Statement, name string, p *Parser) bool {
	if s.TokenLiteral() != "int" {
		t.Errorf("%s:%d: s.TokenLiteral not 'int', got=%q\n",
			p.l.Filename, p.l.Line, s.TokenLiteral())
		return false
	}

	intStmt, ok := s.(*ast.IntStatement)
	if !ok {
		t.Errorf("%s:%d: s not *ast.IntStatement, got=%T\n",
			p.l.Filename, p.l.Line, s)
		return false
	}

	if intStmt.Name.Value != name {
		t.Errorf("%s:%d: intStmt.Name.Value not '%s'. got=%s\n",
			p.l.Filename, p.l.Line, name, intStmt.Name.Value)
		return false
	}

	if intStmt.Name.TokenLiteral() != name {
		t.Errorf("%s:%d: intStmt.Name.TokenLiteral not '%s'. got=%s\n",
			p.l.Filename, p.l.Line, name, intStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestBoolStatements(t *testing.T) {
	input := `
bool x = true;
bool y = false;
`

	l := lexer.New(input, "booltest")
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil\n")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("Expected 2 statements. got=%d\n", len(program.Statements))
	}

	tests := []struct {
		expIdent string
	}{
		{"x"},
		{"y"},
	}

	for index, testStruct := range tests {
		stmt := program.Statements[index]
		if !testBoolStatement(t, stmt, testStruct.expIdent, p) {
			return
		}
	}
}

func testBoolStatement(t *testing.T, s ast.Statement, name string, p *Parser) bool {
	if s.TokenLiteral() != "bool" {
		t.Errorf("%s:%d: s.TokenLiteral: expected='bool', got=%q\n",
			p.l.Filename, p.l.Line, s.TokenLiteral())
		return false
	}

	boolStmt, ok := s.(*ast.BoolStatement)
	if !ok {
		t.Errorf("%s:%d: s: expected=*ast.BoolStatement, got=%T\n",
			p.l.Filename, p.l.Line, s)
		return false
	}

	if boolStmt.Name.Value != name {
		t.Errorf("%s:%d: boolStmt.Name.Value: expected='%s'. got=%s\n",
			p.l.Filename, p.l.Line, name, boolStmt.Name.Value)
		return false
	}

	if boolStmt.Name.TokenLiteral() != name {
		t.Errorf("%s:%d: boolStmt.Name.TokenLiteral: expected=%s. got=%s\n",
			p.l.Filename, p.l.Line, name, boolStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return true;
`

	l := lexer.New(input, "returntest")
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("Expected 2 statements. got=%d.\n",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("%s:%d: stmt: expected=*ast.ReturnStatement. got=%T.\n",
				p.l.Filename, p.l.Line, stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("%s:%d: returnStmt.TokenLiteral(): expected=return. got=%s",
				p.l.Filename, p.l.Line, returnStmt.TokenLiteral())
		}
	}
}

/*
func TestStringStatements(t *testing.T) {
	input := `
string x = "meow";
string y = "10";
string foobar = ":";
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil\n")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements. got=%d\n", len(program.Statements))
	}

	tests := []struct {
		expIdent string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for index, testStruct := range tests {
		stmt := program.Statements[index]
		if !testStringStatement(t, stmt, testStruct.expIdent, p) {
			return
		}
	}
}

func testStringStatement(t *testing.T, s ast.Statement, name string, p *Parser) bool {
	if s.TokenLiteral() != "string" {
        t.Errorf("%s:%d: s.TokenLiteral not 'string', got=%q\n",
            p.l.Filename, p.l.Line,s.TokenLiteral())
		return false
	}

	stringStmt, ok := s.(*ast.StringStatement)
	if !ok {
        t.Errorf("%s:%d: s not *ast.StringStatement, got=%T\n",
            p.l.Filename, p.l.Line,s)
		return false
	}

	if stringStmt.Name.Value != name {
        t.Errorf("%s:%d: stringStmt.Name.Value not '%s'. got=%s\n",
            p.l.Filename, p.l.Line,name, stringStmt.Name.Value)
		return false
	}

	if stringStmt.Name.TokenLiteral() != name {
        t.Errorf(%s:%d: "stringStmt.Name.TokenLiteral not '%s'. got=%s\n",
			p.l.Filename, p.l.Line,name, stringStmt.Name.TokenLiteral())
		return false
	}

	return true
}
*/
