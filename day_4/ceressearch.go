package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")

	allLines := strings.Split(string(file), "\n")

	words := findNumberOfWords(allLines)
	fmt.Println("Number of word XMAS in input file: ", words)
}

func findNumberOfWords(allLines []string) int {
	words := 0
	for i, line := range allLines {
		for j, char := range line {
			if string(char) == "X" && 
				checkLetter(allLines, i, j) {
				continue
			}
		}
	}

	return words
}

func checkLetter(allLines []string, i, j int) bool {
	h := len(allLines)
	w := len(allLines[0])
	// check forward:
	if j < w - 1 {

	}
	// check backward:
	if j > 0 {

	} 
	// check top_down:
	if i < h - 1 {
		
	} 
	// check bottom_up:
	if i > 0 {
		
	}
	// check top_down left:
	if i > 0 && j > 0 {

	}

	// check top_down right:
	if i > 0 && j < w - 1 {

	}
	// check bottom_up left:
	if i < h - 1 && j > 0 {

	}
	// check bottom_up right:
	if i < h -1 && j < w - 1 {

	}
	return false
}