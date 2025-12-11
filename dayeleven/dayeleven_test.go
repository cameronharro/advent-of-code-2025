package dayeleven_test

import (
	"fmt"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/dayeleven"
)

func TestParseInput(t *testing.T) {
	input, err := dayeleven.ParseInput("./dayeleven_test.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	for k, v := range input {
		fmt.Printf("Node: %s, Options: %v\n", k, v)
	}
}

func TestPartOne(t *testing.T) {
	input, err := dayeleven.ParseInput("./dayeleven_test.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result := dayeleven.PartOne(input)
	expected := 5
	if result != 5 {
		t.Errorf("PartOne got %d, expected %d\n", result, expected)
		return
	}

	input, err = dayeleven.ParseInput("./dayeleven.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result = dayeleven.PartOne(input)

	fmt.Println()
	fmt.Printf("Part One answer: %d\n", result)
	fmt.Println()
}

func TestPartTwo(t *testing.T) {
	input, err := dayeleven.ParseInput("./dayeleven_test2.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result := dayeleven.PartTwo(input)
	expected := 2
	if result != expected {
		t.Errorf("PartTwo got %d, expected %d\n", result, expected)
		return
	}

	input, err = dayeleven.ParseInput("./dayeleven.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	// result = dayeleven.PartTwo(input)
	//
	// fmt.Println()
	// fmt.Printf("Part Two answer: %d\n", result)
	// fmt.Println()
}
