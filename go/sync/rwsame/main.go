package main

import (
	"fmt"
	"sync"
	"time"
)

// Obj test obj
type Obj struct {
	Val int
	rw  *sync.RWMutex
}

// NewObj new
func NewObj() *Obj {
	return &Obj{
		Val: -1,
		rw:  &sync.RWMutex{}}
}

// Write write val
func (obj *Obj) Write(val int) {
	obj.rw.Lock()
	defer obj.rw.Unlock()
	obj.Val = val

	return
}

// Read read val
func (obj *Obj) Read() int {
	obj.rw.RLock()
	defer obj.rw.RUnlock()

	return obj.Val
}

func main() {
	var wg sync.WaitGroup

	sum := func(o1, o2 *Obj) {
		defer wg.Done()
		o1.rw.Lock()
		time.Sleep(2 * time.Second)
		o2.rw.Lock()
		v := o1.Val + o2.Val
		// FIXME: fail
		//v := o1.Read() + o2.Read()
		fmt.Println("sum value: ", v)
		o1.rw.Unlock()
		o2.rw.Unlock()
	}
	product := func(o1, o2 *Obj) {
		defer wg.Done()
		o1.rw.Lock()
		time.Sleep(2 * time.Second)
		o2.rw.Lock()
		v := o1.Val * o2.Val
		fmt.Println("product value: ", v)
		o1.rw.Unlock()
		o2.rw.Unlock()
	}

	o1 := NewObj()
	o2 := NewObj()
	o1.Write(1)
	o2.Write(10)
	wg.Add(2)
	go sum(o1, o2)
	go product(o1, o2)
	// fail
	// go sum(o1, o1)
	// go product(o2, o2)
	wg.Wait()
}
