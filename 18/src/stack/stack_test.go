package stack_test

import (
	"fmt"
	"stack"
	"testing"
)

var s *stack.Stack
var capacity int = 5

func TestNewStack(t *testing.T) {
	s = stack.NewStack(capacity)
	fmt.Println(s)

}

func TestPush(t *testing.T) {
	for x := 1; x <= capacity; x++ {
		err := s.Push(x)
		fmt.Println(s)
		if err != nil {
			t.Errorf("Push(%d): capacity %d err %s", x, capacity, err)

		}

	}
}

func TestPop(t *testing.T) {
	for x := 1; x <= capacity; x++ {
		x, err := s.Pop()
		fmt.Println(s)
		if err != nil {
			t.Errorf("Pop(%d): err %s", x, err)

		}
	}
}
