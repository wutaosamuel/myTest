package main

import (
	"fmt"
	"sync"
)

// Obj test obj
type Obj struct {
	Val int
	mu  *sync.Mutex
}

// NewObj new
func NewObj() *Obj {
	return &Obj{
		Val: -1,
		mu:  &sync.Mutex{}}
}

// Write write val
func (obj *Obj) Write(val int) {
	obj.mu.Lock()
	defer obj.mu.Unlock()
	obj.Val = val

	return
}

// Read read val
func (obj *Obj) Read() int {
	obj.mu.Lock()
	defer obj.mu.Unlock()

	return obj.Val
}

// Double get double
func (obj *Obj) Double() int {
	// FIXME: error on defer unlock
	// obj.mu.Lock()
	// defer obj.mu.Unlock()

	return obj.Sum(obj.Val)
}

// Sum get sum
func (obj *Obj) Sum(a int) int {
	obj.mu.Lock()
	defer obj.mu.Unlock()

	return obj.Val + a
}

func main() {
	o := NewObj()
	o.Write(1)
	fmt.Println("Sum: ", o.Sum(10))
	fmt.Println("Double: ", o.Double())
}
