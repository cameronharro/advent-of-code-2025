package dayseven

import (
	"os"
	"strings"
)

func ParseInput(path string) ([][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return [][]string{}, err
	}

	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")
	result := make([][]string, len(lines))
	for i, line := range lines {
		result[i] = strings.Split(line, "")
	}
	return result, nil
}

func PartOne(grid [][]string) int {
	result := 0
	for y, line := range grid {
		if y >= len(grid)-1 {
			continue
		}
		for x, char := range line {
			if char == "S" || char == "|" {
				below := grid[y+1][x]
				if below != "^" {
					grid[y+1][x] = "|"
				} else {
					result++
					grid[y+1][x-1] = "|"
					grid[y+1][x+1] = "|"
				}
			}
		}
	}
	return result
}
