package dayfour

import (
	"os"
	"strings"
)

func ParseInput(path string) ([][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")
	result := make([][]string, len(lines))
	for i, l := range lines {
		result[i] = strings.Split(l, "")
	}
	return result, nil
}
