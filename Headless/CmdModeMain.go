package Headless

import (
	"fmt"
	"github.com/Grimmr/TheFool/Parser"
	"github.com/Grimmr/TheFool/Interp"
)

func RunCmdMode(args []string) {
	//concatinate all parameters into the expresion we are going to evaluate
	var expr string
	for _, element := range args {
		expr += element + " "
	}

	result := Interp.InterpProgramme(Parser.ParseProgramme(Parser.LexProgramme(expr)), nil)
	fmt.Print(result.ToString())
}