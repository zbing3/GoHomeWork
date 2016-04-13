package queue

import (
	"errors"
)

type RingQueue struct {
	data     []int
	capacity int
	head     int
	tail     int
}

var (
	ErrQueueFull  = errors.New("queue full")
	ErrQueueEmpty = errors.New("queue empty")
)

func NewRingQueue(capacity int) *RingQueue {
	return &RingQueue{
		data:     make([]int, capacity),
		capacity: capacity,
	}
}

func (q *RingQueue) Push(x int) error {
	if (q.capacity - (q.tail - q.head)) == 0 {
		return ErrQueueFull
	}

	n := q.tail % q.capacity
	q.data[n] = x

	q.tail++
	return nil
}

func (q *RingQueue) Pop() (int, error) {
	if q.tail-q.head == 0 {
		return 0, ErrQueueEmpty
	}

	n := q.head % q.capacity
	x := q.data[n]

	q.head++
	return x, nil
}
