package main

import (
	"os"
	"fmt"
	"github.com/Grimmr/TheFool/Headless" 
)

func main() {
	//defer error printer
	defer func () {
		err := recover()
		if err != nil {
			fmt.Print(err)
			fmt.Print("\n")
		}
	}()
	
	if len(os.Args) > 1 {
		Headless.RunCmdMode(os.Args[1:])
	}
}