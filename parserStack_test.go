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

