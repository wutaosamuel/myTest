package main

import (
	"fmt"
	"sync"
	"time"
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

func main() {
	var wg sync.WaitGroup

	sum := func(o1, o2 *Obj) {
		defer wg.Done()
		o1.mu.Lock()
		time.Sleep(2 * time.Second)
		o2.mu.Lock()
		v := o1.Val + o2.Val
		// FIXME: fail
		// v := o1.Read() + o2.Read()
		fmt.Println("sum value: ", v)
		o1.mu.Unlock()
		o2.mu.Unlock()
	}
	product := func(o1, o2 *Obj) {
		defer wg.Done()
		o1.mu.Lock()
		time.Sleep(2 * time.Second)
		o2.mu.Lock()
		v := o1.Val * o2.Val
		fmt.Println("product value: ", v)
		o1.mu.Unlock()
		o2.mu.Unlock()
	}

	o1 := NewObj()
	o2 := NewObj()
	o1.Write(1)
	o2.Write(10)
	wg.Add(2)
	go sum(o1, o2)
	go product(o1, o2)
	wg.Wait()
}
