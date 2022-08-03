package Parser

import (
	"testing"
)

//lexer only tests
func TestLexProgrammeKeywords(t *testing.T) {
	//should include one of every token
	programme := "and or a.txt"
	result := LexProgramme(programme)
	expected := []lexToken{lexToken{"and", lexTokenType_and}, lexToken{"or", lexTokenType_or}, lexToken{"a.txt", lexTokenType_name}}

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
	programme := &ParseTreeNode{lexToken{"and", lexTokenType_and}, []*ParseTreeNode{&ParseTreeNode{lexToken{"or", lexTokenType_or}, []*ParseTreeNode{&ParseTreeNode{lexToken{"a", lexTokenType_name}, nil}, &ParseTreeNode{lexToken{"b", lexTokenType_name}, nil}}}, &ParseTreeNode{lexToken{"c", lexTokenType_name}, nil}}}
	result := WalkParseTree(programme)
	expected := "(and (or (a) (b)) (c))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

//parser only tests ()
//a simple one level test of and
func TestParseOnly(t *testing.T) {
	programme := []lexToken{lexToken{"a", lexTokenType_name}, lexToken{"and", lexTokenType_and}, lexToken{"b", lexTokenType_name}}
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
	if a.token != b.token {
		return false
	}

	if len(a.children) != len(b.children){
		return false
	}

	for index := range a.children {
		if !compareTrees(a.children[index], b.children[index]) {
			return false
		}
	} 

	return true
}