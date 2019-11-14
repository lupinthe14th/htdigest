package main

import (
	"os"
	"strings"
)

func split(s string) []string {
	return strings.Split(s, " ")
}

func Example_main() {

	os.Args = split("./example john john password")
	main()

	// Output:
	// john:john:8ebe5fdc92a961c3b8fb190db08e1cd5
}
