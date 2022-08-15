package Csv

import(
	"io/ioutil"
	"strings"
	"testing"
	"os"
	"errors"
	"math/rand"
	"strconv"
	"sort"
	//"fmt"
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
	this.Index = []int{}
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

func (this *Csv) generateIndex () {
	this.Index = []int{}
	unsorted := make([]int, len(this.Data))
	
	for index, _ := range this.Data {
		unsorted[index] = index
	}

	for _, _ = range this.Data {
		selected := 0
		for index, _ := range unsorted {
			if rowLessThen(this.Headers, this.Data[unsorted[index]], this.Data[unsorted[selected]]) {
				selected = index
			}
		}
		this.Index = append(this.Index, unsorted[selected])
		unsorted[selected] = unsorted[len(unsorted)-1]
		unsorted = unsorted[:len(unsorted)-1]

	}
}

func rowLessThen (headers []string, lhs map[string]string, rhs map[string]string) bool {
	for _, field := range headers {
		if lhs[field] < rhs[field] {
			return true
		} else if lhs[field] > rhs[field] {
			return false
		}
	}
	return false
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

func fitHeaders (headers []string, oldRow map[string]string) map[string]string {
	newRow := map[string]string{}
	for k, v := range oldRow {
		newRow[k] = v
	}
	
	//add new headers
	for _, header := range headers {
		if _, ok := newRow[header]; !ok {
			newRow[header] = ""
		}
	}

	//remove unneeded headers
	for k, _ := range newRow {
		remove := true
		for _, header := range headers {
			if k == header {
				remove = false
				break
			}
		}
		if remove {
			delete(newRow, k)
		}
	}

	return newRow
}

func (this *Csv) insertRow(row map[string]string) {
	//add the row
	this.Data = append(this.Data, fitHeaders(this.Headers, row)) 

	//for now we assume index position is guaranteed 
	this.Index = append(this.Index, len(this.Index))
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

	lhsPos := 0
	rhsPos := 0
	for lhsPos < len(this.Data) && rhsPos < len(rhs.Data) {
		lhsRow := fitHeaders(out.Headers, this.Data[this.Index[lhsPos]])
		rhsRow := fitHeaders(out.Headers, rhs.Data[rhs.Index[rhsPos]])
		var addedRow map[string]string 

		if rowLessThen(out.Headers, lhsRow, rhsRow){
			out.insertRow(lhsRow)
			addedRow = lhsRow
		} else if rowLessThen(out.Headers, rhsRow, lhsRow) {
			out.insertRow(rhsRow)
			addedRow = rhsRow
		} else {
			out.insertRow(lhsRow)
			addedRow = lhsRow
		}

		for lhsPos < len(this.Data) {
			if matchRow(out.Headers, fitHeaders(out.Headers, this.Data[this.Index[lhsPos]]), addedRow) {
				lhsPos++
			} else {
				break
			}
		}

		for rhsPos < len(rhs.Data) {
			if matchRow(out.Headers, fitHeaders(out.Headers, rhs.Data[rhs.Index[rhsPos]]), addedRow) {
				rhsPos++
			} else {
				break
			}
		}
	} 

	//only one of these will trigger
	for ; lhsPos < len(this.Data); lhsPos++ {
		out.insertRow(this.Data[this.Index[lhsPos]])
	}
	for ; rhsPos < len(rhs.Data); rhsPos++ {
		out.insertRow(rhs.Data[rhs.Index[rhsPos]])
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
	lhsPos := 0
	rhsPos := 0
	for lhsPos < len(this.Data) && rhsPos < len(rhs.Data) {
		lhsRow := fitHeaders(out.Headers, this.Data[this.Index[lhsPos]])
		rhsRow := fitHeaders(out.Headers, rhs.Data[rhs.Index[rhsPos]])
		var addedRow map[string]string 

		if matchRow(out.Headers, lhsRow, rhsRow){
			out.insertRow(lhsRow)
			addedRow = lhsRow

			for lhsPos < len(this.Data) {
				if matchRow(out.Headers, fitHeaders(out.Headers, this.Data[this.Index[lhsPos]]), addedRow) {
					lhsPos++
				} else {
					break
				}
			}

			for rhsPos < len(rhs.Data) {
				if matchRow(out.Headers, fitHeaders(out.Headers, rhs.Data[rhs.Index[rhsPos]]), addedRow) {
					rhsPos++
				} else {
					break
				}
			}
		} else if rowLessThen(out.Headers, lhsRow, rhsRow) {
			lhsPos++
		} else if rowLessThen(out.Headers, rhsRow, lhsRow) {
			rhsPos++
		}
	} 

	return out
}

func (this *Csv) OperatorLess(rhs *Csv) *Csv {
	out := NewCsv()

	//select headers
	compHeaders := []string{}
	for _, field := range this.Headers {
		out.insertHeader(field)
		for _, rhsField := range rhs.Headers {
			if field == rhsField {
				compHeaders = append(compHeaders, field)
				break
			}
		}
	}

	//select from lhs
	lhsPos := 0
	rhsPos := 0
	for lhsPos < len(this.Data) && rhsPos < len(rhs.Data) {
		lhsRow := fitHeaders(compHeaders, this.Data[this.Index[lhsPos]])
		rhsRow := fitHeaders(compHeaders, rhs.Data[rhs.Index[rhsPos]])

		if matchRow(out.Headers, lhsRow, rhsRow) {
			lhsPos++
		} else if rowLessThen(out.Headers, lhsRow, rhsRow) {
			out.insertRow(this.Data[this.Index[lhsPos]])
			lhsPos++
		} else  {
			rhsPos++
		}
	}

	//add the rest of lhs if any
	for lhsPos < len(this.Data) {
		out.insertRow(this.Data[this.Index[lhsPos]])
		lhsPos++
	}

	return out
}

func (this *Csv) operatorRandomSubset(rhs string, random *rand.Rand) *Csv {
	out := NewCsv()

	//select headers
	for _, field := range this.Headers {
		out.insertHeader(field)
	}

	//select data 
	targetCount, _ := strconv.Atoi(rhs)
	if targetCount < len(this.Data) {
		targets := make([]int, 0)
		options := make([]int, len(this.Data))
		for i := range options {
			options[i] = i
		}
		for i := 0; i < targetCount; i++ {
			selected := random.Intn(len(options))
			targets = append(targets, options[selected])
			options[selected] = options[len(options)-1]
			options = options[:len(options)-1]
		}
		sort.Ints(targets)
		for _, row := range targets {
			out.insertRow(this.Data[this.Index[row]])
		} 
	} else {
		for i := range this.Index {
			out.insertRow(this.Data[this.Index[i]])
		}
	}

	return out
}

func matchRow(headers []string, lhs map[string]string, rhs map[string]string) bool {
	for _, k := range headers {
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

	out.generateIndex()

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