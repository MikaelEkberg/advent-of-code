package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	day := 11
	input := getInput("input.txt")

	solutionPartOne := solutionPartOne(input)
	solutionPartTwo := solutionPartTwo(input)

	printSolution(solutionPartOne, solutionPartTwo, day)
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
	monkeys := strings.Split(input, "\n\n")
	var monkeyMap = make(map[int]Monkey)
	var totalCount []int

	for i, v := range monkeys {
		lines := strings.Split(v, "\n")

		name := i
		items := strings.Split(strings.Split(lines[1], ":")[1], ",")

		itemsInt := deleteEmpty(items)

		operation := strings.Split(strings.Split(lines[2], ":")[1], "=")[1]
		test := strings.Split(lines[3], " ")[5]
		trueMonkey, _ := strconv.Atoi(strings.Split(lines[4], " ")[9])
		falseMonkey, _ := strconv.Atoi(strings.Split(lines[5], " ")[9])

		monkeyMap[i] = Monkey{name, itemsInt, operation, test, trueMonkey, falseMonkey, 0}
	}

	for j := 0; j < 20; j++ {

		for i := range monkeys {
			operation := strings.Split(monkeyMap[i].operation, " ")[2]

			for _, value := range monkeyMap[i].items {
				var increase, test, nextMonkey int

				if strings.Split(monkeyMap[i].operation, " ")[3] == "old" {
					increase = value
				} else {
					temp := strings.Split(monkeyMap[i].operation, " ")[3]
					increase, _ = strconv.Atoi(temp)
				}

				switch operation {
				case "*":
					test = value * increase
				case "+":
					test = value + increase
				}

				res := test / 3

				divisionValue, _ := strconv.Atoi(monkeyMap[i].test)

				if res%divisionValue == 0 {
					nextMonkey = monkeyMap[i].trueMonkey
				} else {
					nextMonkey = monkeyMap[i].falseMonkey
				}

				tmpNext := monkeyMap[nextMonkey]
				tmpNext.items = append(tmpNext.items, res)
				monkeyMap[nextMonkey] = tmpNext

			}

			var emptySlice []int
			tmpCurrent := monkeyMap[i]
			tmpCurrent.count += len(tmpCurrent.items)
			tmpCurrent.items = emptySlice
			monkeyMap[i] = tmpCurrent
		}
	}

	for i := range monkeys {
		totalCount = append(totalCount, monkeyMap[i].count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totalCount)))

	res := totalCount[0] * totalCount[1]

	return res
}

func solutionPartTwo(input string) int {
	monkeys := strings.Split(input, "\n\n")
	var monkeyMap = make(map[int]Monkey)
	var totalCount []int

	for i, v := range monkeys {
		lines := strings.Split(v, "\n")

		name := i
		items := strings.Split(strings.Split(lines[1], ":")[1], ",")

		itemsInt := deleteEmpty(items)

		operation := strings.Split(strings.Split(lines[2], ":")[1], "=")[1]
		test := strings.Split(lines[3], " ")[5]
		trueMonkey, _ := strconv.Atoi(strings.Split(lines[4], " ")[9])
		falseMonkey, _ := strconv.Atoi(strings.Split(lines[5], " ")[9])

		monkeyMap[i] = Monkey{name, itemsInt, operation, test, trueMonkey, falseMonkey, 0}
	}

	divisor := 1

	for i := range monkeys {
		divisionValue, _ := strconv.Atoi(monkeyMap[i].test)
		divisor *= divisionValue
	}

	for j := 0; j < 10000; j++ {

		for i := range monkeys {
			operation := strings.Split(monkeyMap[i].operation, " ")[2]

			for _, value := range monkeyMap[i].items {
				var increase, test, nextMonkey int

				if strings.Split(monkeyMap[i].operation, " ")[3] == "old" {
					increase = value
				} else {
					temp := strings.Split(monkeyMap[i].operation, " ")[3]
					increase, _ = strconv.Atoi(temp)
				}

				switch operation {
				case "*":
					test = value * increase
				case "+":
					test = value + increase
				}

				divisionValue, _ := strconv.Atoi(monkeyMap[i].test)

				res := test % divisor

				if res%divisionValue == 0 {
					nextMonkey = monkeyMap[i].trueMonkey
				} else {
					nextMonkey = monkeyMap[i].falseMonkey
				}

				tmpNext := monkeyMap[nextMonkey]
				tmpNext.items = append(tmpNext.items, res)
				monkeyMap[nextMonkey] = tmpNext
			}

			var emptySlice []int
			tmpCurrent := monkeyMap[i]
			tmpCurrent.count += len(tmpCurrent.items)
			tmpCurrent.items = emptySlice
			monkeyMap[i] = tmpCurrent
		}
	}

	fmt.Println(monkeyMap[0].count)
	fmt.Println(monkeyMap[1].count)
	fmt.Println(monkeyMap[2].count)
	fmt.Println(monkeyMap[3].count)

	for i := range monkeys {
		totalCount = append(totalCount, monkeyMap[i].count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totalCount)))

	res := totalCount[0] * totalCount[1]

	return res
}

func deleteEmpty(s []string) []int {
	var r []int
	for _, str := range s {
		if strings.TrimSpace(str) != "" {
			str = strings.TrimSpace(str)
			v, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}

			r = append(r, v)
		}
	}

	return r
}

type Monkey struct {
	name        int
	items       []int
	operation   string
	test        string
	trueMonkey  int
	falseMonkey int
	count       int
}
