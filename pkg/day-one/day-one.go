package dayone

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func DayOne() {
	// put each list into a sorted slice
	// compare the distance between
	fileName := "part-one.txt"

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(path.Join(cwd, "pkg/day-one", fileName))
	// full stop if the file errors
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	partTwo(scanner)

}

func partOne(scanner *bufio.Scanner) {
	leftSlice := make([]int, 0)
	rightSlice := make([]int, 0)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		lNum, err := strconv.Atoi(input[0])
		if err != nil {
			panic(err)
		}
		rNum, err := strconv.Atoi(input[len(input)-1])
		if err != nil {
			panic(err)
		}

		leftSlice = append(leftSlice, lNum)
		rightSlice = append(rightSlice, rNum)
	}

	sort.Ints(leftSlice)
	sort.Ints(rightSlice)

	total := 0
	for idx, leftVal := range leftSlice {
		distance := math.Abs(float64(leftVal) - float64(rightSlice[idx]))

		total += int(distance)
	}

	fmt.Printf("Total distance is %d\n", total)
}

func partTwo(scanner *bufio.Scanner) {
	leftSlice := make([]int, 0)
	rightSlice := make([]int, 0)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		lNum, err := strconv.Atoi(input[0])
		if err != nil {
			panic(err)
		}
		rNum, err := strconv.Atoi(input[len(input)-1])
		if err != nil {
			panic(err)
		}

		leftSlice = append(leftSlice, lNum)
		rightSlice = append(rightSlice, rNum)
	}

	numMap := make(map[int]int)
	for _, leftVal := range leftSlice {
		numMap[leftVal] = 0
	}

	for _, rightVal := range rightSlice {
		if value, ok := numMap[rightVal]; ok {
			numMap[rightVal] = value + 1
		}
	}

	total := 0
	for _, l := range leftSlice {
		distance := l * numMap[l]
		total += int(distance)
	}

	fmt.Printf("Total distance is %d\n", total)
}
