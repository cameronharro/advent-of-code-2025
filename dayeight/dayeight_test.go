package dayeight_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/dayeight"
)

func TestParseInput(t *testing.T) {
	result, err := dayeight.ParseInput("./dayeight_test.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	expected := []dayeight.Point{
		{X: 162, Y: 817, Z: 812},
		{X: 57, Y: 618, Z: 57},
		{X: 906, Y: 360, Z: 560},
		{X: 592, Y: 479, Z: 940},
		{X: 352, Y: 342, Z: 300},
	}

	if len(result) != 20 {
		t.Errorf("ParseInput() got len %d, expected len %d", len(result), 20)
	}
	if !slices.Equal(result[:5], expected) {
		t.Errorf("ParseInput() got %v, expected %v", result[:5], expected)
	}
}

func TestPartOne(t *testing.T) {
	input, err := dayeight.ParseInput("./dayeight_test.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	expected := 40
	result := dayeight.PartOne(input, 10)
	if result != expected {
		t.Errorf("PartOne() got %d, expected %d", result, expected)
		return
	}

	input, err = dayeight.ParseInput("./dayeight.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result = dayeight.PartOne(input, 1000)

	fmt.Println()
	fmt.Printf("Part One answer: %d\n", result)
	fmt.Println()
}

func TestPartTwo(t *testing.T) {
	input, err := dayeight.ParseInput("./dayeight_test.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	expected := 25272
	result := dayeight.PartTwo(input)
	if result != expected {
		t.Errorf("PartTwo() got %d, expected %d", result, expected)
		return
	}

	input, err = dayeight.ParseInput("./dayeight.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result = dayeight.PartTwo(input)

	fmt.Println()
	fmt.Printf("Part Two answer: %d\n", result)
	fmt.Println()
}
