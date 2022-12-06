package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Err")
	}

	input := string(content)
	day := 6

	firstPartSolution := getLastCharacterIndex(input, 4)
	secondPartSolution := getLastCharacterIndex(input, 14)

	printSolution(firstPartSolution, secondPartSolution, day)
}

func printSolution(firstPart int, secondPart int, day int) {
	fmt.Println("Day", day, "Part 1:", firstPart)
	fmt.Println("Day", day, "Part 2:", secondPart)
}

func getLastCharacterIndex(text string, characterCount int) int {
	length := len(text)
	lastCharacterIndex := 0
	solution := 0

out:
	for i := 0; lastCharacterIndex < length; i++ {
		lastCharacterIndex = i + characterCount
		fourCharacters := text[i:lastCharacterIndex]
		total := 0

		chars := []rune(fourCharacters)

		for j := 0; j < len(chars); j++ {
			char := string(chars[j])
			charCount := strings.Count(fourCharacters, char)
			total += charCount
		}

		if total == characterCount {
			solution = lastCharacterIndex
			break out
		}
	}
	return solution
}
