package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	day := 3

	firstPartSolution := getSolutionPartOne("input.txt")
	secondPartSolution := getSolutionPartTwo("input.txt")

	printSolution(firstPartSolution, secondPartSolution, day)
}

func printSolution(firstPart int, secondPart int, day int) {
	fmt.Println("Day", day, "Part 1:", firstPart)
	fmt.Println("Day", day, "Part 2:", secondPart)
}

func getScore(char string) int {
	var scoreLookup = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}

	score := scoreLookup[char]

	return score
}

func getSolutionPartOne(input string) int {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	result := 0

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		last := len(text)
		half := len(text) / 2
		firstHalf := text[0:half]
		secondHalf := text[half:last]
		test := ""
		score := 0

		secondHalfArray := strings.Split(secondHalf, "")

		for _, character := range secondHalfArray {
			res := strings.Contains(firstHalf, character)

			if res {
				if !(strings.Contains(test, character)) {
					test += character
					score = getScore(character)
					result += score
				}
			}
		}
	}

	readFile.Close()

	return result
}

func getSolutionPartTwo(input string) int {
	content, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Err")
	}

	result := 0

	text := string(content)
	values := strings.Split(text, "\n")
	count := len(values)

	start := 0
	end := 3

	for end <= count {
		group := values[start:end]
		firstRucksack := group[0]
		secondRucksack := group[1]
		lastRucksackCharacters := strings.Split(group[2], "")
		test := ""
		score := 0

		for _, character := range lastRucksackCharacters {
			res := strings.Contains(firstRucksack, character) && strings.Contains(secondRucksack, character)

			if res {
				if !(strings.Contains(test, character)) {
					test += character
					score = getScore(character)
					result += score
				}
			}
		}

		start = start + 3
		end = end + 3
	}

	return result
}
