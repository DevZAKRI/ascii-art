package main

import (
	"fmt"
	"os"

	ascii "ascii/artistTools"
)

func main() {
	// check if the input is valid
	if ascii.ValidateInput() != "" {
		print(ascii.ValidateInput())
		return
	}

	input := os.Args[1]
	// read the file containing the ascii graph represantation and put content in ascii Graph
	// check for errors too in case file no longer exist
	asciiGraph, err := ascii.ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// take the input and asciiGraph and and Process them
	ascii.TheARTIST(input, asciiGraph)
}
