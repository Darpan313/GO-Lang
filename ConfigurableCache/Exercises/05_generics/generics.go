package main

import (
	"errors"
	"fmt"
)


var ErrEmpty = errors.New("stack is empty")

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		// Create a variable of type T to return its default zero-value
		// (e.g, 0 for int, "" for string, nil for pointers)
		var zero T
		return zero, ErrEmpty
	}

	index := len(s.items)-1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func main() {
	s := &Stack[int]{}

	_, err := s.Pop()

	if err != nil {
		wrappedError := fmt.Errorf("loading: %w", err)
		fmt.Printf("Using 'errors.Is: %v\n", errors.Is(wrappedError, ErrEmpty))

		strStack :=&Stack[string]{}
		strStack.Push("Go Generics")
		val, _ := strStack.Pop()
		fmt.Printf("Popped from strStack; %s\n", val)
	}
}