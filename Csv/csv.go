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