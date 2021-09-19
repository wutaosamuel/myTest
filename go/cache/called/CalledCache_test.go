package utils

import (
	"testing"
	"time"
)

const isDebug = true

func TestCache(t *testing.T) {
	var found bool
	cc := NewCalledCache(time.Minute, 10)

	itemLen := cc.ItemCount()
	if itemLen != 0 {
		t.Error("Just stat, it should 0 items")
	}

	cc.Set("a", 1)
	cc.Set("b", "b")

	a, found := cc.Get("a")
	if !found {
		t.Error("Already set a element, should be found")
	}
	if a == nil {
		t.Error("a value is 1, not nil")
	}
	// a_value, ok = a.(int)
	if a_value := a.(int); a_value+2 != 3 {
		t.Error("a value is 1 (int), 1 + 2 should be equal 3")
	}
	debugLog(t, a)

	b, found := cc.Get("b")
	if !found {
		t.Error("Already set a element, should be found")
	}
	if b == nil {
		t.Error("a value is b, not nil")
	}
	if b_value := b.(string); b_value+"B" != "bB" {
		t.Error("a value is 1 (int), 1 + 2 should be equal 3")
	}
	debugLog(t, b)
}

func TestCacheTime(t *testing.T) {
	cc := NewCalledCache(time.Minute, 8)

	cc.Set("a", 1)
	cc.Set("b", "b")
	cc.Set("Del", "delete")

	aGet := func() {
		a, found := cc.Get("a")
		if !found {
			t.Error("Already set a element, should be found")
		}
		if a == nil {
			t.Error("a value is 1, not nil")
		}
		debugLog(t, a)
	}
	bGet := func() {
		b, found := cc.Get("b")
		if !found {
			t.Error("Already set a element, should be found")
		}
		if b == nil {
			t.Error("a value is b, not nil")
		}
		debugLog(t, b)
	}
	cGet := func() {
		c, found := cc.Get("Del")
		if !found {
			t.Error("Already set a element, should be found")
		}
		if c == nil {
			t.Error("c value is delete, not nil")
		}
		debugLog(t, c)
	}

	for i := 0; i <= 20; i++ {
		go aGet()

		if i&1 == 0 {
			go bGet()
		}

		if i%3 == 0 {
			go cGet()
		}
	}

	<-time.After(time.Minute)

	aGet()
	bGet()
	// should not get c now
	c, found := cc.Get("Del")
	if found {
		t.Error("Del(c) key, should not get value")
	}
	if c != nil {
		t.Error("c value should be delete, should be nil now")
	}
}

// TestAutoFlush
func TestAutoFlush(t *testing.T) {
	var found bool
	cc := NewCalledCache(time.Minute, -1)

	itemLen := cc.ItemCount()
	if itemLen != 0 {
		t.Error("Just stat, it should 0 items")
	}

	cc.Set("a", 1)
	cc.Set("b", "b")

	debugLog(t, "cache contains ", cc.ItemCount())
	if cc.ItemCount() != 2 {
		t.Error("items should be contain")
	}

	<-time.After(time.Minute)

	debugLog(t, "cache contains ", cc.ItemCount())
	if cc.ItemCount() != 0 {
		t.Error("items should not be contain any item")
	}

	a, found := cc.Get("a")
	if found {
		t.Error("should not be found a")
	}
	if a != nil {
		t.Error("a should not have the value")
	}

	b, found := cc.Get("b")
	if found {
		t.Error("should not be found b")
	}
	if b != nil {
		t.Error("b should not have the value")
	}
}

func debugLog(t *testing.T, args ...interface{}) {
	if isDebug {
		t.Log(args...)
	}
}
