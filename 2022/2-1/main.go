package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var resultLookup = map[string]string{
		"A X": "draw",
		"A Y": "win",
		"A Z": "lose",
		"B X": "lose",
		"B Y": "draw",
		"B Z": "win",
		"C X": "win",
		"C Y": "lose",
		"C Z": "draw",
	}

	var scoreLookup = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	result := 0

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		res := strings.Split(fileScanner.Text(), " ")
		roundResult := 0
		yourScore := scoreLookup[res[1]]

		if resultLookup[fileScanner.Text()] == "win" {
			roundResult = yourScore + 6
		} else if resultLookup[fileScanner.Text()] == "draw" {
			roundResult = yourScore + 3
		} else if resultLookup[fileScanner.Text()] == "lose" {
			roundResult = yourScore + 0
		}
		result += roundResult
	}
	fmt.Println(result)
	readFile.Close()
}
