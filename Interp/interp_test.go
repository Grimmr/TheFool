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

func TestInterpOr (t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/simple.csv or test_data/simple2.csv"))

	result := InterpProgramme(programme, nil)

	expectedHeaders := []string{"h1", "h2", "h3", "h4"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{
		[]string{"a", "", "", "b"},
		[]string{"a","b","c", ""},
		[]string{"d","e","f", ""}}
	
	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpAnd (t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/andData.csv and test_data/andData2.csv"))

	result := InterpProgramme(programme, nil)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a","b","c"}}
	
	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}