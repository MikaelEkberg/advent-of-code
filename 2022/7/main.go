package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	day := 7
	input := getInput("input.txt")
	lines := strings.Split(input, "\n")
	var slice []int
	var directories []string
	var currentDir string
	var directorySizes = make(map[string]int)
	var directorySizesRecursive = make(map[string]int)

	for i, v := range lines {
		if strings.HasPrefix(v, "$ cd") {
			slice = append(slice, i)
		}
	}

	slice = append(slice, len(lines))

	sliceLength := len(slice)

	for i := 0; i < sliceLength-1; i++ {
		start := slice[i]
		end := slice[i+1]

		section := lines[start+1 : end]

		dir := strings.Split(lines[start], " ")[2]

		if dir == ".." {
			dirs := strings.Split(currentDir, "/")
			dirCount := len(dirs)
			if dirCount <= 2 {
				currentDir = "/"
			} else {
				currentDir = strings.Join(dirs[0:dirCount-1], "/")
			}
		} else {
			if currentDir == "/" {
				currentDir += dir
			} else if len(currentDir) == 0 {
				currentDir += dir
			} else {
				currentDir += "/" + dir
			}
		}
		currentDirSize := 0

		if !(contains(directories, currentDir)) {
			directories = append(directories, currentDir)
		}

		if len(section) > 0 && strings.HasSuffix(section[0], "ls") {
			for _, r := range section {
				if !(strings.HasPrefix(r, "dir")) && !(strings.HasPrefix(r, "$")) {
					fileSize := strings.Split(r, " ")[0]
					fileSizeInt, _ := strconv.Atoi(fileSize)
					currentDirSize += fileSizeInt
				}
			}
			directorySizes[currentDir] = currentDirSize
		}
	}

	for _, v := range directories {
		totalSize := 0
		for k := range directorySizes {
			if strings.HasPrefix(k, v) {
				totalSize += directorySizes[k]
			}
		}
		directorySizesRecursive[v] = totalSize
	}

	solutionPartOne := getSolutionPartOne(directorySizesRecursive)
	solutionPartTwo := getSolutionPartTwo(70000000, 30000000, directorySizesRecursive)

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

func getSolutionPartOne(directorySizesRecursive map[string]int) int {
	sum := 0

	for _, v := range directorySizesRecursive {
		if v <= 100000 {
			sum += v
		}
	}

	return sum
}

func getSolutionPartTwo(diskSize int, diskSpaceSpaceNeeded int, directorySizesRecursive map[string]int) int {
	diskUsed := directorySizesRecursive["/"]
	currentlyFree := diskSize - diskUsed
	diff := diskSpaceSpaceNeeded - currentlyFree
	var directoriesLargerThanDiff []int

	for _, v := range directorySizesRecursive {
		if v >= diff {
			directoriesLargerThanDiff = append(directoriesLargerThanDiff, v)
		}
	}

	sort.Ints(directoriesLargerThanDiff)

	smallestDirLargerThanDiff := directoriesLargerThanDiff[0]

	return smallestDirLargerThanDiff
}
