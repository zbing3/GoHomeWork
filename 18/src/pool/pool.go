package pool

import (
	"errors"
	"unsafe"
)

var (
	ErrPollAddress = errors.New("Poll Address error")
)

type Pool []int

// [1, 2, 3, 4, 5, 6]
// [0, 1, 2, 3, 4, 5]
func NewPool(capacity int) Pool {
	p := make(Pool, capacity+1)

	for i := 0; i <= capacity; i++ {
		p[i] = i + 1
	}

	return p
}

// 获取数据函数
func (p Pool) Get() *int {
	n := p[0]
	if n == len(p) {
		return nil
	}

	p[0] = p[n]
	return &p[n]
}

// TODO: 判断地址合法性；判断是否已经放回。
// bitmap 判断是否已经放回
func (p Pool) Put(x *int) error {
	poll_head_address := uintptr(unsafe.Pointer(&p[0]))
	poll_tail_address := uintptr(unsafe.Pointer(&p[len(p)+1]))
	n := (uintptr(unsafe.Pointer(x)) - uintptr(unsafe.Pointer(&p[0]))) / unsafe.Sizeof(p[0])
	if n < poll_head_address && n > poll_tail_address {
		return ErrPollAddress

	}

	p[n] = p[0]
	p[0] = int(n)
	return nil
}
