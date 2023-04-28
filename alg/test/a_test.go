package test

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestA(t *testing.T) {

	m := make(map[int64]byte)

	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				fmt.Println(len(m))
			}
		}
	}()
	ch := make(chan int64, 10)

	done := make(chan int64, 2)

	go func() {
		for i := 0; i < math.MaxInt32; i++ {
			ch <- rand.Int63()
		}
		done <- 1
	}()

	go func() {
		for i := 0; i < math.MaxInt32; i++ {
			ch <- rand.Int63()
		}
		done <- 1
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		x := 0
		for {
			select {
			case n := <-ch:
				m[n]++
			case <-done:
				x++
				if x == 2 {
					return
				}
			}
		}

	}()
	wg.Wait()

	for k, v := range m {
		if v == 1 {
			fmt.Println(k)
		}
	}

}
