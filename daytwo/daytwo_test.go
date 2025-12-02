package daytwo_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/daytwo"
)

func TestParseData(t *testing.T) {
	input, err := daytwo.ParseInput("./daytwo_test.txt")
	if err != nil {
		t.Error(err.Error())
	}
	expected := [][]int{
		{11, 22},
		{95, 115},
		{998, 1012},
		{1188511880, 1188511890},
		{222220, 222224},
		{1698522, 1698528},
		{446443, 446449},
		{38593856, 38593862},
		{565653, 565659},
		{824824821, 824824827},
		{2121212118, 2121212124},
	}
	if !slices.EqualFunc(input, expected, func(E1, E2 []int) bool {
		if !slices.Equal(E1, E2) {
			return false
		}
		return true
	}) {
		t.Errorf("ParseData() got %v, expected %v", input, expected)
	}
}

func TestPartOne(t *testing.T) {
	partOne := daytwo.Part(daytwo.PartOneValidCheck)
	type TestCase struct {
		input   [][]int
		result  int
		wantErr bool
	}
	testCases := []TestCase{
		{
			input:   [][]int{{11}},
			result:  0,
			wantErr: true,
		},
		{
			input:   [][]int{{12, 11}},
			result:  0,
			wantErr: true,
		},
		{
			input:   [][]int{{11, 12}},
			result:  11,
			wantErr: false,
		},
		{
			input:   [][]int{{11, 22}},
			result:  33,
			wantErr: false,
		},
		{
			input: [][]int{
				{11, 22},
				{95, 115},
				{998, 1012},
				{1188511880, 1188511890},
				{222220, 222224},
				{1698522, 1698528},
				{446443, 446449},
				{38593856, 38593862},
				{565653, 565659},
				{824824821, 824824827},
				{2121212118, 2121212124},
			},
			result:  1227775554,
			wantErr: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.input), func(t *testing.T) {
			result, err := partOne(testCase.input)
			if (err != nil) != testCase.wantErr {
				t.Errorf("PartOne() err %v, wanted %v", err, testCase.wantErr)
			}
			if result != testCase.result {
				t.Errorf("PartOne() got %d, expected %d", result, testCase.result)
			}
		})
	}

	realInput, err := daytwo.ParseInput("./daytwo.txt")
	if err != nil {
		t.Error(err.Error())
	}

	sum, err := partOne(realInput)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println()
	fmt.Println("Part 1 answer:", sum)
	fmt.Println()
}
