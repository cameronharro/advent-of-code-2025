package dayfive_test

import (
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
