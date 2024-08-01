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

	pterm.Info.Println("\nProgram running down, now press enter to exit")
	fmt.Scanln()
}
