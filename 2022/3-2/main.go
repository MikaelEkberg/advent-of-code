package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
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

	result := 0

	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Err")
	}

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
					score = scoreLookup[character]
					result += score
				}
			}
		}

		start = start + 3
		end = end + 3
	}
	fmt.Println(result)
}
