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

	highest := 0

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

		if result > highest {
			highest = result
		}
	}
	fmt.Println(highest)
}
