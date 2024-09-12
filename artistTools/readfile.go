package ascii

import (
	"os"
)

func ReadFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var content []string
	singleByte := make([]byte, 1)
	line := ""
	for {
		_, err = file.Read(singleByte)
		if err != nil {
			break
		}
		if singleByte[0] == '\n' {
			content = append(content, line)
			line = ""
		} else {
			line += string(singleByte)
		}
	}
	content = append(content, line)
	// content = append(content, "")
	return content, nil
}
