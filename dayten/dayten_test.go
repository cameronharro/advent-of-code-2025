package dayten_test

import (
	"fmt"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/dayten"
)

func TestParseInput(t *testing.T) {
	result, err := dayten.ParseInput("./dayten_test.txt")
	if err != nil {
		t.Error(err.Error())
	}
	for i, machine := range result {
		fmt.Printf("ParseInput() machine %d: %v\n", i, machine)
	}
}

func TestPartOne(t *testing.T) {
	input, err := dayten.ParseInput("./dayten_test.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result := dayten.PartOne(input)
	expected := 7
	if result != expected {
		t.Errorf("PartOne() got %d, expected %d", result, expected)
		return
	}

	input, err = dayten.ParseInput("./dayten.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result = dayten.PartOne(input)

	fmt.Println()
	fmt.Printf("Part One Answer: %d\n", result)
	fmt.Println()
}
