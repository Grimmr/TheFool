package main

import (
	"os"
	"github.com/Grimmr/TheFool/Headless" 
)

func main() {
	if len(os.Args) > 1 {
		Headless.RunCmdMode(os.Args[1:])
	}
}