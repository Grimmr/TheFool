package Csv

import(
	"io/ioutil"
	"strings"
	"testing"
	"os"
	"errors"
)

type Csv struct {
	Headers []string
	Data []map[string]string
	Index []int
	Raw string
	RawPos int
	EndRow byte
	EndField byte
}

func NewCsv () *Csv {
	var out Csv
	out.EndRow = byte('\n')
	out.EndField = byte(',')
	out.empty()
	return &out
}

func (this *Csv) empty () {
	this.Headers = []string{}
	this.Data = []map[string]string{}
	this.Raw = ""
	this.RawPos = 0
}

func (this *Csv) Read (path string) bool {
	//reset members
	this.empty()

	//attempt to read file into a buffer
	bytes, err := ioutil.ReadFile(path)
	
	if os.IsNotExist(err)  {
		panic(errors.New("file/path not found " + path))
	}

	this.Raw = string(bytes)


	var lineBuffer []string

	//parse headers
	lineBuffer, _ = this.parseLine()
	for _, col := range lineBuffer {
		this.Headers = append(this.Headers, strings.TrimSpace(col))
	}

	//read lines
	for lineBuffer, cont := this.parseLine(); cont; lineBuffer, cont = this.parseLine() {
		row := make(map[string]string) 
		//check row width
		if len(lineBuffer) != len(this.Headers) {
			continue
		}

		//create row
		for index, col := range lineBuffer {
			row[this.Headers[index]] = strings.TrimSpace(col)
		}

		//add row
		this.Data = append(this.Data, row)
	}

	this.generateIndex()

	return true
}

func (this *Csv) parseLine () ([]string, bool) {
	var out []string
	var buffer string

	//stop on blank line or EOF
	if this.RawPos >= len(this.Raw) || this.Raw[this.RawPos] == this.EndRow {
		return out, false
	}

	for ; this.RawPos < len(this.Raw); this.RawPos++ {
		char := this.Raw[this.RawPos]
		if char == this.EndField {
			out = append(out, buffer)
			buffer = ""
		} else if char == this.EndRow {
			out = append(out, buffer)
			buffer = ""
			this.RawPos++
			return out, true
		} else {
			buffer += string(char)
		}
	}
	if buffer != "" {
		out = append(out, buffer)
	}
	return out, len(out) != 0 
} 

func rowLessThen(headerOrder []string, lhs map[string]string, rhs map[string]string) bool {
	for _, header := range headerOrder {
		if lhs[header] < rhs[header] {
			return true
		} else if lhs[header] > rhs[header] {
			return false
		}
	}

	return false
}

func (this *Csv) generateIndex() {
	//make an initial index
	dumbIndex := make([]int, len(this.Data))
	for i := 0; i < len(dumbIndex); i++ {
		dumbIndex[i] = i
	}
	
	//we're going to use selection sort which in n^2. sooner or later this should be improved
	for _ = range this.Data {
		selected := 0
		for compareIndex := range dumbIndex {
			if rowLessThen(this.Headers, this.Data[dumbIndex[compareIndex]], this.Data[dumbIndex[selected]]) {
				selected = compareIndex
			}
		}
		this.Index = append(this.Index, dumbIndex[selected])

		dumbIndex[selected] = dumbIndex[len(dumbIndex)-1]
		dumbIndex = dumbIndex[:len(dumbIndex)-1]
	} 
}

func (this *Csv) ToString() string {
	var out string
	//headers
	for index, field := range this.Headers {
		if index != 0 {
			out += ", "
		}
		out += field
	}
	out += "\n"
	//data
	for _, row := range this.Data {
		first := true
		for _, field := range this.Headers {
			if !first {
				out += ", "
			}
			first = false
			out += row[field]
		}
		out += "\n" 
	}
	return out
}

//this assumes the header isn't a dupe
func (this *Csv) insertHeader(header string) {
	//add the header
	this.Headers = append(this.Headers, header)
	
	//add entries for each row
	for _, row := range this.Data {
		row[header] = ""
	}
}

func (this *Csv) insertRow(row map[string]string) {
	//create new row object
	newRow := make(map[string]string)
	for k, v := range row {
		newRow[k] = v
	} 

	//add new headers
	for _, header := range this.Headers {
		if _, ok := newRow[header]; !ok {
			newRow[header] = ""
		}
	}

	//remove unneeded headers
	for k, _ := range newRow {
		remove := true
		for _, header := range this.Headers {
			if k == header {
				remove = false
				break
			}
		}
		if remove {
			delete(newRow, k)
		}
	}

	//add the row
	this.Data = append(this.Data, newRow) 
}

func (this *Csv) OperatorOr(rhs *Csv) *Csv {
	out := NewCsv()

	//select Headers
	///copy lhs headers
	lhsHeaders := []string{} //this is all the ehaders only on lhs
	for _, element := range this.Headers {
		out.insertHeader(element)
		add := true
		for _, check := range rhs.Headers {
			if element == check {
				add = false
			}
		}
		if add {
			lhsHeaders = append(lhsHeaders, element)
		}
	}
	//add only needed rhs headers
	rhsHeaders := []string{} //this is all the headers only on rhs
	for _, targetHeader := range rhs.Headers {
		add := true
		for _, checkHeader := range this.Headers {
			if targetHeader == checkHeader {
				add = false
				break
			}
		}
		if add {
			out.insertHeader(targetHeader)
			rhsHeaders = append(rhsHeaders, targetHeader)
		}
	}

	//copy lhs into out
	for _, element := range this.Data {
		out.insertRow(element)
	}

	//copy only needed rhs rows
	for _, targetRow := range rhs.Data {
		add := true
		for _, checkRow := range this.Data {
			if matchRow(out.Headers, targetRow, checkRow) {
				add = false
				break
			}
		}
		if add {
			out.insertRow(targetRow)
		}
	}

	return out 
}

func (this *Csv) OperatorAnd(rhs *Csv) *Csv {
	out := NewCsv()

	//select headers
	for _, targetHeader := range this.Headers {
		for _, checkHeader := range rhs.Headers {
			if targetHeader == checkHeader {
				out.insertHeader(targetHeader)
				break
			}
		} 
	}

	//select rows
	for _, targetRow := range this.Data {
		for _, checkRow := range rhs.Data {
			if matchRow(out.Headers, targetRow, checkRow) {
				out.insertRow(targetRow)
			}
		}
	} 

	return out
}

func matchRow(headers []string, lhs map[string]string, rhs map[string]string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for k, _ := range lhs {
		skip := true
		for _, header := range headers {
			if k == header {
				skip = false
			}
		}
		if skip {
			continue
		}

		if _, ok := rhs[k]; !ok {
			return false
		}
		if lhs[k] != rhs[k] {
			return false
		}
	}
	return true
} 

//helpers that should be in a test file but cant due to export semantics
func ConstructTable(headers []string, data [][]string) *Csv {
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

func CompareData(expected []map[string]string, actual []map[string]string, checkAllWidths bool, t *testing.T) {
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