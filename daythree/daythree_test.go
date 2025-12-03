package daythree_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/daythree"
)

func TestParseInput(t *testing.T) {
	type TestCase struct {
		input    string
		expected [][]int
	}

	testCases := []TestCase{
		{
			input: "./daythree_test.txt",
			expected: [][]int{
				{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
				{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
				{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
				{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result, err := daythree.ParseInput(testCase.input)
			if err != nil {
				t.Error(err.Error())
			}
			if !slices.EqualFunc(result, testCase.expected, func(E1, E2 []int) bool {
				return slices.Equal(E1, E2)
			}) {
				t.Errorf("ParseData() got %v, expected %v", result, testCase.expected)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	type TestCase struct {
		input    []int
		expected int64
	}
	testCases := []TestCase{
		{
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: 98,
		},
		{
			input:    []int{8, 1, 1, 1, 1, 9},
			expected: 89,
		},
		{
			input:    []int{2, 3, 4, 3, 2, 2, 4, 7, 8},
			expected: 78,
		},
		{
			input:    []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			expected: 92,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := daythree.PartOneJolt(testCase.input)
			if result != testCase.expected {
				t.Errorf("PartOneJolt() got %d, expected %d", result, testCase.expected)
			}
		})
	}

	input, err := daythree.ParseInput("./daythree.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println()
	fmt.Printf("Part One result: %d\n", daythree.Sum(input, daythree.PartOneJolt))
	fmt.Println()
}

func TestPartTwo(t *testing.T) {
	type TestCase struct {
		input    []int
		expected int64
	}
	testCases := []TestCase{
		{
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			expected: 987654321111,
		},
		{
			input:    []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			expected: 811111111119,
		},
		{
			input:    []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			expected: 434234234278,
		},
		{
			input:    []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			expected: 888911112111,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := daythree.PartTwoJolt(testCase.input)
			if result != testCase.expected {
				t.Errorf("PartOneJolt() got %d, expected %d", result, testCase.expected)
			}
		})
	}

	input, err := daythree.ParseInput("./daythree.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println()
	fmt.Printf("Part Two result: %d\n", daythree.Sum(input, daythree.PartTwoJolt))
	fmt.Println()
}
