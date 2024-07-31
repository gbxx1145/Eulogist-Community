package main

func main() {
	err := UnfoldEulogist("48285363", "", `type your fb token`)
	if err != nil {
		panic(err)
	}
}
