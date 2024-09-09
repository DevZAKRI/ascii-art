package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error u need to enter the STRING u want in a graphic representation using ASCII ")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Only trait one STRING at once!!")
		return
	}
	input := os.Args[1]
	if !validASCIIrune(input) {
		fmt.Println("input Contain invalid characters")
		return
	}

	AsciiGraph, _ := readFile("banners/thinkertoy.txt")

	fmt.Println(AsciiGraph)
}

// word := "éà^¨¨¨"'"
func validASCIIrune(word string) bool {
	for _, r := range word {
		if !(r >= 32 && r <= 126) && r != '\n' {
			return false
		}
	}
	return true
}

func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var content []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text()+"\n")
	}
	return content, scanner.Err()
}
