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

func TestPartTwo(t *testing.T) {
	testData := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	result := dayone.PartTwo(testData)
	if result != 6 {
		t.Errorf("PartTwo() got %d, expected 6", result)
	}
	testData = []int{-1050}
	result = dayone.PartTwo(testData)
	if result != 11 {
		t.Errorf("PartTwo() got %d, expected 11", result)
	}
	testData = []int{1050, -1}
	result = dayone.PartTwo(testData)
	if result != 11 {
		t.Errorf("PartTwo() got %d, expected 11", result)
	}

	data, err := dayone.ParseInput("./day_one.txt")
	if err != nil {
		t.Error(err.Error())
	}
	result = dayone.PartTwo(data)
	fmt.Println(result)
}
