package Csv

import(
	"io/ioutil"
	"strings"
)

type Csv struct {
	Headers []string
	Data []map[string]string
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

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
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
	checkErr(err)
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
	var newRow map[string]string
	for k, v := range row {
		newRow[k] = v
	}

	//add the row
	this.Data = append(this.Data, newRow) 
}

func (this *Csv) OperatorOr(rhs *Csv) *Csv {
	out := NewCsv()

	//select Headers
	///copy lhs headers
	for _, element := range this.Headers {
		out.insertHeader(element)
	}
	//add only needed rhs headers
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
			if matchRow(targetRow, checkRow) {
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

func matchRow(lhs map[string]string, rhs map[string]string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for k, _ := range lhs {
		if _, ok := rhs[k]; !ok {
			return false
		}
		if lhs[k] != rhs[k] {
			return false
		}
	}
	return true
} 