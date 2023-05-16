package main

import (
	"sync"
	"time"
)

func main() {

}

type SlidingTimeWindow struct {
	reqCount     int
	slots        []int
	limitNum     int
	windowLength int64
	windowNum    int
	lock         sync.RWMutex
}

func New() *SlidingTimeWindow {
	stw := SlidingTimeWindow{
		reqCount:     0,
		slots:        make([]int, 0),
		limitNum:     100,
		windowLength: 100,
		windowNum:    10,
	}

	stw.slots = append(stw.slots, 0)

	go func() {

		ticker := time.NewTicker(time.Duration(stw.windowLength) * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				stw.lock.Lock()
				stw.slots = append(stw.slots, 0)
				if len(stw.slots) > stw.windowNum {
					stw.reqCount -= stw.slots[0]
					stw.slots = stw.slots[1:]
				}
				stw.lock.Unlock()
			}
		}
	}()

	return &stw
}

func (stw *SlidingTimeWindow) Limit() bool {
	stw.lock.Lock()
	defer stw.lock.Unlock()

	if (stw.reqCount + 1) > stw.limitNum {
		return true
	}
	stw.slots[len(stw.slots)-1]++
	stw.reqCount++

	return false
}
