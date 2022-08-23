package Interp

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/Grimmr/TheFool/Csv"
	"github.com/Grimmr/TheFool/Parser"
)

func TestInterpName(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/simple.csv"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"}}

	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpOr(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/simple.csv or test_data/simple2.csv"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2", "h3", "h4"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{
		[]string{"a", "", "", "b"},
		[]string{"a", "b", "c", ""},
		[]string{"d", "e", "f", ""}}

	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpAnd(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/andData.csv and test_data/andData2.csv"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a", "b", "c"}}

	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpLess(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/andData.csv less test_data/andData2.csv"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"d", "e", "f"}}

	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpParen(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("(test_data/simple.csv)"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a", "b", "c"}, []string{"d", "e", "f"}}

	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpRandomSubset(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/simple.csv%1"))

	random := rand.NewSource(100)
	result := InterpProgramme(programme, nil, &random)

	expectedHeaders := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"d", "e", "f"}}

	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpMinus(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/minusData.csv - test_data/minusData2.csv"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a", "b"}, []string{"c", "d"}}

	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpPlus(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/simple.csv+test_data/simple2.csv"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2", "h3", "h4"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a", "", "", "b"}, []string{"a", "b", "c", ""}, []string{"d", "e", "f", ""}}

	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpFilter(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/minusData.csv filter test_data/minusData2.csv"))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a", "b"}, []string{"a", "b"}}

	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestInterpNameWithSpace(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("\"test_data/minus data.csv\" filter \"test_data/minus data2.csv\""))

	result := InterpProgramme(programme, nil, nil)

	expectedHeaders := []string{"h1", "h2"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a", "b"}, []string{"a", "b"}}

	Csv.CompareData(Csv.ConstructTable(expectedHeaders, expectedData).Data, result.Data, true, t)
}

func TestRegression1(t *testing.T) {
	programme := Parser.ParseProgramme(Parser.LexProgramme("test_data/regression1/in1.csv filter test_data/regression1/in2.csv"))

	result := InterpProgramme(programme, nil, nil)
	expect := Csv.NewCsv()
	expect.Read("test_data/regression1/out.csv")

	expectedHeaders := expect.Headers
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("headers: expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := expect.Data

	Csv.CompareData(expectedData, result.Data, true, t)
}
