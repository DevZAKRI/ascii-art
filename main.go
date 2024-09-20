package main

import (
	"fmt"
	"os"

	ascii "ascii/artistTools"
)

var ColorMap = map[string]string{
	"black":     "\033[30m",
	"red":       "\033[31m",
	"green":     "\033[32m",
	"yellow":    "\033[33m",
	"blue":      "\033[34m",
	"magenta":   "\033[35m",
	"cyan":      "\033[36m",
	"white":     "\033[37m",
	"ColorSTOP": "\033[0m",
}

// main is the entry point of the program.
// it validates the input argument, prints any validation errors,
// and then calls the Artist function to process the input.
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
