package main

import (
	"os"
	"fmt"
)

func main() {
	//concatinate all parameters into the expresion we are going to evaluate
	var expr string
	for _, element := range os.Args[1:] {
		expr += element + " "
	}

	evaluateProgramme(expr)
}

func evaluateProgramme(prog string) {
	fmt.Println(prog)
	tokens := lexProgramme(prog)
	treeRoot := parseProgramme(tokens)
	os.Stdout.WriteString(walkParseTree(treeRoot))
	interpProgramme(treeRoot)
}