package dayfour_test

import (
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
