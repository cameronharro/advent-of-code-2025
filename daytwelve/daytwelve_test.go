package daytwelve_test

import (
	"fmt"
	"testing"

	"github.com/cameronharro/advent-of-code-2025/daytwelve"
)

func TestParseInput(t *testing.T) {
	input, err := daytwelve.ParseInput("./daytwelve.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	result := daytwelve.PartOne(input)
	fmt.Println()
	fmt.Printf("Part One answer: %d\n", result)
	fmt.Println()
}
