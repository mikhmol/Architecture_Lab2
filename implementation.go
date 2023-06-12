package lab2

import (
	"fmt"
	"strings"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", fmt.Errorf("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func isOperator(token string) bool {
	operators := []string{"+", "-", "*", "/"}
	for _, operator := range operators {
		if operator == token {
			return true
		}
	}
	return false
}

func PrefixToPostfix(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("empty prefix expression")
	}

	stack := Stack{}
	tokens := strings.Split(input, " ")

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if isOperator(token) {
			if stack.IsEmpty() {
				return "", fmt.Errorf("invalid prefix expression")
			}

			operand1, err := stack.Pop()
			if err != nil {
				return "", err
			}

			operand2, err := stack.Pop()
			if err != nil {
				return "", err
			}

			expression := operand1 + " " + operand2 + " " + token
			stack.Push(expression)
		} else {
			stack.Push(token)
		}
	}

	result, err := stack.Pop()
	if err != nil {
		return "", err
	}

	return result, nil
}
