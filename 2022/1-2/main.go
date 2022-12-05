package main

import (
	"fmt"
	"os"
	"sort"
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
	fmt.Println("1. Most calories:", slice[0])
	fmt.Println("2. Top three calories", sum)
}
