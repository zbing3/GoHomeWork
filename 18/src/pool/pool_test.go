package pool_test

import (
	"fmt"
	"pool"
	"testing"
)

var p pool.Pool
var capacity int = 5

func TestNewPool(t *testing.T) {
	p = pool.NewPool(capacity)
	fmt.Println(p)

}

func TestGet() {

}
