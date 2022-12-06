package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	day := 2

	firstPartSolution, secondPartSolution := getSolution("input.txt")

	printSolution(firstPartSolution, secondPartSolution, day)
}

func printSolution(firstPart int, secondPart int, day int) {
	fmt.Println("Day", day, "Part 1:", firstPart)
	fmt.Println("Day", day, "Part 2:", secondPart)
}

func resultLookupPartOne(input string) string {
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

	result := resultLookup[input]

	return result
}

func resultLookupPartTwo(input string) string {
	resultLookup := map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}

	result := resultLookup[input]

	return result
}

func getMove(neededResult string, opponentChoice string) string {
	lose := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	draw := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	win := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	var move string

	if neededResult == "lose" {
		move = lose[opponentChoice]
	} else if neededResult == "draw" {
		move = draw[opponentChoice]
	} else if neededResult == "win" {
		move = win[opponentChoice]
	}

	return move
}

func getScore(result string, move string) int {
	scoreLookup := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	var score int

	if result == "lose" {
		score = scoreLookup[move] + 0
	} else if result == "draw" {
		score = scoreLookup[move] + 3
	} else if result == "win" {
		score = scoreLookup[move] + 6
	}

	return score
}

func getSolutionPartOne(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var result int

	for fileScanner.Scan() {
		res := strings.Split(fileScanner.Text(), " ")
		roundResult := 0
		outcome := resultLookupPartOne(fileScanner.Text())
		roundResult = getScore(outcome, res[1])
		result += roundResult
	}

	readFile.Close()

	return result
}

func getSolutionPartTwo(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := 0

	for fileScanner.Scan() {
		res := strings.Split(fileScanner.Text(), " ")
		opponentChoice := res[0]
		neededResult := resultLookupPartTwo(res[1])

		move := getMove(neededResult, opponentChoice)
		result += getScore(neededResult, move)
	}

	readFile.Close()

	return result
}

func getSolution(input string) (int, int) {
	firstPart := getSolutionPartOne(input)
	secondPart := getSolutionPartTwo(input)

	return firstPart, secondPart
}
