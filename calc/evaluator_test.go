package calc

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	{
		s, _ := Evaluate(&ExpressionStatement{&NumberExpression{"123"}}, Env{})
		if s != "123" {
			t.Errorf("Expect a statement \"123;\" to be evaluated as \"123\", but got %q", s)
		}
	}
	{
		s, err := Evaluate(&ExpressionStatement{&IdentifierExpression{"a"}}, Env{})
		if err == nil {
			t.Errorf("Expect an undefined variable \"a\" not to be evaluated, but got %q", s)
		}
	}
	{
		env := Env{"a": 16}
		s, _ := Evaluate(&VarDefStatement{"a", &NumberExpression{"123"}}, env)
		if s != "Assign 123 to a" {
			t.Errorf("Expect a VarDef statement do an assignment, but it didn't")
		}
		if env["a"] != 123 {
			t.Errorf("Expect a VarDef statement do an assignment, but it didn't")
		}
	}
	{
		s, err := Evaluate(&VarDefStatement{"a", &IdentifierExpression{"a"}}, Env{})
		if err == nil {
			t.Errorf("Expect an undefined variable \"a\" not to be evaluated, but got %q", s)
		}
	}
}

func TestEvaluateExpr(t *testing.T) {
	if v, _ := evaluateExpr(&NumberExpression{"123"}, Env{}); v != 123 {
		t.Error("Expect a number 123 to be evaluated as 123, but got", v)
	}
	if v, _ := evaluateExpr(&IdentifierExpression{"a"}, Env{"a": 123}); v != 123 {
		t.Error("Expect a variable \"a\" to be evaluated as 123, but got", v)
	}
	if v, err := evaluateExpr(&IdentifierExpression{"a"}, Env{}); err == nil {
		t.Error("Expect an undefined variable \"a\" not to be evaluated, but got", v)
	}
	if v, _ := evaluateExpr(&UnaryMinusExpression{&NumberExpression{"123"}}, Env{}); v != -123 {
		t.Error("Expect a number -123 to be evaluated as -123, but got", v)
	}
	if v, err := evaluateExpr(&UnaryMinusExpression{&IdentifierExpression{"a"}}, Env{}); err == nil {
		t.Error("Expect an undefined variable \"a\" not to be evaluated, but got", v)
	}
	if v, _ := evaluateExpr(&ParenExpression{&NumberExpression{"123"}}, Env{}); v != 123 {
		t.Error("Expect an expression (123) to be evaluated as 123, but got", v)
	}
	if v, err := evaluateExpr(&ParenExpression{&IdentifierExpression{"a"}}, Env{}); err == nil {
		t.Error("Expect an undefined variable \"a\" not to be evaluated, but got", v)
	}

	expr1 := &NumberExpression{"42"}
	expr2 := &NumberExpression{"12"}
	exprE := &IdentifierExpression{"a"}
	if v, _ := evaluateExpr(&BinOpExpression{expr1, '+', expr2}, Env{}); v != 54 {
		t.Error("Expect an expression 42+12 to be evaluated as 54, but got", v)
	}
	if v, _ := evaluateExpr(&BinOpExpression{expr1, '-', expr2}, Env{}); v != 30 {
		t.Error("Expect an expression 42-12 to be evaluated as 30, but got", v)
	}
	if v, _ := evaluateExpr(&BinOpExpression{expr1, '*', expr2}, Env{}); v != 504 {
		t.Error("Expect an expression 42*12 to be evaluated as 3, but got", v)
	}
	if v, _ := evaluateExpr(&BinOpExpression{expr1, '/', expr2}, Env{}); v != 3 {
		t.Error("Expect an expression 42/12 to be evaluated as 3, but got", v)
	}
	if v, _ := evaluateExpr(&BinOpExpression{expr1, '%', expr2}, Env{}); v != 6 {
		t.Error("Expect an expression 42%12 to be evaluated as 6, but got", v)
	}
	if v, err := evaluateExpr(&BinOpExpression{exprE, '+', expr2}, Env{}); err == nil {
		t.Error("Expect an undefined variable \"a\" not to be evaluated, but got", v)
	}
	if v, err := evaluateExpr(&BinOpExpression{expr1, '+', exprE}, Env{}); err == nil {
		t.Error("Expect an undefined variable \"a\" not to be evaluated, but got", v)
	}
}
