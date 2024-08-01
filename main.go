package main

import Eulogist "Eulogist/eulogist"

func main() {
	err := Eulogist.Eulogist()
	if err != nil {
		panic(err)
	}
}
