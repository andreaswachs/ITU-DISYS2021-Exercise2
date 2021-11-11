package main

import "sync"

type Clock struct {
	count uint64
	lock  sync.Mutex
}

func NewClock() Clock {
	return Clock{count: 0}
}

func (c *Clock) Increment() uint64 {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count++
	return c.count
}

func (c *Clock) GetCount() uint64 {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.count
}

func (c1 *Clock) Update(c2 uint64) {
	c1.lock.Lock()
	defer c1.lock.Unlock()

	if c1.count > c2 {
		c1.count++
	} else {
		c1.count = c2 + 1
	}
}
