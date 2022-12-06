package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	day := 6
	input := getInput("input.txt")

	firstPartSolution, secondPartSolution := getCalories(input)

	printSolution(firstPartSolution, secondPartSolution, day)
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

func getCalories(input string) (int, int) {
	values := strings.Split(input, "\n\n")

	var slice []int

	for _, v := range values {
		j := strings.Fields(v)

		result := 0

		for _, t := range j {
			test, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			result += test
		}

		slice = append(slice, result)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(slice)))

	sum := 0
	for _, k := range slice[0:3] {
		sum += k
	}

	return slice[0], sum
}
