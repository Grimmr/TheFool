package Csv

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	table := NewCsv()
	table.Raw = "h1,h2,h3\na,b,c\nd,e,f"

	//0
	t.Logf("Row 0")
	row, cont := table.parseLine()
	expected := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expected, row) {
		t.Errorf("expected %v, got %v", expected, row)
	}
	if !cont {
		t.Errorf("unexpected end")
	}
	//1
	t.Logf("Row 1")
	row, cont = table.parseLine()
	expected = []string{"a", "b", "c"}
	if !reflect.DeepEqual(expected, row) {
		t.Errorf("expected %v, got %v", expected, row)
	}
	if !cont {
		t.Errorf("unexpected end")
	}
	//2
	t.Logf("Row 2")
	row, cont = table.parseLine()
	expected = []string{"d", "e", "f"}
	if !reflect.DeepEqual(expected, row) {
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
	expected := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expected, row) {
		t.Errorf("expected %v, got %v", expected, row)
	}
	if !cont {
		t.Errorf("unexpected end")
	}
	//1
	t.Logf("Row 1")
	row, cont = table.parseLine()
	expected = []string{"a", "b", "c"}
	if !reflect.DeepEqual(expected, row) {
		t.Errorf("expected %v, got %v", expected, row)
	}
	if !cont {
		t.Errorf("unexpected end")
	}
	//2
	t.Logf("Row 2")
	row, cont = table.parseLine()
	expected = []string{"d", "e", "f"}
	if !reflect.DeepEqual(expected, row) {
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

func TestRead(t *testing.T) {
	table := NewCsv()
	table.Read("test_data/dogs2.csv")

	t.Logf("headers")
	expectedHeaders := []string{"name", "breed", "age"}
	if !reflect.DeepEqual(expectedHeaders, table.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, table.Headers)
	}

	expectedData := [][]string{
		[]string{"spike", "greyhound", "2"},
		[]string{"clara", "wolfhound", "5"},
		[]string{"mike", "jack russel", "12"}}
	expectedTable := ConstructTable(expectedHeaders, expectedData)

	CompareData(expectedTable.Data, table.Data, true, t)
}

func TestReadTrailingNewLine(t *testing.T) {
	table := NewCsv()
	table.Read("test_data/dogs1.csv")

	t.Logf("headers")
	expectedHeaders := []string{"name", "breed", "age"}
	if !reflect.DeepEqual(expectedHeaders, table.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, table.Headers)
	}

	expectedData := [][]string{
		[]string{"spike", "greyhound", "2"},
		[]string{"clara", "wolfhound", "5"},
		[]string{"mike", "jack russel", "12"}}
	expectedTable := ConstructTable(expectedHeaders, expectedData)

	CompareData(expectedTable.Data, table.Data, true, t)
}

func TestReadDouble(t *testing.T) {
	table := NewCsv()
	table.Read("test_data/dogs1.csv")
	table.Read("test_data/people.csv")

	//headers
	t.Logf("headers")
	expectedHeaders := []string{"name", "age", "pay"}
	if !reflect.DeepEqual(expectedHeaders, table.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, table.Headers)
	}

	expectedData := [][]string{
		[]string{"tim", "22", "21000"},
		[]string{"tina", "40", "38000"},
		[]string{"clara", "35", "30000"}}

	expectedTable := ConstructTable(expectedHeaders, expectedData)

	CompareData(expectedTable.Data, table.Data, true, t)
}

func TestOperatorOrSimple(t *testing.T) {
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

func TestOperatorOrNewHeaderSimple(t *testing.T) {
	lhsheaders := []string{"h1", "lh"}
	lhsData := [][]string{[]string{"a", "b"}}
	lhs := ConstructTable(lhsheaders, lhsData)

	rhsheaders := []string{"h1", "rh"}
	rhsData := [][]string{[]string{"c", "d"}}
	rhs := ConstructTable(rhsheaders, rhsData)

	result := lhs.OperatorOr(rhs)

	expectedHeaders := []string{"h1", "lh", "rh"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{
		[]string{"a", "b", ""},
		[]string{"c", "", "d"}}
	expected := ConstructTable(expectedHeaders, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestToString(t *testing.T) {
	Headers := []string{"h", "l"}
	Data := [][]string{
		[]string{"a", "b"},
		[]string{"c", "d"}}
	table := ConstructTable(Headers, Data)

	actual := table.ToString()
	expected := "h, l\na, b\nc, d\n"
	if actual != expected {
		t.Errorf("\nexpected\n%s\ngot\n%s", expected, actual)
	}
}

func TestOperatorOrOperandsNotModified(t *testing.T) {
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

	if !reflect.DeepEqual(lhsHeaders, lhs.Headers) {
		t.Errorf("lhs headers: expected %v, got %v", lhsHeaders, lhs.Headers)
	}
	if !reflect.DeepEqual(rhsHeaders, rhs.Headers) {
		t.Errorf("rhs headers: expected %v, got %v", rhsHeaders, rhs.Headers)
	}

	t.Logf("lhs:")
	CompareData(ConstructTable(lhsHeaders, lhsData).Data, lhs.Data, true, t)

	t.Logf("\nrhs:")
	CompareData(ConstructTable(rhsHeaders, rhsData).Data, rhs.Data, true, t)
}

func TestOperatorAndSimple(t *testing.T) {
	headers := []string{"h1"}
	lhsData := [][]string{[]string{"a"}, []string{"b"}}
	lhs := ConstructTable(headers, lhsData)

	rhsData := [][]string{[]string{"b"}}
	rhs := ConstructTable(headers, rhsData)

	result := lhs.OperatorAnd(rhs)

	expectedData := [][]string{[]string{"b"}}
	expected := ConstructTable(headers, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestOperatorAndNewHeaderSimple(t *testing.T) {
	lhsheaders := []string{"h1", "lh"}
	lhsData := [][]string{[]string{"a", "b"}}
	lhs := ConstructTable(lhsheaders, lhsData)

	rhsheaders := []string{"h1", "rh"}
	rhsData := [][]string{[]string{"a", "d"}}
	rhs := ConstructTable(rhsheaders, rhsData)

	result := lhs.OperatorAnd(rhs)

	expectedHeaders := []string{"h1"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a"}}
	expected := ConstructTable(expectedHeaders, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestOperatorLessSimple(t *testing.T) {
	lhsheaders := []string{"h1"}
	lhsData := [][]string{[]string{"a"}, []string{"b"}}
	lhs := ConstructTable(lhsheaders, lhsData)

	rhsheaders := []string{"h1"}
	rhsData := [][]string{[]string{"b"}}
	rhs := ConstructTable(rhsheaders, rhsData)

	result := lhs.OperatorLess(rhs)

	expectedHeaders := []string{"h1"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a"}}
	expected := ConstructTable(expectedHeaders, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestOperatorRandomSubsetSimple(t *testing.T) {
	random := rand.New(rand.NewSource(8))

	lhsheaders := []string{"h1"}
	lhsData := [][]string{[]string{"a"}, []string{"b"}, []string{"c"}, []string{"d"}, []string{"e"}, []string{"f"}, []string{"g"}}
	lhs := ConstructTable(lhsheaders, lhsData)

	rhs := "3"

	result := lhs.OperatorRandomSubset(rhs, random)

	expectedHeaders := []string{"h1"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a"}, []string{"c"}, []string{"e"}}
	expectedTable := ConstructTable(expectedHeaders, expectedData)

	CompareData(expectedTable.Data, result.Data, true, t)
}

func TestOperatorRandomSubsetRhsToBig(t *testing.T) {
	random := rand.New(rand.NewSource(8))

	lhsheaders := []string{"h1"}
	lhsData := [][]string{[]string{"a"}, []string{"b"}, []string{"c"}}
	lhs := ConstructTable(lhsheaders, lhsData)

	rhs := "8"

	result := lhs.OperatorRandomSubset(rhs, random)

	expectedHeaders := []string{"h1"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a"}, []string{"b"}, []string{"c"}}
	expectedTable := ConstructTable(expectedHeaders, expectedData)

	CompareData(expectedTable.Data, result.Data, true, t)
}

func TestOperatorPlusSimple(t *testing.T) {
	headers := []string{"h1"}
	lhsData := [][]string{[]string{"a"}, []string{"b"}}
	lhs := ConstructTable(headers, lhsData)

	rhsData := [][]string{[]string{"a"}}
	rhs := ConstructTable(headers, rhsData)

	result := lhs.OperatorPlus(rhs)

	expectedData := [][]string{[]string{"a"}, []string{"a"}, []string{"b"}}
	expected := ConstructTable(headers, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestOperatorPlusMultiHeader(t *testing.T) {
	lhsHeaders := []string{"h1"}
	lhsData := [][]string{[]string{"a"}, []string{"b"}}
	lhs := ConstructTable(lhsHeaders, lhsData)

	rhsHeaders := []string{"h2"}
	rhsData := [][]string{[]string{"a"}}
	rhs := ConstructTable(rhsHeaders, rhsData)

	result := lhs.OperatorPlus(rhs)

	expectedHeaders := []string{"h1", "h2"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"", "a"}, []string{"a", ""}, []string{"b", ""}}
	expected := ConstructTable(expectedHeaders, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestOperatorMinusSimple(t *testing.T) {
	lhsHeaders := []string{"h1"}
	lhsData := [][]string{[]string{"a"}, []string{"b"}, []string{"b"}, []string{"b"}}
	lhs := ConstructTable(lhsHeaders, lhsData)

	rhsHeaders := []string{"h1", "h2"}
	rhsData := [][]string{[]string{"b", "a"}, []string{"b", "a"}}
	rhs := ConstructTable(rhsHeaders, rhsData)

	result := lhs.OperatorMinus(rhs)

	expectedHeaders := []string{"h1"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a"}, []string{"b"}}
	expected := ConstructTable(expectedHeaders, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestOperatorMinusTooManyInRhs(t *testing.T) {
	lhsHeaders := []string{"h1"}
	lhsData := [][]string{[]string{"a"}, []string{"b"}}
	lhs := ConstructTable(lhsHeaders, lhsData)

	rhsHeaders := []string{"h1"}
	rhsData := [][]string{[]string{"b"}, []string{"b"}}
	rhs := ConstructTable(rhsHeaders, rhsData)

	result := lhs.OperatorMinus(rhs)

	expectedHeaders := []string{"h1"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"a"}}
	expected := ConstructTable(expectedHeaders, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestOperatorFilterSimple(t *testing.T) {
	lhsHeaders := []string{"h1"}
	lhsData := [][]string{[]string{"a"}, []string{"b"}, []string{"b"}, []string{"b"}}
	lhs := ConstructTable(lhsHeaders, lhsData)

	rhsHeaders := []string{"h1"}
	rhsData := [][]string{[]string{"b"}, []string{"b"}}
	rhs := ConstructTable(rhsHeaders, rhsData)

	result := lhs.OperatorFilter(rhs)

	expectedHeaders := []string{"h1"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"b"}, []string{"b"}, []string{"b"}}
	expected := ConstructTable(expectedHeaders, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestOperatorFilterHeaderMissmatch(t *testing.T) {
	lhsHeaders := []string{"h1", "h2"}
	lhsData := [][]string{[]string{"a", "e"}, []string{"b", "f"}, []string{"b", "i"}, []string{"b", "i"}}
	lhs := ConstructTable(lhsHeaders, lhsData)

	rhsHeaders := []string{"h1", "h3"}
	rhsData := [][]string{[]string{"b", "a"}, []string{"b", "a"}, []string{"b", "a"}, []string{"b", "a"}, []string{"b", "a"}}
	rhs := ConstructTable(rhsHeaders, rhsData)

	result := lhs.OperatorFilter(rhs)

	expectedHeaders := []string{"h1", "h2"}
	if !reflect.DeepEqual(expectedHeaders, result.Headers) {
		t.Errorf("expected %v, got %v", expectedHeaders, result.Headers)
	}

	expectedData := [][]string{[]string{"b", "f"}, []string{"b", "i"}, []string{"b", "i"}}
	expected := ConstructTable(expectedHeaders, expectedData)

	CompareData(expected.Data, result.Data, true, t)
}

func TestRowLessThen(t *testing.T) {
	headers := []string{"1", "2", "3"}
	data := [][]string{
		[]string{"a", "b", "a"},
		[]string{"b", "b", "b"},
		[]string{"a", "a", "c"},
		[]string{"c", "c", "c"}}

	table := ConstructTable(headers, data)

	expected := true
	actual := rowLessThen(headers, table.Data[0], table.Data[3])
	if expected != actual {
		t.Errorf("test 1: expected %t, got %t", expected, actual)
	}

	expected = false
	actual = rowLessThen(headers, table.Data[0], table.Data[0])
	if expected != actual {
		t.Errorf("test 2: expected %t, got %t", expected, actual)
	}

	expected = false
	actual = rowLessThen(headers, table.Data[3], table.Data[0])
	if expected != actual {
		t.Errorf("test 3: expected %t, got %t", expected, actual)
	}
}

func TestGenerateIndex(t *testing.T) {
	headers := []string{"1", "2", "3"}
	data := [][]string{
		[]string{"a", "b", "a"},
		[]string{"b", "b", "b"},
		[]string{"a", "a", "c"},
		[]string{"c", "c", "c"}}

	table := ConstructTable(headers, data)
	table.generateIndex()

	expectedIndex := []int{2, 0, 1, 3}

	expectedLen := len(expectedIndex)
	actualLen := len(table.Index)
	if expectedLen != actualLen {
		t.Fatalf("expected Index length %d, got %d", expectedLen, actualLen)
	}

	if !reflect.DeepEqual(expectedIndex, table.Index) {
		t.Errorf("expected %d, got %d", expectedIndex, table.Index)
	}
}

func TestFitHeaders(t *testing.T) {
	headers := []string{"a", "b"}
	row := map[string]string{"a": "a", "c": "c"}

	expected := map[string]string{"a": "a", "b": ""}
	actual := fitHeaders(headers, row)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestMatchRow(t *testing.T) {
	headers := []string{"a", "b"}
	data := [][]string{
		[]string{"a", "b"},
		[]string{"b", "b"}}
	table := ConstructTable(headers, data)

	expected := true
	actual := matchRow(headers, table.Data[0], table.Data[0])
	if expected != actual {
		t.Errorf("test 1: expected %t, got %t", expected, actual)
	}

	expected = false
	actual = matchRow(headers, table.Data[0], table.Data[1])
	if expected != actual {
		t.Errorf("test 2: expected %t, got %t", expected, actual)
	}
}
