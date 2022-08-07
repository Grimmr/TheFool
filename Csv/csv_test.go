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
	
	//row count
	if len(table.Data) != 3 {
		t.Errorf("expected 3 rows, got %d", len(table.Data))
	}
	//headers 
	t.Logf("headers")
	expected := []string{"name", "breed", "age"}
	if !reflect.DeepEqual(expected,table.Headers) {
		t.Errorf("expected %v, got %v", expected, table.Headers)
	}
	//row 0
	t.Logf("Row 0")
	expected = []string{"spike", "greyhound", "2"}
	act := []string{table.Data[0]["name"], table.Data[0]["breed"], table.Data[0]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.Data[0]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.Data[0]))
	}
	//row 1
	t.Logf("Row 1")
	expected = []string{"clara", "wolfhound", "5"}
	act = []string{table.Data[1]["name"], table.Data[1]["breed"], table.Data[1]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.Data[1]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.Data[1]))
	}
	//row 2
	t.Logf("Row 2")
	expected = []string{"mike", "jack russel", "12"}
	act = []string{table.Data[2]["name"], table.Data[2]["breed"], table.Data[2]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.Data[2]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.Data[2]))
	}
}

func TestReadTrailingNewLine (t *testing.T) {
	table := NewCsv()
	table.Read("test_data/dogs1.csv")
	
	//row count
	if len(table.Data) != 3 {
		t.Errorf("expected 3 rows, got %d", len(table.Data))
	}
	//headers 
	t.Logf("headers")
	expected := []string{"name", "breed", "age"}
	if !reflect.DeepEqual(expected,table.Headers) {
		t.Errorf("expected %v, got %v", expected, table.Headers)
	}
	//row 0
	t.Logf("Row 0")
	expected = []string{"spike", "greyhound", "2"}
	act := []string{table.Data[0]["name"], table.Data[0]["breed"], table.Data[0]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.Data[0]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.Data[0]))
	}
	//row 1
	t.Logf("Row 1")
	expected = []string{"clara", "wolfhound", "5"}
	act = []string{table.Data[1]["name"], table.Data[1]["breed"], table.Data[1]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.Data[1]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.Data[1]))
	}
	//row 2
	t.Logf("Row 2")
	expected = []string{"mike", "jack russel", "12"}
	act = []string{table.Data[2]["name"], table.Data[2]["breed"], table.Data[2]["age"]}
	if !reflect.DeepEqual(expected,act) {
		t.Errorf("expected %v, got %v", expected, act)
	}
	if len(table.Data[2]) != len(expected) {
		t.Errorf("expected width %d, got %d", len(expected), len(table.Data[2]))
	}
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

	expectedTable := constructTable(expectedHeaders, expectedData)

	compareData(expectedTable.Data, table.Data, true, t)
}

func TestOperatorOrSimple (t *testing.T) {
	headers := []string{"h1"}
	lhsData := [][]string{[]string{"a"}}
	lhs := constructTable(headers, lhsData)

	rhsData := [][]string{[]string{"b"}}
	rhs := constructTable(headers, rhsData)

	result := lhs.OperatorOr(rhs)

	expectedData := [][]string{
		[]string{"a"},
		[]string{"b"}}
	expected := constructTable(headers, expectedData)

	compareData(expected.Data, result.Data, true, t)
}


//helpers
func constructTable(headers []string, data [][]string) *Csv {
	out := NewCsv()
	out.Headers = headers

	for _, row := range data {
		newRow := make(map[string]string)
		for index,  field := range row {
			newRow[headers[index]] = field
		}
		out.Data = append(out.Data, newRow)
	}

	return out
}

func compareData(expected []map[string]string, actual []map[string]string, checkAllWidths bool, t *testing.T) {
	if len(expected) != len(actual) {
		t.Fatalf("expected %d rows, but got %d", len(expected), len(actual))
	}

	var checkWidth = true
	for index := range expected {
		if checkWidth && len(expected[index]) != len(actual[index]) {
			t.Errorf("expected row %d to have %d fields, but found %d", index, len(expected[index]), len(actual[index]))
			checkWidth = checkAllWidths
		}

		for k, _ := range expected[index] {
			if expected[index][k] != actual[index][k] {
				t.Errorf("expected row %d, field %s to have value %s, but found %s", index, k, expected[index][k], actual[index][k])
			}
		}
	} 
}