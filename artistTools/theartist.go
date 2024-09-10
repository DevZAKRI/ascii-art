package ascii

import "strings"

func TheARTIST(input string, asciiGraph []string) {
	lines := strings.Split(input, "\\n")
	for idx, line := range lines {
		// Print Line by Line
		PrintLineAsAscii(line, idx, asciiGraph)
	}
}
