package main

import (
	"errors"
	"fmt"
)
// 1. Define the package-level sentinel error
var ErrEmpty = errors.New("stack is empty")

type Stack struct {
	items []int
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrEmpty
	}

	index := len(s.items)-1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func main() {
	s := &Stack{}

	// Try popping empty stack
	_, err := s.Pop()

	if err != nil {
		// Wrap the error to add context
		wrappedError := fmt.Errorf("loading: %w", err)
		fmt.Printf("Error string: %v\n\n", wrappedError)

		// Check using standard equality
		fmt.Printf("Using '==': %v\n", wrappedError == ErrEmpty)

		// Check using errors.Is
		fmt.Printf("Using 'errors.Is': %v\n", errors.Is(wrappedError, ErrEmpty))
	}
}