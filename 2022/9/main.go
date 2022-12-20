package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	day := 9
	input := getInput("input.txt")

	solutionPartOne := solutionPartOne(input)

	solutionPartTwo := solutionPartTwo(input)

	printSolution(solutionPartOne, solutionPartTwo, day)
}

func printSolution(firstPart int, secondPart int, day int) {
	fmt.Println("Day", day, "Part 1:", firstPart)
	fmt.Println("Day", day, "Part 2:", secondPart)
}

func getInput(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	input := string(data)

	return input
}

func contains(elements []string, name string) bool {
	for _, value := range elements {
		if value == name {
			return true
		}
	}
	return false
}

func solutionPartOne(input string) int {
	lines := strings.Split(input, "\n")

	var hX, hY, sX, sY int
	var sXYPlacements []string

	for _, line := range lines {
		move := strings.Split(line, " ")
		direction := move[0]
		steps, _ := strconv.Atoi(move[1])

		for i := 1; i <= steps; i++ {
			if direction == "R" {
				hX = hX + 1
			} else if direction == "L" {
				hX = hX - 1
			} else if direction == "U" {
				hY = hY + 1
			} else if direction == "D" {
				hY = hY - 1
			}

			if direction == "R" && hX >= sX+2 {
				sX = hX - 1
				sY = hY
			} else if direction == "L" && hX <= sX-2 {
				sX = hX + 1
				sY = hY
			} else if direction == "U" && hY >= sY+2 {
				sY = hY - 1
				sX = hX
			} else if direction == "D" && hY <= sY-2 {
				sY = hY + 1
				sX = hX
			}

			sXY := strconv.Itoa(sX) + ":" + strconv.Itoa(sY)

			if !(contains(sXYPlacements, sXY)) {
				sXYPlacements = append(sXYPlacements, sXY)
			}
		}
	}

	return len(sXYPlacements)
}

func solutionPartTwo(input string) int {
	lines := strings.Split(input, "\n")

	var hX, hY int
	var sXYPlacements []string
	knots := 9
	var knotPlacements = make(map[int]string)

	for i := 0; i < knots; i++ {
		knotPlacements[i] = "0:0"
	}

	for _, line := range lines {
		move := strings.Split(line, " ")
		direction := move[0]
		steps, _ := strconv.Atoi(move[1])

		fmt.Println(line)

		for i := 1; i <= steps; i++ {
			if direction == "R" {
				hX = hX + 1
			} else if direction == "L" {
				hX = hX - 1
			} else if direction == "U" {
				hY = hY + 1
			} else if direction == "D" {
				hY = hY - 1
			}

			hXY := strconv.Itoa(hX) + ":" + strconv.Itoa(hY)
			fmt.Println("  Head", hXY)

			lastX := hX
			lastY := hY

			for j := 0; j < knots; j++ {
				temp := strings.Split(knotPlacements[j], ":")
				sX, _ := strconv.Atoi(temp[0])
				sY, _ := strconv.Atoi(temp[1])

				fmt.Println("  Last:", lastX, lastY)
				if direction == "R" && lastX >= sX+2 {
					sX = lastX - 1
					sY = lastY
				} else if direction == "L" && lastX <= sX-2 {
					sX = lastX + 1
					sY = lastY
				} else if direction == "U" && lastY >= sY+2 {
					sY = lastY - 1
					sX = lastX
				} else if direction == "D" && lastY <= sY-2 {
					sY = lastY + 1
					sX = lastX
				}

				sXY := strconv.Itoa(sX) + ":" + strconv.Itoa(sY)
				knotPlacements[j] = sXY

				fmt.Println("    Knot", j, sXY)

				lastX = sX
				lastY = sY
			}

			// if !(contains(sXYPlacements, sXY)) {
			// 	sXYPlacements = append(sXYPlacements, sXY)
			// }
		}
	}

	return len(sXYPlacements)
}
