package Interp

import (
	"testing"
	"reflect"
)

func TestBufferTableGetOrLoadNewFile(t *testing.T) {
	table := NewBufferTable()
	csv := table.GetOrLoad("test_data/simple.csv")

	//we only check the headers here to verify the file
	expected := []string{"h1", "h2", "h3"}
	if !reflect.DeepEqual(expected,csv.Headers) {
		t.Errorf("expected %v, got %v", expected, csv.Headers)
	}
}

func TestBufferTableGetOrLoadOldFile(t *testing.T) {
	table := NewBufferTable()
	csvA := table.GetOrLoad("test_data/simple.csv")

	csvA.Data[1]["h1"] = "x"

	csvB := table.GetOrLoad("test_data/simple.csv")

	//we only check the headers here to verify the file
	expected := "x"
	actual := csvB.Data[1]["h1"]
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}