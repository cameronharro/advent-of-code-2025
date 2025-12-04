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
	positions := []Point{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			positions = append(positions, Point{X: point.X + x, Y: point.Y + y})
		}
	}

	result := make([]string, 8)
	validPoints := 0
	for _, position := range positions {
		if position.X < 0 || position.Y < 0 || position.X > len(grid[0])-1 || position.Y > len(grid)-1 {
			continue
		}
		result[validPoints] = grid[position.Y][position.X]
		validPoints++
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
