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
	length := len(input)
	lastCharacterIndex := 0

out:
	for i := 0; lastCharacterIndex < length; i++ {
		lastCharacterIndex = i + 14
		fourCharacters := input[i:lastCharacterIndex]
		total := 0

		chars := []rune(fourCharacters)

		for j := 0; j < len(chars); j++ {
			char := string(chars[j])
			charCount := strings.Count(fourCharacters, char)
			total += charCount
		}

		if total == 14 {
			fmt.Println(fourCharacters, lastCharacterIndex)
			break out
		}
	}
}
