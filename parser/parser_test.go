package parser

import (
	"fmt"
	"github.com/piglang/ast"
	"github.com/piglang/ast/expressions"
	"github.com/piglang/ast/literals"
	"github.com/piglang/ast/statements"
	"github.com/piglang/lexer"
	"github.com/piglang/token"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `let x = 5;
	let y = 10;
	let foo = 80321;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements doesn't have 3 statements, theinput was not parsed properly. Got=%d", len(program.Statements))

	}

	tests := []struct {
		expectedIdent string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdent) {
			return
		}
	}
}

func TestParserErrorCreation(t *testing.T) {
	input := `let x = 5;
	let y = 10;
	let foo = 80321;
	`
	l := lexer.New(input)
	p := New(l)
	p.ParseProgram()
	checkParserErrors(t, p)

}

// Check if any errors were registered during the parsing.
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors.", len(errors))
	for _, msg := range errors {
		t.Errorf("Parser Error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
	}

	letStmt, ok := s.(*statements.LetStatement)
	if !ok {
		t.Errorf("s not a *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s. got=%s", name,
			letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 9398321;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("Program does not have 3 statements. got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*statements.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. Got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return'. Got %q", returnStmt.TokenLiteral())
		}
	}

}

func TestIdentifierExpression(t *testing.T) {
	input := "testIdent;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program doesn't have correct number of statements. got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*statements.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statement[0] is not of the type ExpressionStatement. got %T", stmt)
	}

	ident, ok := stmt.Expression.(*statements.Identifier)
	if !ok {
		t.Fatalf("stmt is not of the type Identifier. got %T", ident)
	}

	if ident.Token.Type != token.IDENT {
		t.Fatalf("token type is incorrect, expected %s got %s", token.IDENT, ident.Token.Type)
	}
	if ident.Value != "testIdent" {
		t.Fatalf("token value is incorrect, expecter %s got %s", input, ident.Value)
	}

}

func TestIntegerLiteral(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("program doesn't have correct number of statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*statements.ExpressionStatement)
	if !ok {
		t.Fatalf("Statement[0] is not of the type ExpressionStatement. Got %T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*literals.IntegerLiteral)
	if !ok {
		t.Errorf("Expression is not of the type IntegerLiteral. got %T", stmt.Expression)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("Value of the literal is not %s. got %s", "5", literal.TokenLiteral())
	}

}

// Testing parsing for prefix expressions like !5, -2, etc.
func TestParsingPrefixExpression(t *testing.T){

	prefix := []struct{
		input string
		prefixOperator string
		operand interface{}
	}{
		{"-5", "-", 5},
		{"!6", "!", 6},
		{"!false", "!", false},
		{"!true", "!", true},
	}

	for _, tc := range prefix {
		l := lexer.New(tc.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1{
			t.Fatalf("The number of statements is not correct. Expected 1, got %d",len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*statements.ExpressionStatement)
		if !ok {
			t.Fatalf("statement is not of the type ExpressionStatement. Got %T", program.Statements[0])
		}


		expression, ok := stmt.Expression.(*expressions.PrefixExpression)
		if !ok {
			t.Fatalf("Statement is not of the type PrefixExpression. Got %T", stmt)
		}

		if expression.Operator != tc.prefixOperator {
			t.Fatalf("The Operator is not correct. Expected %s got %s", tc.prefixOperator, expression.Operator)
		}

		if !testIsProperExpressionOperand(t, expression.RightOperand, tc.operand){
			return
		}

	}

}

// Helper function to explicitly check if an Expression is an Integer Literal
func testIsIntegerLiteral(t *testing.T, il ast.Expression, val int64) bool{
	intLiteral, ok := il.(*literals.IntegerLiteral)
	if !ok {
		t.Errorf("Expression is not IntegerLiteral type. Got %T", il)
		return false
	}

	if intLiteral.Value != val {
		t.Errorf("The value for the integer is not correct. Expecter %d got %d", val, intLiteral.Value)
		return false
	}

	if intLiteral.TokenLiteral() != fmt.Sprint(val){
		t.Errorf("Integer literal not %d. Got %s", val, intLiteral.TokenLiteral())
		return false
	}
	return true
}

//Generic helper to validate that the expression operand are of the valid type and have the valid value.
func testIsProperExpressionOperand(t *testing.T, exp ast.Expression,val interface{}) bool {
	switch v := val.(type) {
	case int:
		return testIsIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIsIntegerLiteral(t, exp, v)
	case string:
		return testIsIdentifier(t, exp, v)
	case bool:
		return testIsBooleanLiteral(t, exp, v)
	}
	t.Errorf("Type of Expression not handled yet. Got %T", exp)
	return false
}

//Helper function to test if an expression is an identifier
func testIsIdentifier(t *testing.T, exp ast.Expression, value string) bool{
	ident, ok := exp.(*statements.Identifier)
	if !ok {
		t.Errorf("Expression is not of the type Identifier. Got %T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("Incorrect value for Identifier. Exprected %s got %s", value,ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("Incorrect token literal. Expected %s got %s", value, ident.TokenLiteral())
		return false
	}

	return true
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*expressions.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.OperatorExpression. got=%T(%s)", exp, exp)
		return false }
	if !testIsProperExpressionOperand(t, opExp.LeftOperand, left) {
		return false
	}
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}
	if !testIsProperExpressionOperand(t, opExp.RightOperand, right) {
		return false
	}
	return true
}

// Extended generic testing of infix expressions.
func TestParsingInfixExpression(t *testing.T){

	testCases := []struct {
		input string
		LeftOperand interface{}
		Operator string
		RightOperand interface{}
	}{
		{"6+5",6,"+",5},
		{"6-5",6,"-",5},
		{"6/5",6,"/",5},
		{"6*5",6,"*",5},
		{"6>5",6,">",5},
		{"6<5",6,"<",5},
		{"6==5",6,"==",5},
		{"6!=5",6,"!=",5},
		{"true == true", true,"==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, tc := range testCases{
		l := lexer.New(tc.input)
		parser := New(l)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("Number of statements incorrect. Expected 1 got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*statements.ExpressionStatement)
		if !ok {
			t.Fatalf("Statement[0] is not of the type ExpressionStatement. Got %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*expressions.InfixExpression)
		if !ok {
			t.Fatalf("Statement expression is not of the type InfixExpression. Got %T", stmt.Expression)
		}

		if !testIsProperExpressionOperand(t,exp.LeftOperand, tc.LeftOperand ){
			return
		}

		if exp.Operator != tc.Operator {
			t.Fatalf("Incorrect operator. Exptected %s got %s", tc.Operator, exp.Operator)
		}

		if !testIsProperExpressionOperand(t, exp.RightOperand, tc.RightOperand){
			return
		}
	}

}

// Generic testing of precedence.
func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
				"true",
				"true",
		},
		{
			"false",
			"false",
		},
		{
			"a > b == true",
			"((a > b) == true)",
		},
		{
				"a < b != false",
				"((a < b) != false)",
		},
		{
			"a + (b + c) + d",
			"((a + (b + c)) + d)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

// Test parsing boolean expressions explicitly.
func TestParsingBooleanExpression(t *testing.T){
	input := []struct {
		input string
		expected bool
	}{
		{input: "true;", expected: true},
		{input: "false;", expected:false},

	}

	for _, tc := range input {
		l := lexer.New(tc.input)
		parser := New(l)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1{
			t.Fatalf("Number of statements incorrect. Expected %d. Got %d",1, len(program.Statements))
		}

		stmt, ok  := program.Statements[0].(*statements.ExpressionStatement)
		if !ok {
			t.Fatalf("Statement is not of the type ExpressionStatement. Got %T", program.Statements[0])
		}

		boolLiteral, ok  := stmt.Expression.(*literals.BooleanLiteral)

		if !ok {
			t.Fatalf("Expression is not of the type Boolean Literal. Got %T", stmt.Expression)
		}

		if boolLiteral.Value != tc.expected {
			t.Fatalf("Incorrect Boolean Literal Value. Expected %v got %v", tc.expected, boolLiteral.Value)
		}

	}

}

func testIsBooleanLiteral(t *testing.T,bl ast.Expression, val bool ) bool {

	literal, ok := bl.(*literals.BooleanLiteral)
	if !ok {
		t.Errorf("Expression is not of the type bool. Got %T", bl)
		return false
	}

	if literal.Value != val {
		t.Errorf("Incorrect boolean literal value. Expected %v got %v", val, literal.Value)
		return false
	}

	if literal.TokenLiteral() != fmt.Sprintf("%t", val){
		t.Errorf("Incorrect boolean token literal. Exptected %t, got %s", val, literal.TokenLiteral())
		return false
	}

	return true
}