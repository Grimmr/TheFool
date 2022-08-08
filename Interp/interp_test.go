package Interp

import (
	"testing"
	"reflect"
	"github.com/Grimmr/TheFool/Parser"
	"github.com/Grimmr/TheFool/Csv"
)

func TestInterpName (t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/simple.csv"))

	result := InterpProgramme(programme, nil)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{
		[]string{"a","b","c"},
		[]string{"d","e","f"}}
	
	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}