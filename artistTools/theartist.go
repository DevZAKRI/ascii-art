package ascii

import (
	"fmt"
	"strings"
)

func Artist(input string) {
	// read the file containing the ascii graph represantation and put content in ascii Graph
	// check for errors too in case file no longer exist
	asciiGraph, err := ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(input, "\\n")
	if strings.ReplaceAll(input, "\\n", "") == "" {
		lines = lines[1:]
	}
	for _, line := range lines {
		// Print Line by Line
		PrintLineAsAscii(line, asciiGraph)
	}
}
