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

func Part(parseNumbers func(row []string) []int) func(input [][]string) int {
	return func(input [][]string) int {
		result := 0
		for _, row := range input {
			ns := parseNumbers(row[:len(row)-1])
			operator := row[len(row)-1]
			result += pickMath(ns, operator)
		}
		return result
	}
}

func pickMath(ns []int, operator string) int {
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

func PartOne(strs []string) []int {
	result := []int{}
	for _, str := range strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		result = append(result, n)
	}
	return result
}
