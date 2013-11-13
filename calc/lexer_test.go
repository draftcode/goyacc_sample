package calc

import (
	"testing"
)

func testScanner(t *testing.T, src string, expectTok int) {
	s := new(Scanner)
	s.Init(src)
	tok, lit, _ := s.Scan()
	if tok != expectTok {
		t.Errorf("Expect Scanner{%q}.Scan() = %#v, _ want %#v", src, tok, expectTok)
	}
	if lit != src {
		t.Errorf("Expect Scanner{%q}.Scan() = _, %#v want %#v", src, lit, src)
	}
	tok, lit, _ = s.Scan()
	if tok != EOF {
		t.Errorf("Expect Scanner{%q}.Scan() = %#v, _ want %#v", src, tok, EOF)
	}
}

func TestScanner(t *testing.T) {
	testScanner(t, "var", VAR)
	testScanner(t, "abc", IDENT)
	testScanner(t, "123", NUMBER)
	testScanner(t, "(", '(')
	testScanner(t, ")", ')')
	testScanner(t, ";", ';')
	testScanner(t, "+", '+')
	testScanner(t, "-", '-')
	testScanner(t, "*", '*')
	testScanner(t, "/", '/')
	testScanner(t, "%", '%')
	testScanner(t, "=", '=')
}
