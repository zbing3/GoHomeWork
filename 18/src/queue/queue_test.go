package queue_test

import (
	"fmt"
	"queue"
	"testing"
)

func TestNewRingQueue(t *testing.T) {
	rq := queue.NewRingQueue(5)
	fmt.Println(rq)

}

func TestPush(t *testing.T) {
	capacity := 5
	var x int = 1
	rq := queue.NewRingQueue(capacity)

	push_status := rq.Push(x)
	if push_status != nil {
		t.Errorf("Pus(%d): capacity %d", x, capacity)

	}

}
