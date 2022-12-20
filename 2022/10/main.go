package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	day := 10
	input := getInput("input.txt")

	solutionPartOne := solutionPartOne(input)

	printSolution(solutionPartOne, 2, day)

	solutionPartTwo(input)
}

func getInput(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	input := string(data)

	return input
}

func printSolution(firstPart int, secondPart int, day int) {
	fmt.Println("Day", day, "Part 1:", firstPart)
	fmt.Println("Day", day, "Part 2:", secondPart)
}

func solutionPartOne(input string) int {
	lines := strings.Split(input, "\n")
	value := 1
	cycles := 0
	var cyclesToRun, instructionValue, sum int
	var instructionType string

	for _, line := range lines {
		if strings.HasPrefix(line, "noop") {
			instructionType = "noop"
			cyclesToRun = 1
		} else {
			instruction := strings.Split(line, " ")
			instructionType = instruction[0]
			instructionValue, _ = strconv.Atoi(instruction[1])
			cyclesToRun = 2
		}

		for i := 0; i < cyclesToRun; i++ {
			cycles++

			if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
				res := cycles * value
				sum += res
			}
		}

		if instructionType == "addx" {
			value += instructionValue
		}
	}

	return sum
}

func solutionPartTwo(input string) {
	lines := strings.Split(input, "\n")
	spritePosition := 1
	cycles := 0
	var instructionType string
	var instructionValue, cyclesToRun int
	res := ""

	for _, line := range lines {
		spritePosition1 := spritePosition - 1
		spritePosition2 := spritePosition
		spritePosition3 := spritePosition + 1

		if strings.HasPrefix(line, "noop") {
			instructionType = "noop"
			cyclesToRun = 1
		} else {
			instruction := strings.Split(line, " ")
			instructionType = instruction[0]
			instructionValue, _ = strconv.Atoi(instruction[1])
			cyclesToRun = 2
		}

		for i := 0; i < cyclesToRun; i++ {
			// fmt.Println(cycles, "-", spritePosition1, spritePosition2, spritePosition3)

			if cycles == spritePosition1 || cycles == spritePosition2 || cycles == spritePosition3 {
				res += "#"
			} else {
				res += " "
			}

			if cycles == 39 {
				fmt.Println(res)
				res = ""
				cycles = 0
			} else {
				cycles++
			}
		}

		if instructionType == "addx" {
			spritePosition += instructionValue
		}
	}
}
