package Interp

import (
	"testing"
	"reflect"
	"math/rand"
	"github.com/Grimmr/TheFool/Parser"
	"github.com/Grimmr/TheFool/Csv"
)

func TestInterpName (t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/simple.csv"))

	result := InterpProgramme(programme, nil, nil)

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

	result := InterpProgramme(programme, nil, nil)

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

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a","b","c"}}
	
	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpLess (t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/andData.csv less test_data/andData2.csv"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"d","e","f"}}
	
	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t) 
}

func TestInterpParen (t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("(test_data/simple.csv)"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a","b","c"}, []string{"d","e","f"}}
	
	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t) 
}

func TestInterpRandomSubset (t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/simple.csv%1"))

	random := rand.NewSource(100)
	result := InterpProgramme(programme, nil, &random)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"d","e","f"}}
	
	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t) 
}

func TestInterpMinus (t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/minusData.csv - test_data/minusData2.csv"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a","b"}, []string{"c", "d"}}
	
	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t) 
}

func TestInterpPlus (t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/simple.csv+test_data/simple2.csv"))

	random := rand.NewSource(100)
	result := InterpProgramme(programme, nil, &random)

	expectedHeaders := []string{"h1", "h2", "h3", "h4"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a", "", "", "b"}, []string{"a","b","c",""}, []string{"d","e","f",""}}
	
	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t) 
}