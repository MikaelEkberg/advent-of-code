package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	day := 8
	input := getInput("input.txt")
	lines := strings.Split(input, "\n")
	lineMaxIndex := len(lines) - 1
	var visibleCount int
	var scenicScore int

	for i, line := range lines {
		lineLength := len(line)
		if i == 0 || i == lineMaxIndex {
			visibleCount += lineLength
		} else {
			chars := []rune(line)
			lastIndex := lineLength - 1

			for j, v := range chars {
				if j == 0 || j == lastIndex {
					visibleCount++
				} else {
					var left []int
					var right []int
					var up []int
					var down []int
					valueInt := int(v - '0')
					visibleLeft, visibleRight, visibleUp, visibleDown := true, true, true, true
					var leftViewCount, rightViewCount, upViewCount, downViewCount int

					for _, k := range chars[0:j] {
						left = append(left, int(k-'0'))
					}

					left = reverseInts(left)

					for _, k := range chars[j+1 : lineLength] {
						right = append(right, int(k-'0'))
					}

					for k := 0; k < i; k++ {
						lineChars := []rune(lines[k])
						up = append(up, int(lineChars[j]-'0'))
					}

					up = reverseInts(up)

					for k := i + 1; k < lineLength; k++ {
						lineChars := []rune(lines[k])
						down = append(down, int(lineChars[j]-'0'))
					}

					for _, k := range left {
						if k >= valueInt {
							visibleLeft = false
						}
					}

					for _, k := range left {
						if k < valueInt {
							leftViewCount++
						} else {
							leftViewCount++
							break
						}
					}

					for _, k := range right {
						if k >= valueInt {
							visibleRight = false
						}
					}

					for _, k := range right {
						if k < valueInt {
							rightViewCount++
						} else {
							rightViewCount++
							break
						}
					}

					for _, k := range up {
						if k >= valueInt {
							visibleUp = false
						}
					}

					for _, k := range up {
						if k < valueInt {
							upViewCount++
						} else {
							upViewCount++
							break
						}
					}

					for _, k := range down {
						if k >= valueInt {
							visibleDown = false
						}
					}

					for _, k := range down {
						if k < valueInt {
							downViewCount++
						} else {
							downViewCount++
							break
						}
					}

					if visibleLeft || visibleRight || visibleUp || visibleDown {
						visibleCount++
					}

					result := upViewCount * leftViewCount * rightViewCount * downViewCount

					if result > scenicScore {
						scenicScore = result
					}
				}
			}
		}
	}

	printSolution(visibleCount, scenicScore, day)
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

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
