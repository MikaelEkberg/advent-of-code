package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Err")
	}

	text := string(content)
	values := strings.Split(text, "\n\n")
	stack := values[0]
	rows := strings.Split(stack, "\n")
	stackNumbers := rows[len(rows)-1]
	stackNumbersArrary := strings.Fields(stackNumbers)
	moves := strings.Split(values[1], "\n")
	start := 1
	lastRowNumber := len(rows)

	var stackMap = make(map[string][]string)

	for _, v := range stackNumbersArrary {
		arrayTest := []string{}
		for i, j := range rows {
			if i < lastRowNumber-1 {
				character := string(j[start])
				if strings.TrimSpace(character) != "" {
					arrayTest = append(arrayTest, character)
				}
			}
		}
		stackMap[v] = arrayTest
		start = start + 4
	}

	for _, v := range moves {
		count, from, to := rearrangement(v)

		for i := 0; i < count; i++ {
			characterToMove := stackMap[from][0]
			// fmt.Println("Move ", characterToMove, " to ", stackMap[to])
			stackMap[to] = append([]string{characterToMove}, stackMap[to]...)
			temp := stackMap[from]
			updatedIndex := RemoveIndex(temp, 0)
			stackMap[from] = updatedIndex
		}
	}

	result := ""

	for i := 1; i <= len(stackMap); i++ {
		index := strconv.Itoa(i)
		result += stackMap[index][0]
	}

	fmt.Println(result)
}

func rearrangement(text string) (int, string, string) {
	numbers := strings.Split(text, " ")
	count, _ := strconv.Atoi(numbers[1])
	from := numbers[3]
	to := numbers[5]

	return count, from, to
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
