package ascii

import "fmt"

func PrintLineAsAscii(line string, idx int, asciiGraph []string) {
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
	} else if idx != 0 {
		fmt.Println()
	}
}

// find the last line after the char
func findLastLine(char rune) int {
	return int((char - 31) * (9))
}
