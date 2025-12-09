package daynine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func ParseInput(path string) ([]Point, error) {
	file, err := os.Open(path)
	if err != nil {
		return []Point{}, err
	}

	scanner := bufio.NewScanner(file)
	result := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		nStrings := strings.Split(line, ",")
		if len(nStrings) != 2 {
			return []Point{}, fmt.Errorf("Invalid point definition: %s", line)
		}
		x, err := strconv.Atoi(nStrings[0])
		if err != nil {
			return []Point{}, fmt.Errorf("Invalid point definition: %s", line)
		}
		y, err := strconv.Atoi(nStrings[1])
		if err != nil {
			return []Point{}, fmt.Errorf("Invalid point definition: %s", line)
		}
		result = append(result, Point{X: x, Y: y})
	}
	return result, nil
}

func getArea(p1, p2 Point) int {
	length, width := 1+p1.X-p2.X, 1+p1.Y-p2.Y
	return length * width
}

func PartOne(points []Point) int {
	result := 0
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			area := getArea(p1, p2)
			if area > result {
				result = area
			}
		}
	}
	return result
}
