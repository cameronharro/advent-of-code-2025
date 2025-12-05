package dayfive

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Input struct {
	FreshRanges [][]int64
	Ids         []int64
}

func ParseInput(path string) (Input, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Input{}, err
	}

	result := Input{}
	parts := strings.Split(string(data), "\n\n")
	if len(parts) != 2 {
		return Input{}, fmt.Errorf("ParseInput() got %d parts, expected 2", len(parts))
	}

	for rangeString := range strings.SplitSeq(parts[0], "\n") {
		if rangeString == "" {
			continue
		}
		freshRange := make([]int64, 2)
		bounds := strings.Split(rangeString, "-")
		for i, boundStr := range bounds {
			bound, err := strconv.ParseInt(boundStr, 10, 64)
			if err != nil {
				return Input{}, err
			}
			freshRange[i] = bound
		}
		result.FreshRanges = append(result.FreshRanges, freshRange)
	}
	slices.SortFunc(result.FreshRanges, func(a, b []int64) int {
		diff := a[0] - b[0]
		if diff > 0 {
			return 1
		} else if diff < 0 {
			return -1
		}
		return 0
	})
	result.FreshRanges = getDiscontinuousRanges(result.FreshRanges)

	for idString := range strings.SplitSeq(parts[1], "\n") {
		if idString == "" {
			continue
		}
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			return Input{}, err
		}
		result.Ids = append(result.Ids, id)
	}
	return result, nil
}

func getDiscontinuousRanges(sortedRanges [][]int64) [][]int64 {
	result := [][]int64{}
	foundRanges := 0
	for _, newRange := range sortedRanges {
		if len(result) == 0 {
			foundRanges++
			result = append(result, newRange)
			continue
		}
		currentRange := result[foundRanges-1]
		if newRange[1] < currentRange[1] {
			continue
		}
		if newRange[0] <= currentRange[1] {
			currentRange[1] = newRange[1]
			continue
		}
		foundRanges++
		result = append(result, newRange)
	}
	return result
}

func PartOne(input Input) int {
	result := 0
	for _, id := range input.Ids {
		for _, freshRange := range input.FreshRanges {
			if id < freshRange[0] || id > freshRange[1] {
				continue
			}
			result++
			break
		}
	}
	return result
}

func PartTwo(input Input) int64 {
	var result int64 = 0
	for _, r := range input.FreshRanges {
		result += r[1] - r[0] + 1
	}
	return result
}
