package daythree

import (
	"fmt"
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

func Sum(input [][]int, lineJolt func(bank []int) int) int {
	sum := 0
	for _, bank := range input {
		sum += lineJolt(bank)
	}
	return sum
}

func PartOneJolt(bank []int) int {
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
	return (maximum * 10) + nextMax
}
