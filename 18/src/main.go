package main

import (
	"fmt"
)

// 看什么看！难道我模仿你们的测试方法不像？
// 赶紧写成单元测试是正经。
// 这回记得交作业。

func q() {
	q := NewRingQueue(3)

	for i := 1; i < 5; i++ {
		err := q.Push(i)
		fmt.Printf("push %d, %v | %+v\n", i, err, *q)
	}

	for i := 0; i < q.capacity+2; i++ {
		x, err := q.Pop()
		fmt.Printf("pop %d, %v | %+v\n", x, err, *q)
	}

	fmt.Println(q.Push(2))
	fmt.Printf("%+v\n", *q)

	fmt.Println(q.Push(3))
	fmt.Printf("%+v\n", *q)
	fmt.Println(q.Pop())
	fmt.Printf("%+v\n", *q)

	fmt.Println(q.Pop())
	fmt.Printf("%+v\n", *q)
}

func p() {
	p := NewPool(3)
	fmt.Println(p)

	x1 := p.Get()
	x2 := p.Get()
	x3 := p.Get()
	fmt.Println(p, x1, x2, x3)

	p.Put(x2)
	p.Put(x1)
	p.Put(x3)
	fmt.Println(p)
}

func s() {
	s := NewStack(5)

	for i := 0; i < 7; i++ {
		fmt.Printf("push %d: %v, %v\n", i, s.Push(i), *s)
	}

	for i := 0; i < 7; i++ {
		x, err := s.Pop()
		fmt.Printf("pop: %d, %v, %v\n", x, err, *s)
	}
}

func main() {
	q()
	p()
	s()
}
