package main

import "time"

func main() {

}

type CounterLimiter struct {
	timeStamp time.Time
	reqCount  int
	limitNum  int
	interval  int64
}

func New() *CounterLimiter {

	return &CounterLimiter{
		timeStamp: time.Now(),
		reqCount:  0,
		limitNum:  100,
		interval:  1000,
	}
}

func (cl *CounterLimiter) limit() bool {

	now := time.Now()
	if now.Unix() < cl.timeStamp.Unix()+cl.interval {
		if cl.reqCount+1 > cl.limitNum {
			// 限制
			return true
		}
		cl.reqCount++
		return false
	} else {
		cl.timeStamp = now
		cl.reqCount = 1
		return false
	}
}
