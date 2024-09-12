package ascii

import "fmt"

// PrintLineAsAscii prints a given line of text as ASCII art using a provided ASCII graph.
// The line parameter represents the text to be printed as ASCII art.
// The asciiGraph parameter is a slice of strings that represents the ASCII graph used for printing.
// Each character in the line is converted to ASCII art by looping through it to the corresponding lines in the asciiGraph.
// The ASCII art is printed line by line, with each line of the ASCII art represented by a string in the asciiGraph.
// If the line parameter is empty, a new line is printed.
func PrintLineAsAscii(line string, asciiGraph []string) {
	var asciiChars []string
	if line != "" {
		for _, char := range line {
			// add each chars to the asciiChars line by line
			for i := 8; i >= 0; i-- {
				asciiChars = append(asciiChars, string(asciiGraph[findLastLine(char)-i]))
			}
		}

		for i := 0; i < 8; i++ {
			for j := 0; j < len(asciiChars); j += 9 {
				fmt.Print(asciiChars[i+j])
			}
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}

// find the last line after the char
func findLastLine(char rune) int {
	return int((char - 31) * (9))
}
