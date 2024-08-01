package main

import (
	Eulogist "Eulogist/eulogist"
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	err := Eulogist.Eulogist()
	if err != nil {
		pterm.Error.Println(err)
	}

	fmt.Println()
	pterm.Info.Println("Program running down, now press enter to exit.")
	fmt.Scanln()
}
