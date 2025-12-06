package daysix_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/daysix"
)

func TestParseInput(t *testing.T) {
	result, err := daysix.ParseInput("./daysix_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	expected := [][]string{
		{"123", "45", "6", "*"},
		{"328", "64", "98", "+"},
		{"51", "387", "215", "*"},
		{"64", "23", "314", "+"},
	}

	if !slices.EqualFunc(result, expected, func(E1, E2 []string) bool {
		return slices.Equal(E1, E2)
	}) {
		t.Errorf("ParseInput() got %v, expected %v", result, expected)
	}
}

func TestPartOne(t *testing.T) {
	type TestCase struct {
		row      []string
		expected int
	}
	testCases := []TestCase{
		{
			row:      []string{"123", "45", "6", "*"},
			expected: 33210,
		},
		{
			row:      []string{"328", "64", "98", "+"},
			expected: 490,
		},
		{
			row:      []string{"51", "387", "215", "*"},
			expected: 4243455,
		},
		{
			row:      []string{"64", "23", "314", "+"},
			expected: 401,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := daysix.PartOne(testCase.row)
			if result != testCase.expected {
				t.Errorf("PartOne() got %d, expected %d", result, testCase.expected)
			}
		})
	}

	input, err := daysix.ParseInput("./daysix.txt")
	if err != nil {
		t.Error(err.Error())
	}
	partOne := daysix.Part(daysix.PartOne)
	result := partOne(input)

	fmt.Println()
	fmt.Printf("Part One Result: %d\n", result)
	fmt.Println()

}
