package daysix

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	Numbers  []int
	Operator string
}

func ParseInputOne(path string) ([]Problem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := []Problem{}
	for line := range strings.SplitSeq(string(data), "\n") {
		if line == "" {
			continue
		}

		for i, str := range strings.Fields(line) {
			if len(result) <= i {
				result = append(result, Problem{})
			}

			if str == "*" || str == "+" {
				result[i].Operator = str
				continue
			}

			n, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			result[i].Numbers = append(result[i].Numbers, n)
		}
	}
	return result, nil
}

func ParseInputTwo(path string) ([]Problem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := []Problem{}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")
	problem := Problem{}
	for i := range len(lines[0]) {
		isBlankColumn := true
		nStr := ""
		for _, line := range lines {
			if len(line) <= i {
				continue
			}
			r := line[i]
			if r == 42 || r == 43 {
				problem.Operator = string(r)
				isBlankColumn = false
			}
			if r >= 48 && r <= 57 {
				nStr += string(r)
				isBlankColumn = false
			}
		}

		if nStr != "" {
			n, err := strconv.Atoi(nStr)
			if err != nil {
				return nil, err
			}
			problem.Numbers = append(problem.Numbers, n)
		}

		if isBlankColumn || i >= len(lines[0])-1 {
			result = append(result, problem)
			problem = Problem{}
		}
	}
	return result, nil
}

func EvalProblems(input []Problem) int {
	result := 0
	for _, problem := range input {
		result += SolveProblem(problem)
	}
	return result
}

func SolveProblem(problem Problem) int {
	operator := problem.Operator
	switch operator {
	case "+":
		return operateNumbers(problem.Numbers, func(n1, n2 int) int {
			return n1 + n2
		})
	case "*":
		return operateNumbers(problem.Numbers, func(n1, n2 int) int {
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
