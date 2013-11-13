// vim: noet sw=8 sts=8
%{
package calc

import (
	"log"
)

type Token struct {
	tok int
	lit string
	pos Position
}

%}

%union{
	statements []Statement
	statement  Statement
	expr       Expression
	tok        Token
}

%type<statements> statements
%type<statement> statement
%type<expr> expr

%token<tok> IDENT NUMBER VAR

%left '+' '-'
%left '*' '/' '%'
%right UNARY

%%

statements
	:
	{
		$$ = nil
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.statements = $$
		}
	}
	| statement statements
	{
		$$ = append([]Statement{$1}, $2...)
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.statements = $$
		}
	}

statement
	: expr ';'
	{
		$$ = &ExpressionStatement{Expr: $1}
	}
	| VAR IDENT '=' expr ';'
	{
		$$ = &VarDefStatement{VarName: $2.lit, Expr: $4}
	}

expr	: NUMBER
	{
		$$ = &NumberExpression{Lit: $1.lit}
	}
	| IDENT
	{
		$$ = &IdentifierExpression{Lit: $1.lit}
	}
	| '-' expr      %prec UNARY
	{
		$$ = &UnaryMinusExpression{SubExpr: $2}
	}
	| '(' expr ')'
	{
		$$ = &ParenExpression{SubExpr: $2}
	}
	| expr '+' expr
	{ $$ = &BinOpExpression{LHS: $1, Operator: int('+'), RHS: $3} }
	| expr '-' expr
	{ $$ = &BinOpExpression{LHS: $1, Operator: int('-'), RHS: $3} }
	| expr '*' expr
	{ $$ = &BinOpExpression{LHS: $1, Operator: int('*'), RHS: $3} }
	| expr '/' expr
	{ $$ = &BinOpExpression{LHS: $1, Operator: int('/'), RHS: $3} }
	| expr '%' expr
	{ $$ = &BinOpExpression{LHS: $1, Operator: int('%'), RHS: $3} }

%%

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
