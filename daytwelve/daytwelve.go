package daytwelve

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Tree struct {
	Area     int
	Presents []int
}

func Map[T, V any](slice []T, conversionFunc func(E T) (V, error)) ([]V, error) {
	result := make([]V, len(slice))
	for i, e := range slice {
		v, err := conversionFunc(e)
		if err != nil {
			return nil, err
		}
		result[i] = v
	}
	return result, nil
}

func ParseInput(path string) ([]Tree, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	result := []Tree{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		treeRe := regexp.MustCompile(`(\d\dx\d\d):(.*)`)
		matches := treeRe.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}
		if len(matches) != 3 {
			return nil, fmt.Errorf("Invalid tree: %v\n", line)
		}

		width, _ := strconv.Atoi(matches[1][:2])
		height, _ := strconv.Atoi(matches[1][3:])
		numbers, _ := Map(strings.Fields(matches[2]), func(e string) (int, error) {
			return strconv.Atoi(e)
		})
		tree := Tree{Area: width * height, Presents: numbers}
		result = append(result, tree)
	}

	return result, nil
}

func PartOne(input []Tree) int {
	result := 0
	for _, tree := range input {
		if solveTree(tree) {
			result++
		}
	}
	return result
}

func solveTree(tree Tree) bool {
	presentSizes := [6]int{6, 7, 7, 7, 5, 7}
	minimumSpace := 0
	for i, size := range presentSizes {
		if i >= len(tree.Presents) {
			continue
		}
		minimumSpace += size * tree.Presents[i]
	}
	fmt.Println(minimumSpace, tree.Area)
	return minimumSpace <= tree.Area
}
