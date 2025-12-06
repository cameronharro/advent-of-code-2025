package daysix

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(path string) ([][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return [][]string{}, err
	}

	result := [][]string{}
	for line := range strings.SplitSeq(string(data), "\n") {
		if line == "" {
			continue
		}
		for i, str := range strings.Fields(line) {
			if len(result) <= i {
				result = append(result, []string{str})
				continue
			}
			result[i] = append(result[i], str)
		}
	}
	return result, nil
}

func Part(valFunc func(row []string) int) func(input [][]string) int {
	return func(input [][]string) int {
		result := 0
		for _, row := range input {
			result += valFunc(row)
		}
		return result
	}
}

func PartOne(input []string) int {
	operator := input[len(input)-1]
	ns := []int{}
	for _, str := range input {
		n, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		ns = append(ns, n)
	}
	switch operator {
	case "+":
		return operateNumbers(ns, func(n1, n2 int) int {
			return n1 + n2
		})
	case "*":
		return operateNumbers(ns, func(n1, n2 int) int {
			return n1 * n2
		})
	default:
		panic(fmt.Sprintf("invalid math operator %v", operator))
	}
}

func operateNumbers(ns []int, mathFunc func(n1, n2 int) int) int {
	result := ns[0]
	for _, n := range ns[1:] {
		result = mathFunc(result, n)
	}
	return result
}
