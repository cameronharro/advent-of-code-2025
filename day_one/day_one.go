package dayone

import (
	"os"
	"strconv"
	"strings"
)

func PartOne(nums []int) int {
	result := 0
	displacement := 0
	for _, n := range nums {
		displacement += n
		if displacement%100 == 50 || displacement%100 == -50 {
			result++
		}
	}
	return result
}

func ParseInput(path string) ([]int, error) {
	input, err := os.ReadFile(path)
	if err != nil {
		return []int{}, err
	}

	lines := strings.Split(string(input), "\n")
	result := make([]int, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		if len(line) == 0 {
			continue
		}
		direction, distance := line[:1], line[1:]
		n, err := strconv.Atoi(distance)
		if err != nil {
			return []int{}, err
		}
		if direction == "L" {
			n = n * -1
		}
		result[i] = n
	}
	return result, nil
}

func PartTwo(nums []int) int {
	result := 0
	position := 50

	for _, n := range nums {
		crossed := 0
		startingPosition := position
		position += n
		if position < 0 {
			crossed += 1 - (position / 100)
			if position%100 == 0 {
				crossed--
			}
			position += crossed * 100
			if startingPosition == 0 {
				crossed--
			}
		} else if position >= 100 {
			crossed += position / 100
			position -= crossed * 100
			if position == 0 {
				crossed--
			}
		}
		result += crossed
		if position == 0 {
			result++
		}
	}
	return result
}
