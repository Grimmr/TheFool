package Parser

import (
	"testing"
)

// lexer only tests
func TestLexProgrammeKeywords(t *testing.T) {
	//should include one of every token
	programme := "and or a.txt less()%+ - filter"
	result := LexProgramme(programme)
	expected := []lexToken{lexToken{"and", LexTokenType_and}, lexToken{"or", LexTokenType_or}, lexToken{"a.txt", LexTokenType_name}, lexToken{"less", LexTokenType_less}, lexToken{"'('", LexTokenType_lParen}, lexToken{"')'", LexTokenType_rParen}, lexToken{"%", LexTokenType_percent}, lexToken{"+", LexTokenType_plus}, lexToken{"-", LexTokenType_minus}, lexToken{"filter", LexTokenType_filter}}

	t.Logf("Lex returned %v", result)

	if len(result) != len(expected) {
		t.Fatalf("Lex return wrong size, expected %d got %d", len(expected), len(result))
	}

	for index := range result {
		if result[index] != expected[index] {
			t.Errorf("Lex returned wrong token at index %d, expected %+v got %+v", index, expected[index], result[index])
		}
	}
}

func TestLexQoutesDontEffectNormalLex(t *testing.T) {
	programme := "\"a\" and \"b\""
	result := LexProgramme(programme)
	expected := []lexToken{lexToken{"a", LexTokenType_name}, lexToken{"and", LexTokenType_and}, lexToken{"b", LexTokenType_name}}

	t.Logf("Lex returned %v", result)

	if len(result) != len(expected) {
		t.Fatalf("Lex return wrong size, expected %d got %d", len(expected), len(result))
	}

	for index := range result {
		if result[index] != expected[index] {
			t.Errorf("Lex returned wrong token at index %d, expected %+v got %+v", index, expected[index], result[index])
		}
	}
}

func TestLexQoutesAllowSpaceInName(t *testing.T) {
	programme := "\"a b\" and c"
	result := LexProgramme(programme)
	expected := []lexToken{lexToken{"a b", LexTokenType_name}, lexToken{"and", LexTokenType_and}, lexToken{"c", LexTokenType_name}}

	t.Logf("Lex returned %v", result)

	if len(result) != len(expected) {
		t.Fatalf("Lex return wrong size, expected %d got %d", len(expected), len(result))
	}

	for index := range result {
		if result[index] != expected[index] {
			t.Errorf("Lex returned wrong token at index %d, expected %+v got %+v", index, expected[index], result[index])
		}
	}
}

func TestLexQoutesSpecialSymbolAtNameEnd(t *testing.T) {
	programme := "\"a b%\" and c"
	result := LexProgramme(programme)
	expected := []lexToken{lexToken{"a b%", LexTokenType_name}, lexToken{"and", LexTokenType_and}, lexToken{"c", LexTokenType_name}}

	t.Logf("Lex returned %v", result)

	if len(result) != len(expected) {
		t.Fatalf("Lex return wrong size, expected %d got %d", len(expected), len(result))
	}

	for index := range result {
		if result[index] != expected[index] {
			t.Errorf("Lex returned wrong token at index %d, expected %+v got %+v", index, expected[index], result[index])
		}
	}
}

func TestLexQoutesAllowKeywordAsName(t *testing.T) {
	programme := "\"and\" and c"
	result := LexProgramme(programme)
	expected := []lexToken{lexToken{"and", LexTokenType_name}, lexToken{"and", LexTokenType_and}, lexToken{"c", LexTokenType_name}}

	t.Logf("Lex returned %v", result)

	if len(result) != len(expected) {
		t.Fatalf("Lex return wrong size, expected %d got %d", len(expected), len(result))
	}

	for index := range result {
		if result[index] != expected[index] {
			t.Errorf("Lex returned wrong token at index %d, expected %+v got %+v", index, expected[index], result[index])
		}
	}
}

func TestLexQoutesMidBufferOk(t *testing.T) {
	programme := "\"a \"b and c"
	result := LexProgramme(programme)
	expected := []lexToken{lexToken{"a b", LexTokenType_name}, lexToken{"and", LexTokenType_and}, lexToken{"c", LexTokenType_name}}

	t.Logf("Lex returned %v", result)

	if len(result) != len(expected) {
		t.Fatalf("Lex return wrong size, expected %d got %d", len(expected), len(result))
	}

	for index := range result {
		if result[index] != expected[index] {
			t.Errorf("Lex returned wrong token at index %d, expected %+v got %+v", index, expected[index], result[index])
		}
	}
}

// WalkParseTree tests
func TestWalkParseTree(t *testing.T) {
	programme := &ParseTreeNode{lexToken{"and", LexTokenType_and}, []*ParseTreeNode{&ParseTreeNode{lexToken{"or", LexTokenType_or}, []*ParseTreeNode{&ParseTreeNode{lexToken{"a", LexTokenType_name}, nil}, &ParseTreeNode{lexToken{"b", LexTokenType_name}, nil}}}, &ParseTreeNode{lexToken{"c", LexTokenType_name}, nil}}}
	result := WalkParseTree(programme)
	expected := "(and (or (a) (b)) (c))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

// parser only tests ()
// a simple one level test of and
func TestParseOnly(t *testing.T) {
	programme := []lexToken{lexToken{"a", LexTokenType_name}, lexToken{"and", LexTokenType_and}, lexToken{"b", LexTokenType_name}}
	result := WalkParseTree(ParseProgramme(programme))
	expected := "(and (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

// full stack tests
// operator checks
func TestParseAnd(t *testing.T) {
	programme := "a and b"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(and (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseOr(t *testing.T) {
	programme := "a or b"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(or (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseAndOr(t *testing.T) {
	programme := "a and b or c"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(or (and (a) (b)) (c))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseLess(t *testing.T) {
	programme := "a less b"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(less (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseLessTighterThenAndOr(t *testing.T) {
	programme := "a and b less c"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(and (a) (less (b) (c)))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseLessLess(t *testing.T) {
	programme := "a less b less c"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(less (less (a) (b)) (c))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseManyLess(t *testing.T) {
	programme := "a less b less c less d less e"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(less (less (less (less (a) (b)) (c)) (d)) (e))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseParen(t *testing.T) {
	programme := "(a and b)"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "('(' (and (a) (b)))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseBypassesOrder(t *testing.T) {
	programme := "(a and b) less c"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(less ('(' (and (a) (b))) (c))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseRandomSubset(t *testing.T) {
	programme := "a%2"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(% (a) (2))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseRandomSubsetMultiple(t *testing.T) {
	programme := "a%2%1"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(% (% (a) (2)) (1))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParsePlusSimple(t *testing.T) {
	programme := "a+b"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(+ (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParsePlusBindsLikeAndOr(t *testing.T) {
	programme := "a or b + c and d"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(and (+ (or (a) (b)) (c)) (d))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseMinusSimple(t *testing.T) {
	programme := "a - b"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(- (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseMinusBindsTighterThenAndOr(t *testing.T) {
	programme := "a - b and c - d"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(and (- (a) (b)) (- (c) (d)))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseMinusBindsLikeLess(t *testing.T) {
	programme := "a - b less c - d"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(- (less (- (a) (b)) (c)) (d))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseFilterSimple(t *testing.T) {
	programme := "a filter b"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(filter (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseFilterTighterThenAnd(t *testing.T) {
	programme := "a and b filter c"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(and (a) (filter (b) (c)))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseFilterBindsLikeLess(t *testing.T) {
	programme := "a less b filter c"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(filter (less (a) (b)) (c))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

// example programmes
func TestParseProgrammeAndOr(t *testing.T) {
	programme := "dogs1.csv and dogs2.csv or cats.csv"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(or (and (dogs1.csv) (dogs2.csv)) (cats.csv))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

// errors
func TestRandomSubestRhsLetterError(t *testing.T) {
	defer func() {
		err, ok := recover().(error)
		if !ok {
			t.Fatalf("expected error but got nothing")
		}
		actual := err.Error()
		expected := "expected number but found a"
		if actual != expected {
			t.Errorf("expected error '%s' but got '%s'", expected, actual)
		}

	}()

	programme := "a%a"
	WalkParseTree(ParseProgramme(LexProgramme(programme)))

}

// helpers
func compareTrees(a *ParseTreeNode, b *ParseTreeNode) bool {
	if a.Token != b.Token {
		return false
	}

	if len(a.Children) != len(b.Children) {
		return false
	}

	for index := range a.Children {
		if !compareTrees(a.Children[index], b.Children[index]) {
			return false
		}
	}

	return true
}
