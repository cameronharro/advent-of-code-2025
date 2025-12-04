package dayfour_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/dayfour"
)

func TestParseInput(t *testing.T) {
	result, err := dayfour.ParseInput("./dayfour_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	expectedLength := 10
	if len(result) != expectedLength || len(result[0]) != expectedLength {
		t.Errorf("ParseInput() got %dx%d grid, expected %dx%d\n", len(result), len(result[0]), expectedLength, expectedLength)
	}
}

func TestGetSurrounding(t *testing.T) {
	grid, err := dayfour.ParseInput("./dayfour_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	type TestCase struct {
		inputPoint dayfour.Point
		expected   []string
	}

	testCases := []TestCase{
		{
			inputPoint: dayfour.Point{
				X: 2,
				Y: 0,
			},
			expected: []string{".", "@", "@", "@", "."},
		},
		{
			inputPoint: dayfour.Point{
				X: 4,
				Y: 2,
			},
			expected: []string{".", "@", "@", "@", "@", ".", ".", "@"},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := dayfour.GetSurrounding(testCase.inputPoint, grid)
			if !slices.Equal(result, testCase.expected) {
				t.Errorf("GetSurrounding() got %v, expected %v", result, testCase.expected)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	grid, err := dayfour.ParseInput("./dayfour_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	result, _ := dayfour.PartOne(grid)
	if result != 13 {
		t.Errorf("PartOne() got %d, expected %d", result, 13)
	}

	grid, err = dayfour.ParseInput("./dayfour.txt")
	if err != nil {
		t.Error(err.Error())
	}

	result, _ = dayfour.PartOne(grid)
	fmt.Println()
	fmt.Printf("Part One result: %d\n", result)
	fmt.Println()
}

func TestPartTwo(t *testing.T) {
	grid, err := dayfour.ParseInput("./dayfour_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	result := dayfour.PartTwo(grid)
	if result != 43 {
		t.Errorf("PartTwo() got %d, expected %d", result, 43)
	}

	grid, err = dayfour.ParseInput("./dayfour.txt")
	if err != nil {
		t.Error(err.Error())
	}

	result = dayfour.PartTwo(grid)
	fmt.Println()
	fmt.Printf("Part Two result: %d\n", result)
	fmt.Println()
}
