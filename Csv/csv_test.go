package Csv

import (
	"testing"
	"reflect"
)

func TestParseLine(t *testing.T) {
	table := NewCsv()
	table.raw = "h1,h2,h3\na,b,c\nd,e,f"
	
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
	table.raw = "h1,h2,h3\na,b,c\nd,e,f\n"
	
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
	
	//row count
	if len(table.data) != 3 {
		t.Errorf("expected 3 rows, got %d", len(table.data))
	}
	//headers 
	t.Logf("headers")
	expected := []string{"name", "breed", "age"}
	if !reflect.DeepEqual(expected,table.headers) {
		t.Errorf("expected %v, got %v", expected, table.headers)
	}
	//row 0
	t.Logf("Row 0")
	expected = []string{"spike", "greyhound", "2"}
	act := []string{table.data[0]["name"], table.data[0]["breed"], table.data[0]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.data[0]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.data[0]))
	}
	//row 1
	t.Logf("Row 1")
	expected = []string{"clara", "wolfhound", "5"}
	act = []string{table.data[1]["name"], table.data[1]["breed"], table.data[1]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.data[1]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.data[1]))
	}
	//row 2
	t.Logf("Row 2")
	expected = []string{"mike", "jack russel", "12"}
	act = []string{table.data[2]["name"], table.data[2]["breed"], table.data[2]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.data[2]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.data[2]))
	}
}

func TestReadTrailingNewLine (t *testing.T) {
	table := NewCsv()
	table.Read("test_data/dogs1.csv")
	
	//row count
	if len(table.data) != 3 {
		t.Errorf("expected 3 rows, got %d", len(table.data))
	}
	//headers 
	t.Logf("headers")
	expected := []string{"name", "breed", "age"}
	if !reflect.DeepEqual(expected,table.headers) {
		t.Errorf("expected %v, got %v", expected, table.headers)
	}
	//row 0
	t.Logf("Row 0")
	expected = []string{"spike", "greyhound", "2"}
	act := []string{table.data[0]["name"], table.data[0]["breed"], table.data[0]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.data[0]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.data[0]))
	}
	//row 1
	t.Logf("Row 1")
	expected = []string{"clara", "wolfhound", "5"}
	act = []string{table.data[1]["name"], table.data[1]["breed"], table.data[1]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.data[1]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.data[1]))
	}
	//row 2
	t.Logf("Row 2")
	expected = []string{"mike", "jack russel", "12"}
	act = []string{table.data[2]["name"], table.data[2]["breed"], table.data[2]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.data[2]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.data[2]))
	}
}

func TestReadDouble (t *testing.T) {
	table := NewCsv()
	table.Read("test_data/dogs1.csv")
	table.Read("test_data/people.csv")

	//row count
	if len(table.data) != 3 {
		t.Errorf("expected 3 rows, got %d", len(table.data))
	}
	//headers 
	t.Logf("headers")
	expected := []string{"name", "age", "pay"}
	if !reflect.DeepEqual(expected,table.headers) {
		t.Errorf("expected %v, got %v", expected, table.headers)
	}
	//row 0
	t.Logf("Row 0")
	expected = []string{"tim", "22", "21000"}
	act := []string{table.data[0]["name"], table.data[0]["age"], table.data[0]["pay"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.data[0]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.data[0]))
	}
	//row 1
	t.Logf("Row 1")
	expected = []string{"tina", "40", "38000"}
	act = []string{table.data[1]["name"], table.data[1]["age"], table.data[1]["pay"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.data[1]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.data[1]))
	}
	//row 2
	t.Logf("Row 2")
	expected = []string{"clara", "35", "30000"}
	act = []string{table.data[2]["name"], table.data[2]["age"], table.data[2]["pay"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.data[2]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.data[2]))
	}
} 