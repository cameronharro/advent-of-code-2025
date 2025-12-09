package daynine_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/daynine"
)

func TestParseInput(t *testing.T) {
	input, err := daynine.ParseInput("./daynine_test.txt")
	if err != nil {
		t.Error(err)
		return
	}

	expected := []daynine.Point{
		{X: 7, Y: 1},
		{X: 11, Y: 1},
		{X: 11, Y: 7},
		{X: 9, Y: 7},
		{X: 9, Y: 5},
		{X: 2, Y: 5},
		{X: 2, Y: 3},
		{X: 7, Y: 3},
	}

	if !slices.Equal(input, expected) {
		t.Errorf("ParseInput got %v, expected %v", input, expected)
	}
}

func TestPartOne(t *testing.T) {
	input, err := daynine.ParseInput("./daynine_test.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result := daynine.PartOne(input)
	expected := 50
	if result != expected {
		t.Errorf("PartOne() got %d, expected %d", result, expected)
		return
	}

	input, err = daynine.ParseInput("./daynine.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result = daynine.PartOne(input)

	fmt.Println()
	fmt.Printf("Part One answer: %d\n", result)
	fmt.Println()
}
