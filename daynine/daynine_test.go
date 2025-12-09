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
		t.Errorf("PartOne() test got %d, expected %d", result, expected)
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

func TestPartTwo(t *testing.T) {
	input, err := daynine.ParseInput("./daynine_test.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result := daynine.PartTwo(input, daynine.Vector{P: daynine.Point{X: 1, Y: 4}, Direction: 1})
	expected := 24
	if result != expected {
		t.Errorf("PartTwo() test got %d, expected %d", result, expected)
		return
	}

	input, err = daynine.ParseInput("./daynine_test2.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result = daynine.PartTwo(input, daynine.Vector{P: daynine.Point{X: 2, Y: 4}, Direction: 1})
	expected = 66
	if result != expected {
		t.Errorf("PartTwo() test 2 got %d, expected %d", result, expected)
		return
	}

	input, err = daynine.ParseInput("./daynine.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result = daynine.PartTwo(input, daynine.Vector{P: daynine.Point{X: 1848, Y: 52341}, Direction: 1})
	if result <= 102844128 {
		t.Errorf("PartTwo() got %d, which is too low", result)
		return
	}
	if result >= 4616716962 {
		t.Errorf("PartTwo() got %d, which is too high", result)
		return
	}

	fmt.Println()
	fmt.Printf("Part Two answer: %d\n", result)
	fmt.Println()
}
