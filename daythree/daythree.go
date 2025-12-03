package daythree

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func ParseInput(path string) ([][]int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return [][]int{}, err
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")
	result := make([][]int, len(lines))
	for i, line := range lines {
		bank := make([]int, len(line))
		for j, r := range line {
			if r < 48 || r > 57 {
				return [][]int{}, fmt.Errorf("Invalid character: %c", r)
			}
			bank[j] = int(r - '0')
		}
		result[i] = bank
	}
	return result, nil
}

func Sum(input [][]int, lineJolt func(bank []int) int64) int64 {
	var sum int64 = 0
	for _, bank := range input {
		sum += lineJolt(bank)
	}
	return sum
}

func PartOneJolt(bank []int) int64 {
	maximum, maxIndex := 0, 0
	for i, jolt := range bank {
		if i == len(bank)-1 {
			continue
		}
		if jolt > maximum {
			maximum = jolt
			maxIndex = i
		}
	}
	nextMax := slices.Max(bank[maxIndex+1:])
	return int64((maximum * 10) + nextMax)
}

func PartTwoJolt(bank []int) int64 {
	var result int64 = 0
	lowLimit := 0
	for i := range 12 {
		highLimit := len(bank) - (11 - i)
		// fmt.Printf("i: %d, set: %v\n", i, bank[lowLimit:highLimit])
		maximum, maxIndex := 0, 0
		for j, jolt := range bank[lowLimit:highLimit] {
			if jolt > maximum {
				maximum = jolt
				maxIndex = j
			}
		}
		// fmt.Printf("max: %d, maxIndex: %d\n", maximum, maxIndex)
		result += int64(maximum * int(math.Pow10(11-i)))
		lowLimit = lowLimit + maxIndex + 1
		// fmt.Println()
	}
	return result
}
