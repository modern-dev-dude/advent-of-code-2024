package problems

import (
	"aoc/pkg/utils"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

func DayTwo() {
	// put each list into a sorted slice
	// compare the distance between
	fileName := "test-data.txt"
	relativePath := "pkg/problems/day-two"

	scanner, err := utils.ScanFile(fileName, relativePath)
	if err != nil {
		panic(err)
	}

	defer scanner.CloseFile()

	partTwo(scanner)
}

func partOne(scanner *utils.File) {
	numOfSafeReports := 0

	for {
		line, err := scanner.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		listOfChars := strings.Split(line, " ")

		prevVal := 0
		desc := false

		for idx, char := range listOfChars {
			val := convertCharToInt(char)
			if idx == 0 {
				continue
			}

			prevVal = convertCharToInt(listOfChars[idx-1])

			if idx == 1 {
				desc = val < prevVal
			}

			if !desc && val < prevVal {
				break
			}

			if desc && val > prevVal {
				break
			}

			valDiff := math.Abs(float64(val - prevVal))

			if valDiff > 3 || valDiff <= 0 {
				break
			}

			if idx != len(listOfChars)-1 {
				continue
			}

			numOfSafeReports += 1

		}

	}

	fmt.Printf("Safe reports: %v\n", numOfSafeReports)
}

func partTwo(scanner *utils.File) {
	safeReport := 0

	for {
		line, err := scanner.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		hasRemovedOneChar := false
		desc := false
		listOfLines := strings.Split(line, " ")
		prevVal := convertCharToInt(listOfLines[0])

		for idx, char := range listOfLines {
			val := convertCharToInt(char)
			if idx == 0 {
				continue
			}
			if idx == 1 {
				desc = val > prevVal
			}

			if !desc && val > prevVal {
				if hasRemovedOneChar {
					break
				}
				hasRemovedOneChar = true
				continue
			}

			if desc && val < prevVal {
				if hasRemovedOneChar {
					break
				}
				hasRemovedOneChar = true
				continue
			}

			diff := math.Abs(float64(prevVal - val))
			fmt.Printf("prev, val %d - %d\n", prevVal, val)
			if diff > 3 || diff <= 0 {
				if hasRemovedOneChar {
					break
				}

				hasRemovedOneChar = true
				continue
			}

			if idx != len(listOfLines)-1 {
				prevVal = val
				continue
			}

			safeReport += 1
		}
		fmt.Println("================")

	}

	fmt.Printf("Safe reports : %d\n", safeReport)

}

func convertCharToInt(char string) int {
	val, err := strconv.Atoi(char)
	if err != nil {
		panic(err)
	}

	return val
}
