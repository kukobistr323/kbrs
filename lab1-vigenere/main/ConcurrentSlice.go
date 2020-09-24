package main

import "sync"

type ConcurrentSlice struct {
	sync.RWMutex
	items []int
}

func (cs *ConcurrentSlice) Increment(pos int) {
	cs.Lock()
	defer cs.Unlock()
	cs.items[pos] = cs.items[pos] + 1
}
