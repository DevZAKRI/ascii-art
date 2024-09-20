package ascii

import (
	"fmt"
	"strings"
)

// Artist is a function that takes an input string and read the standard file then pass both to PrintLineAsAscii.
// It reads the file containing the ASCII graph representation and puts the content in asciiGraph.
// If there is an error reading the file, it prints the error message and returns.
// It splits the input string into lines using the "\\n" delimiter.
func Artist(input string) {
	// check for errors too in case file no longer exist
	asciiGraph, err := ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(input, "\\n")
	// If the input string is empty after removing "\\n", it removes the first line.
	if strings.ReplaceAll(input, "\\n", "") == "" {
		lines = lines[1:]
	}
	for _, line := range lines {
		// prints each line as ASCII art using the PrintLineAsAscii function.
		PrintLineAsAscii(line, asciiGraph)
	}
}
