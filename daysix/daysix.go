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
