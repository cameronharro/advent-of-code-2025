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
