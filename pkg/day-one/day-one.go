package dayone

import (
	"aoc/pkg/utils"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

func DayOne() {
	// put each list into a sorted slice
	// compare the distance between
	fileName := "problem-data.txt"

	scanner, err := utils.ScanFile(fileName)
	if err != nil {
		panic(err)
	}

	defer scanner.CloseFile()

	partTwo(scanner)

}

func partOne(scanner *utils.File) {
	leftSlice := make([]int, 0)
	rightSlice := make([]int, 0)

	for {
		line, err := scanner.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		input := strings.Split(line, " ")
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

func partTwo(scanner *utils.File) {
	leftSlice := make([]int, 0)
	rightSlice := make([]int, 0)

	for {
		line, err := scanner.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		input := strings.Split(line, " ")
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
