package calc

type (
	Statement interface {
		statement()
	}

	Expression interface {
		expression()
	}
)

type (
	ExpressionStatement struct {
		Expr Expression
	}

	VarDefStatement struct {
		VarName string
		Expr    Expression
	}
)

func (x *ExpressionStatement) statement() {}
func (x *VarDefStatement) statement()     {}

type (
	NumberExpression struct {
		Lit string
	}

	IdentifierExpression struct {
		Lit string
	}

	UnaryMinusExpression struct {
		SubExpr Expression
	}

	ParenExpression struct {
		SubExpr Expression
	}

	BinOpExpression struct {
		LHS      Expression
		Operator int
		RHS      Expression
	}
)

func (x *NumberExpression) expression()     {}
func (x *IdentifierExpression) expression() {}
func (x *UnaryMinusExpression) expression() {}
func (x *ParenExpression) expression()      {}
func (x *BinOpExpression) expression()      {}
