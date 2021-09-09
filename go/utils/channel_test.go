package utils

import (
	"time"
	"testing"
	"errors"
)

// TestIsChannelError
func TestIsChannelError(t *testing.T) {
	ch1 := make(chan error)
	ch2 := make(chan error)
	ch3 := make(chan error)
	e := "ch3 error works"
	go func() {
		time.Sleep(10 * time.Second)
		ch1 <- nil
	} ()
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- nil
	} ()
	go func() {
		time.Sleep(15 * time.Second)
		ch3 <- errors.New(e)
	} ()

	if err := IsChannelError(ch1); err != nil {
		t.Fatal(err)
	} else {
		t.Log("ch1 work")
	}
	if err := IsChannelError(ch2); err != nil {
		t.Fatal(err)
	} else {
		t.Log("ch2 work")
	}
	if err := IsChannelError(ch3); err != nil {
		if err.Error() != e {
			t.Fatal("ch3 error not equal")
		}
		t.Log(err)
	} else {
		t.Fatal("ch3 should output err")
	}
}
