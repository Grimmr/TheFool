package Parser

import (
	"testing"
)

//lexer only tests
func TestLexProgrammeKeywords(t *testing.T) {
	//should include one of every token
	programme := "and or a.txt less()%"
	result := LexProgramme(programme)
	expected := []lexToken{lexToken{"and", LexTokenType_and}, lexToken{"or", LexTokenType_or}, lexToken{"a.txt", LexTokenType_name}, lexToken{"less", LexTokenType_less}, lexToken{"'('", LexTokenType_lParen}, lexToken{"')'", LexTokenType_rParen}, lexToken{"%", LexTokenType_percent}}

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

//WalkParseTree tests
func TestWalkParseTree(t *testing.T) {
	programme := &ParseTreeNode{lexToken{"and", LexTokenType_and}, []*ParseTreeNode{&ParseTreeNode{lexToken{"or", LexTokenType_or}, []*ParseTreeNode{&ParseTreeNode{lexToken{"a", LexTokenType_name}, nil}, &ParseTreeNode{lexToken{"b", LexTokenType_name}, nil}}}, &ParseTreeNode{lexToken{"c", LexTokenType_name}, nil}}}
	result := WalkParseTree(programme)
	expected := "(and (or (a) (b)) (c))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

//parser only tests ()
//a simple one level test of and
func TestParseOnly(t *testing.T) {
	programme := []lexToken{lexToken{"a", LexTokenType_name}, lexToken{"and", LexTokenType_and}, lexToken{"b", LexTokenType_name}}
	result := WalkParseTree(ParseProgramme(programme)) 
	expected := "(and (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

//full stack tests
//operator checks
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


//example programmes
func TestParseProgrammeAndOr(t *testing.T) {
	programme := "dogs1.csv and dogs2.csv or cats.csv"
	result := WalkParseTree(ParseProgramme(LexProgramme(programme)))
	expected := "(or (and (dogs1.csv) (dogs2.csv)) (cats.csv))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}


//helpers
func compareTrees(a *ParseTreeNode, b *ParseTreeNode) bool {
	if a.Token != b.Token {
		return false
	}

	if len(a.Children) != len(b.Children){
		return false
	}

	for index := range a.Children {
		if !compareTrees(a.Children[index], b.Children[index]) {
			return false
		}
	} 

	return true
}