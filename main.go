package main

func main() {
	err := UnfoldEulogist("server_code", "server_password", `type_your_token`, "https://user.fastbuilder.pro")
	if err != nil {
		panic(err)
	}
}
