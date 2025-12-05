package dayfive

import (
	"fmt"
	"os"
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
