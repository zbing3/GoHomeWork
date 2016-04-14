package queue_test

import (
	"fmt"
	"queue"
	"testing"
)

var rq *queue.RingQueue
var capacity int = 5

func TestNewRingQueue(t *testing.T) {
	rq = queue.NewRingQueue(capacity)
	fmt.Println(rq)

}

func TestPush(t *testing.T) {

	for x := 1; x <= capacity+1; x++ {
		push_status := rq.Push(x)
		fmt.Println(rq)
		if push_status != nil {
			t.Errorf("Pus(%d): capacity %d err %s", x, capacity, push_status)

		}
	}

}

func TestPop(t *testing.T) {
	for x := 1; x <= capacity; x++ {
		value, err := rq.Pop()
		if err != nil {
			t.Errorf("TestPop(): err %s", err)
		}
		fmt.Println(rq, err)
		fmt.Println(value)
	}
}
