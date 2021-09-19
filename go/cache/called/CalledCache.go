package utils

import (
	"runtime"
	"sync"
	"time"
)

const (
	DefaultFrequency = int64(60)
	DefaultCycleTime = time.Hour
)

// TODO: update value automatically

// Cache
type CalledCache struct {
	*cache
}

// Item
type Item struct {
	Object interface{}
}

// ItemCounter
type ItemCounter struct {
	Value int64
	mu    *sync.RWMutex
}

// cash
type cache struct {
	cycleTime time.Duration
	frequency int64
	items     map[string]*Item
	counter   map[string]*ItemCounter

	stop chan bool
	mu   *sync.RWMutex
}

// Set
func (c *cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, found := c.items[key]; found {
		item.set(value)
		return
	}
	item := NewItem(value)
	counter := NewItemCounter()
	item.set(value)
	c.items[key] = item
	c.counter[key] = counter
}

// set
func (i *Item) set(value interface{}) {
	i.Object = value
}

// Get
func (c *cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.get(key)
}

// get
func (c *cache) get(key string) (interface{}, bool) {
	if item, found := c.items[key]; found {
		counter := c.counter[key]
		counter.called()
		return item.Object, true
	}
	return nil, false
}

// called
func (ic *ItemCounter) called() {
	ic.mu.Lock()
	defer ic.mu.Unlock()

	ic.Value += 1
}

// GetKeys
func (c *cache) GetKeys() []string {
	c.mu.RLock()
	defer c.mu.Unlock()

	keys := make([]string, len(c.items))
	for k := range c.items {
		keys = append(keys, k)
	}
	return keys
}

// Add
func (c *cache) Add(key string, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, found := c.items[key]; found {
		return false
	}
	item := NewItem(value)
	counter := NewItemCounter()
	item.set(value)
	c.items[key] = item
	c.counter[key] = counter
	return true
}

// Delete
func (c *cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
	delete(c.counter, key)
}

// Flush
func (c *cache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.flush()
}

// flush
func (c *cache) flush() {
	c.items = make(map[string]*Item)
	c.counter = make(map[string]*ItemCounter)
}

// Items
func (c *cache) Items() map[string]Item {
	c.mu.RLock()
	defer c.mu.RUnlock()

	items := make(map[string]Item)
	for k, v := range c.items {
		items[k] = *v
	}
	return items
}

// ItemCount
func (c *cache) ItemCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}

// DeleteNoneFrequency
func (c *cache) DeleteNoneFrequency() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.frequency == -1 {
		c.flush()
		return
	}

	for k, v := range c.counter {
		if v.Value < c.frequency {
			delete(c.counter, k)
			delete(c.items, k)
			continue
		}
		// reset counter
		v.Value = 0
	}
}

// ResetCounter
func (c *cache) ResetCounter() {
	for _, v := range c.counter {
		v.ResetCounter()
	}
}

// ResetCounter
func (ic *ItemCounter) ResetCounter() {
	ic.mu.Lock()
	defer ic.mu.Unlock()

	ic.Value = 0
}

// NewItem
func NewItem(object interface{}) *Item {
	return &Item{
		Object: object,
	}
}

// NewItemCounter
func NewItemCounter() *ItemCounter {
	return &ItemCounter{
		Value: 0,
		mu:    &sync.RWMutex{},
	}
}

// newcache
func newcache(cycleTime time.Duration, frequency int64) *cache {
	cTime := DefaultCycleTime
	f := DefaultFrequency
	if cycleTime != 0 {
		cTime = cycleTime
	}
	if frequency != 0 {
		f = frequency
	}

	return &cache{
		cycleTime: cTime,
		frequency: f,
		items:     make(map[string]*Item),
		counter:   make(map[string]*ItemCounter),

		stop: make(chan bool),
		mu:   &sync.RWMutex{},
	}
}

// run
func (c *cache) run() {
	ticker := time.NewTicker(c.cycleTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.DeleteNoneFrequency()
		case <-c.stop:
			return
		}
	}
}

// stop
func stop(C *CalledCache) {
	C.cache.stop <- true
}

// New
func NewCalledCache(defaultCycleTime time.Duration, defaultFrequency int64) *CalledCache {
	c := newcache(defaultCycleTime, defaultFrequency)
	C := &CalledCache{c}

	go c.run()
	runtime.SetFinalizer(C, stop)

	return C
}
