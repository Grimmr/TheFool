package main

import (
	"testing"
)

//lexer only tests
func TestLexProgrammeKeywords(t *testing.T) {
	//should include one of every token
	programme := "and or a.txt"
	result := lexProgramme(programme)
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

//walkParseTree tests
func TestWalkParseTree(t *testing.T) {
	programme := &parseTreeNode{lexToken{"and", lexTokenType_and}, []*parseTreeNode{&parseTreeNode{lexToken{"or", lexTokenType_or}, []*parseTreeNode{&parseTreeNode{lexToken{"a", lexTokenType_name}, nil}, &parseTreeNode{lexToken{"b", lexTokenType_name}, nil}}}, &parseTreeNode{lexToken{"c", lexTokenType_name}, nil}}}
	result := walkParseTree(programme)
	expected := "(and (or (a) (b)) (c))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

//parser only tests ()
//a simple one level test of and
func TestParseOnly(t *testing.T) {
	programme := []lexToken{lexToken{"a", lexTokenType_name}, lexToken{"and", lexTokenType_and}, lexToken{"b", lexTokenType_name}}
	result := walkParseTree(parseProgramme(programme)) 
	expected := "(and (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

//full stack tests
//operator checks
func TestParseAnd(t *testing.T) {
	programme := "a and b"
	result := walkParseTree(parseProgramme(lexProgramme(programme)))
	expected := "(and (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseOr(t *testing.T) {
	programme := "a or b"
	result := walkParseTree(parseProgramme(lexProgramme(programme)))
	expected := "(or (a) (b))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestParseAndOr(t *testing.T) {
	programme := "a and b or c"
	result := walkParseTree(parseProgramme(lexProgramme(programme)))
	expected := "(or (and (a) (b)) (c))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

//example programmes
func TestParseProgrammeAndOr(t *testing.T) {
	programme := "dogs1.csv and dogs2.csv or cats.csv"
	result := walkParseTree(parseProgramme(lexProgramme(programme)))
	expected := "(or (and (dogs1.csv) (dogs2.csv)) (cats.csv))"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

//helpers
func compareTrees(a *parseTreeNode, b *parseTreeNode) bool {
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