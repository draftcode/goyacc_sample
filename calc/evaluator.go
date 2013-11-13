package calc

import (
	"fmt"
	"strconv"
)

type Env map[string]int

func Evaluate(statement Statement, env Env) (string, error) {
	switch stmt := statement.(type) {
	case *ExpressionStatement:
		v, err := evaluateExpr(stmt.Expr, env)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(v), nil
	case *VarDefStatement:
		v, err := evaluateExpr(stmt.Expr, env)
		if err != nil {
			return "", err
		}
		env[stmt.VarName] = v
		return fmt.Sprintf("Assign %v to %s", v, stmt.VarName), nil
	default:
		panic("Unknown Statement type")
	}
}

func evaluateExpr(expr Expression, env Env) (int, error) {
	switch e := expr.(type) {
	case *NumberExpression:
		v, err := strconv.Atoi(e.Lit)
		if err != nil {
			return 0, err
		}
		return v, nil
	case *IdentifierExpression:
		if v, ok := env[e.Lit]; ok {
			return v, nil
		} else {
			return 0, fmt.Errorf("Undefined variable: %s", e.Lit)
		}
	case *UnaryMinusExpression:
		v, err := evaluateExpr(e.SubExpr, env)
		if err != nil {
			return 0, err
		}
		return -v, nil
	case *ParenExpression:
		v, err := evaluateExpr(e.SubExpr, env)
		if err != nil {
			return 0, err
		}
		return v, nil
	case *BinOpExpression:
		lhsV, err := evaluateExpr(e.LHS, env)
		if err != nil {
			return 0, err
		}
		rhsV, err := evaluateExpr(e.RHS, env)
		if err != nil {
			return 0, err
		}
		switch e.Operator {
		case '+':
			return lhsV + rhsV, nil
		case '-':
			return lhsV - rhsV, nil
		case '*':
			return lhsV * rhsV, nil
		case '/':
			return lhsV / rhsV, nil
		case '%':
			return lhsV % rhsV, nil
		default:
			panic("Unknown operator")
		}
	default:
		panic("Unknown Expression type")
	}
}
