package Headless

import (
	"os"
	"github.com/Grimmr/TheFool/Parser"
)

func RunCmdMode(args []string) {
	//concatinate all parameters into the expresion we are going to evaluate
	var expr string
	for _, element := range args {
		expr += element + " "
	}

	os.Stdout.WriteString(Parser.WalkParseTree(Parser.ParseProgramme(Parser.LexProgramme(expr))) + "\n")
}