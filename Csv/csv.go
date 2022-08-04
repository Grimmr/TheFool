package Csv

import(
	"io/ioutil"
	"strings"
)

type Csv struct {
	headers []string
	data []map[string]string
	raw string
	rawPos int
	endRow byte
	endField byte
}

func NewCsv () *Csv {
	var out Csv
	out.endRow = byte('\n')
	out.endField = byte(',')
	return &out
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func (this *Csv) empty () {
	this.headers = []string{}
	this.data = []map[string]string{}
	this.raw = ""
	this.rawPos = 0
}

func (this *Csv) Read (path string) bool {
	//reset members
	this.empty()

	//attempt to read file into a buffer
	bytes, err := ioutil.ReadFile(path)
	checkErr(err)
	this.raw = string(bytes)


	var lineBuffer []string

	//parse headers
	lineBuffer, _ = this.parseLine()
	for _, col := range lineBuffer {
		this.headers = append(this.headers, strings.TrimSpace(col))
	}

	//read lines
	for lineBuffer, cont := this.parseLine(); cont; lineBuffer, cont = this.parseLine() {
		row := make(map[string]string) 
		//check row width
		if len(lineBuffer) != len(this.headers) {
			continue
		}

		//create row
		for index, col := range lineBuffer {
			row[this.headers[index]] = strings.TrimSpace(col)
		}

		//add row
		this.data = append(this.data, row)
	}

	return true
}

func (this *Csv) parseLine () ([]string, bool) {
	var out []string
	var buffer string

	//stop on blank line or EOF
	if this.rawPos >= len(this.raw) || this.raw[this.rawPos] == this.endRow {
		return out, false
	}

	for ; this.rawPos < len(this.raw); this.rawPos++ {
		char := this.raw[this.rawPos]
		if char == this.endField {
			out = append(out, buffer)
			buffer = ""
		} else if char == this.endRow {
			out = append(out, buffer)
			buffer = ""
			this.rawPos++
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