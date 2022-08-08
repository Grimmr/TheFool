package Csv

import (
	"testing"
	"reflect"
)

func TestParseLine(t *testing.T) {
	table := NewCsv()
	table.Raw = "h1,h2,h3\na,b,c\nd,e,f"
	
	//0
	t.Logf("Row 0")
	row, cont := table.parseLine()
	expected := []string{"h1","h2","h3"}
	if !reflect.DeepEqual(expected,row) {
		t.Errorf("expected %v, got %v", expected, row)
	}
	if !cont {
		t.Errorf("unexpected end")
	}
	//1
	t.Logf("Row 1")
	row, cont = table.parseLine()
	expected = []string{"a","b","c"}
	if !reflect.DeepEqual(expected,row) {
		t.Errorf("expected %v, got %v", expected, row)
	}
	if !cont {
		t.Errorf("unexpected end")
	}
	//2
	t.Logf("Row 2")
	row, cont = table.parseLine()
	expected = []string{"d","e","f"}
	if !reflect.DeepEqual(expected,row) {
		t.Errorf("expected %v, got %v", expected, row)
	}
	if !cont {
		t.Errorf("unexpected end")
	}
	//3
	t.Logf("Row 3")
	row, cont = table.parseLine()
	if len(row) != 0 {
		t.Errorf("expected empty, got %d", len(row))
	}
	if cont {
		t.Errorf("expected end, got no end")
	}
}

func TestParseLineTrailingNewLine(t *testing.T) {
	table := NewCsv()
	table.Raw = "h1,h2,h3\na,b,c\nd,e,f\n"
	
	//0
	t.Logf("Row 0")
	row, cont := table.parseLine()
	expected := []string{"h1","h2","h3"}
	if !reflect.DeepEqual(expected,row) {
		t.Errorf("expected %v, got %v", expected, row)
	}
	if !cont {
		t.Errorf("unexpected end")
	}
	//1
	t.Logf("Row 1")
	row, cont = table.parseLine()
	expected = []string{"a","b","c"}
	if !reflect.DeepEqual(expected,row) {
		t.Errorf("expected %v, got %v", expected, row)
	}
	if !cont {
		t.Errorf("unexpected end")
	}
	//2
	t.Logf("Row 2")
	row, cont = table.parseLine()
	expected = []string{"d","e","f"}
	if !reflect.DeepEqual(expected,row) {
		t.Errorf("expected %v, got %v", expected, row)
	}
	if !cont {
		t.Errorf("unexpected end")
	}
	//3
	t.Logf("Row 3")
	row, cont = table.parseLine()
	if len(row) != 0 {
		t.Errorf("expected empty, got %d", len(row))
	}
	if cont {
		t.Errorf("expected end, got no end")
	}
}

func TestRead (t *testing.T) {
	table := NewCsv()
	table.Read("test_data/dogs2.csv")

	t.Logf("headers")
	expectedHeaders := []string{"name", "breed", "age"}
	if !reflect.DeepEqual(expectedHeaders,table.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, table.Headers)
	}

	expectedData := [][]string{
		[]string{"spike", "greyhound", "2"},
		[]string{"clara", "wolfhound", "5"},
		[]string{"mike", "jack russel", "12"}}
	expectedTable := ConstructTable(expectedHeaders, expectedData)

	CompareData(expectedTable.Data, table.Data, true, t)
}

func TestReadTrailingNewLine (t *testing.T) {
	table := NewCsv()
	table.Read("test_data/dogs1.csv")

	t.Logf("headers")
	expectedHeaders := []string{"name", "breed", "age"}
	if !reflect.DeepEqual(expectedHeaders,table.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, table.Headers)
	}

	expectedData := [][]string{
		[]string{"spike", "greyhound", "2"},
		[]string{"clara", "wolfhound", "5"},
		[]string{"mike", "jack russel", "12"}}
	expectedTable := ConstructTable(expectedHeaders, expectedData)

	CompareData(expectedTable.Data, table.Data, true, t)
}

func TestReadDouble (t *testing.T) {
	table := NewCsv()
	table.Read("test_data/dogs1.csv")
	table.Read("test_data/people.csv")

	//headers 
	t.Logf("headers")
	expectedHeaders := []string{"name", "age", "pay"}
	if !reflect.DeepEqual(expectedHeaders,table.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, table.Headers)
	}

	expectedData := [][]string{
		[]string{"tim", "22", "21000"},
		[]string{"tina", "40", "38000"},
		[]string{"clara", "35", "30000"}}

	expectedTable := ConstructTable(expectedHeaders, expectedData)

	CompareData(expectedTable.Data, table.Data, true, t)
}

func TestOperatorOrSimple (t *testing.T) {
	headers := []string{"h1"}
	lhsData := [][]string{[]string{"a"}}
	lhs := ConstructTable(headers, lhsData)

	rhsData := [][]string{[]string{"b"}}
	rhs := ConstructTable(headers, rhsData)

	result := lhs.OperatorOr(rhs)

	expectedData := [][]string{
		[]string{"a"},
		[]string{"b"}}
	expected := ConstructTable(headers, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestOperatorOrNewHeaderSimple (t *testing.T) {
	lhsheaders := []string{"h1", "lh"}
	lhsData := [][]string{[]string{"a", "b"}}
	lhs := ConstructTable(lhsheaders, lhsData)

	rhsheaders := []string{"h1", "rh"}
	rhsData := [][]string{[]string{"c", "d"}}
	rhs := ConstructTable(rhsheaders, rhsData)

	result := lhs.OperatorOr(rhs)

	expectedHeaders := []string{"h1", "lh", "rh"}
	if !reflect.DeepEqual(expectedHeaders,result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{
		[]string{"a", "b", ""}, 
		[]string{"c", "", "d"}}
	expected := ConstructTable(expectedHeaders, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestOperatorOrOperandsNotModified (t *testing.T) {
	lhsHeaders := []string{"h", "l"}
	lhsData := [][]string{
		[]string{"a", "b"},
		[]string{"c", "d"}}
	lhs := ConstructTable(lhsHeaders, lhsData)

	rhsHeaders := []string{"h", "r"}
	rhsData := [][]string{
		[]string{"e", "f"},
		[]string{"g", "h"}}
	rhs := ConstructTable(rhsHeaders, rhsData)

	lhs.OperatorOr(rhs)

	if !reflect.DeepEqual(lhsHeaders,lhs.Headers) {
		t.Errorf("lhs headers: expected %v, got %v", lhsHeaders, lhs.Headers)
	}
	if !reflect.DeepEqual(rhsHeaders,rhs.Headers) {
		t.Errorf("rhs headers: expected %v, got %v", rhsHeaders, rhs.Headers)
	}

	t.Logf("lhs:")
	CompareData(ConstructTable(lhsHeaders, lhsData).Data, lhs.Data, true, t)

	t.Logf("\nrhs:")
	CompareData(ConstructTable(rhsHeaders, rhsData).Data, rhs.Data, true, t)
}