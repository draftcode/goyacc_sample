package calc

import (
	"reflect"
	"testing"
)

func parseStmt(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init(src)
	statements := Parse(s)
	if len(statements) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if !reflect.DeepEqual(statements[0], expect) {
		t.Errorf("Expect %+#v to be %+#v", statements[0], expect)
		return
	}
}

func TestParseStatement(t *testing.T) {
	parseStmt(t, "123;", &ExpressionStatement{&NumberExpression{"123"}})
	parseStmt(t, "var a = 123;", &VarDefStatement{"a", &NumberExpression{"123"}})
}

func parseExpr(t *testing.T, src string, expect interface{}) {
	s := new(Scanner)
	s.Init(src + ";")
	statements := Parse(s)
	if len(statements) != 1 {
		t.Errorf("Expect %q to be parsed", src)
		return
	}
	if exprStmt, isExprStmt := statements[0].(*ExpressionStatement); isExprStmt {
		if !reflect.DeepEqual(exprStmt.Expr, expect) {
			t.Errorf("Expect %+#v to be %+#v", exprStmt.Expr, expect)
			return
		}
	}
}

func TestParseExpr(t *testing.T) {
	parseExpr(t, "123", &NumberExpression{"123"})
	parseExpr(t, "abc", &IdentifierExpression{"abc"})
	parseExpr(t, "-abc", &UnaryMinusExpression{&IdentifierExpression{"abc"}})
	parseExpr(t, "(abc)", &ParenExpression{&IdentifierExpression{"abc"}})

	aExp := &IdentifierExpression{"a"}
	bExp := &IdentifierExpression{"b"}
	parseExpr(t, "a+b", &BinOpExpression{aExp, '+', bExp})
	parseExpr(t, "a-b", &BinOpExpression{aExp, '-', bExp})
	parseExpr(t, "a*b", &BinOpExpression{aExp, '*', bExp})
	parseExpr(t, "a/b", &BinOpExpression{aExp, '/', bExp})
	parseExpr(t, "a%b", &BinOpExpression{aExp, '%', bExp})
}
