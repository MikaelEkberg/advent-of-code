package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var resultLookup = map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}

	var lose = map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	var draw = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	var win = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
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
		yourScore := 0
		opponentChoice := res[0]
		neededResult := resultLookup[res[1]]

		if neededResult == "lose" {
			yourChoice := lose[opponentChoice]
			yourScore = scoreLookup[yourChoice] + 0
		} else if neededResult == "draw" {
			yourChoice := draw[opponentChoice]
			yourScore = scoreLookup[yourChoice] + 3
		} else if neededResult == "win" {
			yourChoice := win[opponentChoice]
			yourScore = scoreLookup[yourChoice] + 6
		}
		result += yourScore
	}
	fmt.Println(result)
	readFile.Close()
}
