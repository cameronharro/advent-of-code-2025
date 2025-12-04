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

type Point struct {
	X int
	Y int
}

func GetSurrounding(point Point, grid [][]string) []string {
	result := make([]string, 8)
	validPoints := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			X, Y := point.X+x, point.Y+y
			if X < 0 || Y < 0 || X > len(grid[0])-1 || Y > len(grid)-1 {
				continue
			}
			result[validPoints] = grid[Y][X]
			validPoints++
		}
	}

	return result[:validPoints]
}

func PartOne(grid [][]string) int {
	result := 0
	for y, line := range grid {
		for x, s := range line {
			if s != "@" {
				continue
			}
			surrounding := GetSurrounding(Point{X: x, Y: y}, grid)
			countRolls := 0
			for _, char := range surrounding {
				if char == "@" {
					countRolls++
				}
			}
			if countRolls < 4 {
				result++
			}
		}
	}
	return result
}
