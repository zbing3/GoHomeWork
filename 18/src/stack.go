package main

import (
	"errors"
)

type Stack []int

func NewStack(capacity int) *Stack {
	s := make(Stack, 0, capacity)
	return &s
}

func (s *Stack) Push(x int) error {
	n := len(*s)
	if n == cap(*s) {
		return errors.New("stack is full")
	}

	*s = (*s)[:n+1]
	(*s)[n] = x
	return nil
}

func (s *Stack) Pop() (int, error) {
	n := len(*s)
	if n == 0 {
		return 0, errors.New("stack is empty")
	}

	x := (*s)[n-1]
	*s = (*s)[:n-1]

	return x, nil
}
