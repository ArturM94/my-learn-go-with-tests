package main

import (
	"sync"
	"sync/atomic"
)

type Incrementer interface {
	Inc()
	Value() int
}

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

type AtomicCounter struct {
	value int32
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt32(&c.value, 1)
}

func (c *AtomicCounter) Value() int {
	return int(atomic.LoadInt32(&c.value))
}
