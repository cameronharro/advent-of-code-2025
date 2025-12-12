package daysix_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/daysix"
)

func TestParseInputOne(t *testing.T) {
	result, err := daysix.ParseInputOne("./daysix_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	expected := []daysix.Problem{
		{[]int{123, 45, 6}, "*"},
		{[]int{328, 64, 98}, "+"},
		{[]int{51, 387, 215}, "*"},
		{[]int{64, 23, 314}, "+"},
	}

	if !slices.EqualFunc(result, expected, func(E1, E2 daysix.Problem) bool {
		return slices.Equal(E1.Numbers, E2.Numbers) && E1.Operator == E2.Operator
	}) {
		t.Errorf("ParseInputOne() got %v, expected %v", result, expected)
	}
}

func TestParseInputTwo(t *testing.T) {
	result, err := daysix.ParseInputTwo("./daysix_test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	expected := []daysix.Problem{
		{[]int{1, 24, 356}, "*"},
		{[]int{369, 248, 8}, "+"},
		{[]int{32, 581, 175}, "*"},
		{[]int{623, 431, 4}, "+"},
	}

	if !slices.EqualFunc(result, expected, func(E1, E2 daysix.Problem) bool {
		return slices.Equal(E1.Numbers, E2.Numbers) && E1.Operator == E2.Operator
	}) {
		t.Errorf("ParseInputTwo() got %v, expected %v", result, expected)
	}
}

func TestSolveProblem(t *testing.T) {
	type TestCase struct {
		row      daysix.Problem
		expected int
	}
	testCases := []TestCase{
		{
			row:      daysix.Problem{[]int{123, 45, 6}, "*"},
			expected: 33210,
		},
		{
			row:      daysix.Problem{[]int{328, 64, 98}, "+"},
			expected: 490,
		},
		{
			row:      daysix.Problem{[]int{51, 387, 215}, "*"},
			expected: 4243455,
		},
		{
			row:      daysix.Problem{[]int{64, 23, 314}, "+"},
			expected: 401,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := daysix.SolveProblem(testCase.row)
			if result != testCase.expected {
				t.Errorf("SolveProblem() got %d, expected %d", result, testCase.expected)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	input, err := daysix.ParseInputOne("./daysix_test.txt")
	if err != nil {
		t.Error(err.Error())
	}
	result := daysix.EvalProblems(input)
	expected := 4277556
	if result != expected {
		t.Errorf("PartOne() got %d, expected %d", result, expected)
	}

	input, err = daysix.ParseInputOne("./daysix.txt")
	if err != nil {
		t.Error(err.Error())
	}
	result = daysix.EvalProblems(input)

	fmt.Println()
	fmt.Printf("Part One Result: %d\n", result)
	fmt.Println()

}

func TestPartTwo(t *testing.T) {
	input, err := daysix.ParseInputTwo("./daysix_test.txt")
	if err != nil {
		t.Error(err.Error())
	}
	result := daysix.EvalProblems(input)
	expected := 3263827
	if result != expected {
		t.Errorf("PartTwo() got %d, expected %d", result, expected)
	}

	input, err = daysix.ParseInputTwo("./daysix.txt")
	if err != nil {
		t.Error(err.Error())
	}
	result = daysix.EvalProblems(input)

	fmt.Println()
	fmt.Printf("Part Two Result: %d\n", result)
	fmt.Println()

}
