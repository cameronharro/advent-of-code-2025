package dayfive_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/dayfive"
)

func TestParseInput(t *testing.T) {
	result, err := dayfive.ParseInput("./dayfive_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	expected := dayfive.Input{
		FreshRanges: [][]int64{
			{3, 5},
			{10, 14},
			{16, 20},
			{12, 18},
		},
		Ids: []int64{1, 5, 8, 11, 17, 32},
	}

	if !slices.EqualFunc(result.FreshRanges, expected.FreshRanges, func(E1, E2 []int64) bool {
		return slices.Equal(E1, E2)
	}) {
		t.Errorf("ParseInput() FreshRanges got %v, expected %v", result.FreshRanges, expected.FreshRanges)
	}
	if !slices.Equal(result.Ids, expected.Ids) {
		t.Errorf("ParseInput() Ids got %v, expected %v", result.Ids, expected.Ids)
	}
}

func TestPartOne(t *testing.T) {
	input := dayfive.Input{
		FreshRanges: [][]int64{
			{3, 5},
			{10, 14},
			{16, 20},
			{12, 18},
		},
		Ids: []int64{1, 5, 8, 11, 17, 32},
	}

	result := dayfive.PartOne(input)
	if result != 3 {
		t.Errorf("PartOne() got %d, expected 3", result)
	}

	input, err := dayfive.ParseInput("./dayfive.txt")
	if err != nil {
		t.Error(err.Error())
	}

	result = dayfive.PartOne(input)
	fmt.Println()
	fmt.Printf("Part One result: %d\n", result)
	fmt.Println()
}
