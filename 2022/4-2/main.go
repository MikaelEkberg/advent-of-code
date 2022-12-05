package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	count := 0

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		pairs := strings.Split(fileScanner.Text(), ",")
		firstSection := pairs[0]
		secondSection := pairs[1]

		firstNumberFirstSection, lastNumberFirstSection := numbers(firstSection)
		firstNumberSecondSection, lastNumberSecondSection := numbers(secondSection)

		if lastNumberFirstSection >= firstNumberSecondSection && firstNumberFirstSection <= lastNumberSecondSection {
			count++
		} else if lastNumberSecondSection >= firstNumberFirstSection && firstNumberSecondSection <= lastNumberFirstSection {
			count++
		}
	}
	fmt.Println(count)
}

func numbers(text string) (int, int) {
	numbers := strings.Split(text, "-")
	firstNumber, _ := strconv.Atoi(numbers[0])
	secondNumber, _ := strconv.Atoi(numbers[1])

	return firstNumber, secondNumber
}
