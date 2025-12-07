package dayseven_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/dayseven"
)

func TestParseInput(t *testing.T) {
	grid, err := dayseven.ParseInput("./dayseven_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	expected := [][]string{
		{".", ".", ".", ".", ".", ".", ".", "S", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "^", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "^", ".", "^", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "^", ".", "^", ".", "^", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "^", ".", "^", ".", ".", ".", "^", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", "^", ".", "^", ".", ".", ".", "^", ".", "^", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "^", ".", ".", ".", "^", ".", ".", ".", ".", ".", "^", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "^", ".", "^", ".", "^", ".", "^", ".", "^", ".", ".", ".", "^", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	if !slices.EqualFunc(grid, expected, func(E1, E2 []string) bool {
		return slices.Equal(E1, E2)
	}) {
		t.Errorf("ParseInput() got %v, expected %v", grid, expected)
	}
}

func TestPartOne(t *testing.T) {
	grid, err := dayseven.ParseInput("./dayseven_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	result := dayseven.PartOne(grid)
	expected := 21
	if result != expected {
		t.Errorf("PartOne got %d, expected %d", result, expected)
	}

	grid, err = dayseven.ParseInput("./dayseven.txt")
	if err != nil {
		t.Error(err.Error())
	}

	result = dayseven.PartOne(grid)
	fmt.Println()
	fmt.Printf("Part One answer: %d\n", result)
	fmt.Println()
}

func TestPartTwo(t *testing.T) {
	grid, err := dayseven.ParseInput("./dayseven_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	result := dayseven.PartTwo(grid)
	expected := 40
	if result != expected {
		t.Errorf("PartTwo got %d, expected %d", result, expected)
	}

	grid, err = dayseven.ParseInput("./dayseven.txt")
	if err != nil {
		t.Error(err.Error())
	}

	result = dayseven.PartTwo(grid)
	fmt.Println()
	fmt.Printf("Part Two answer: %d\n", result)
	fmt.Println()
}
