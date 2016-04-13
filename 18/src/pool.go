package main

import (
	"unsafe"
)

type Pool []int

func NewPool(capacity int) Pool {
	p := make(Pool, capacity+1)

	for i := 0; i <= capacity; i++ {
		p[i] = i + 1
	}

	return p
}

func (p Pool) Get() *int {
	n := p[0]
	if n == len(p) {
		return nil
	}

	p[0] = p[n]
	return &p[n]
}

// TODO: 判断地址合法性；判断是否已经放回。
func (p Pool) Put(x *int) {
	n := (uintptr(unsafe.Pointer(x)) - uintptr(unsafe.Pointer(&p[0]))) / unsafe.Sizeof(p[0])
	p[n] = p[0]
	p[0] = int(n)
}
