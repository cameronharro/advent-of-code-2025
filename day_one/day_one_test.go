package dayone_test

import (
	"fmt"
	"slices"
	"testing"

	dayone "github.com/cameronharro/advent-of-code-2025/day_one"
)

func TestParseInput(t *testing.T) {
	numbers, err := dayone.ParseInput("./day_one_test.txt")
	if err != nil {
		t.Error(err.Error())
	}
	expected := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	if !slices.Equal(numbers, expected) {
		t.Errorf("ParseInput() got %v expected %v", numbers, expected)
	}
}

func TestPartOne(t *testing.T) {
	testData := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	result := dayone.PartOne(testData)
	if result != 3 {
		t.Errorf("PartOne() got %d, expected 3", result)
	}

	data, err := dayone.ParseInput("./day_one.txt")
	if err != nil {
		t.Error(err.Error())
	}
	result = dayone.PartOne(data)
	fmt.Println(result)
}
