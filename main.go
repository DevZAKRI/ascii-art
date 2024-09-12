package main

import (
	"fmt"
	"os"

	ascii "ascii/artistTools"
)

func main() {
	// check if the input is valid
	if ascii.ValidateArgument(os.Args) != "" {
		fmt.Println(ascii.ValidateArgument(os.Args))
		return
	}

	input := os.Args[1]
	// take the input and and Process it
	ascii.Artist(input)
}
