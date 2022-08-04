package Interp

import (
	"github.com/Grimmr/TheFool/Csv"
)

type BufferTable struct {
	Table map[string]*Csv.Csv
}

func NewBufferTable() *BufferTable {
	var out BufferTable
	out.Table = make(map[string]*Csv.Csv)
	return &out
}

func (this *BufferTable) GetOrLoad (name string) *Csv.Csv {
	//return open buffer if it is in the table
	if csv, ok := this.Table[name]; ok {
		return csv
	}

	//open new buffer
	csv := Csv.NewCsv()
	csv.Read(name)
	this.Table[name] = csv

	return csv
}