
//line parser.go.y:3
package calc
import __yyfmt__ "fmt"
//line parser.go.y:3
		
import (
	"log"
)

type Token struct {
	tok int
	lit string
	pos Position
}


//line parser.go.y:17
type yySymType struct{
	yys int
	statements []Statement
	statement  Statement
	expr       Expression
	tok        Token
}

const IDENT = 57346
const NUMBER = 57347
const VAR = 57348
const UNARY = 57349

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	"VAR",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UNARY",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:89


type LexerWrapper struct {
	s          *Scanner
	recentLit  string
	recentPos  Position
	statements []Statement
}

func (l *LexerWrapper) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.Scan()
	if tok == EOF {
		return 0
	}
	lval.tok = Token{tok: tok, lit: lit, pos: pos}
	l.recentLit = lit
	l.recentPos = pos
	return tok
}

func (l *LexerWrapper) Error(e string) {
	log.Fatalf("Line %d, Column %d: %q %s",
		l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func Parse(s *Scanner) []Statement {
	l := LexerWrapper{s: s}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.statements
}

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 14
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 50

var yyAct = []int{

	3, 11, 12, 13, 14, 15, 24, 16, 17, 18,
	25, 2, 19, 20, 21, 22, 23, 6, 5, 4,
	0, 7, 0, 6, 5, 26, 0, 7, 8, 11,
	12, 13, 14, 15, 8, 27, 11, 12, 13, 14,
	15, 0, 10, 13, 14, 15, 1, 0, 0, 9,
}
var yyPact = []int{

	13, -1000, 13, 29, 3, -1000, -1000, 19, 19, -1000,
	-1000, 19, 19, 19, 19, 19, -8, -1000, -6, 34,
	34, -1000, -1000, -1000, 19, -1000, 22, -1000,
}
var yyPgo = []int{

	0, 46, 11, 0,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3,
}
var yyR2 = []int{

	0, 0, 2, 2, 5, 1, 1, 2, 3, 3,
	3, 3, 3, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, 6, 5, 4, 8, 15, -1,
	13, 7, 8, 9, 10, 11, 4, -3, -3, -3,
	-3, -3, -3, -3, 14, 16, -3, 13,
}
var yyDef = []int{

	1, -2, 1, 0, 0, 5, 6, 0, 0, 2,
	3, 0, 0, 0, 0, 0, 0, 7, 0, 9,
	10, 11, 12, 13, 0, 8, 0, 4,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 11, 3, 3,
	15, 16, 9, 7, 3, 8, 3, 10, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 13,
	3, 14,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 12,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:38
		{
			yyVAL.statements = nil
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.statements = yyVAL.statements
			}
		}
	case 2:
		//line parser.go.y:45
		{
			yyVAL.statements = append([]Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.statements = yyVAL.statements
			}
		}
	case 3:
		//line parser.go.y:54
		{
			yyVAL.statement = &ExpressionStatement{Expr: yyS[yypt-1].expr}
		}
	case 4:
		//line parser.go.y:58
		{
			yyVAL.statement = &VarDefStatement{VarName: yyS[yypt-3].tok.lit, Expr: yyS[yypt-1].expr}
		}
	case 5:
		//line parser.go.y:63
		{
			yyVAL.expr = &NumberExpression{Lit: yyS[yypt-0].tok.lit}
		}
	case 6:
		//line parser.go.y:67
		{
			yyVAL.expr = &IdentifierExpression{Lit: yyS[yypt-0].tok.lit}
		}
	case 7:
		//line parser.go.y:71
		{
			yyVAL.expr = &UnaryMinusExpression{SubExpr: yyS[yypt-0].expr}
		}
	case 8:
		//line parser.go.y:75
		{
			yyVAL.expr = &ParenExpression{SubExpr: yyS[yypt-1].expr}
		}
	case 9:
		//line parser.go.y:79
		{ yyVAL.expr = &BinOpExpression{LHS: yyS[yypt-2].expr, Operator: int('+'), RHS: yyS[yypt-0].expr} }
	case 10:
		//line parser.go.y:81
		{ yyVAL.expr = &BinOpExpression{LHS: yyS[yypt-2].expr, Operator: int('-'), RHS: yyS[yypt-0].expr} }
	case 11:
		//line parser.go.y:83
		{ yyVAL.expr = &BinOpExpression{LHS: yyS[yypt-2].expr, Operator: int('*'), RHS: yyS[yypt-0].expr} }
	case 12:
		//line parser.go.y:85
		{ yyVAL.expr = &BinOpExpression{LHS: yyS[yypt-2].expr, Operator: int('/'), RHS: yyS[yypt-0].expr} }
	case 13:
		//line parser.go.y:87
		{ yyVAL.expr = &BinOpExpression{LHS: yyS[yypt-2].expr, Operator: int('%'), RHS: yyS[yypt-0].expr} }
	}
	goto yystack /* stack new state and value */
}
