package test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

func Test1(t *testing.T) {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
		}
		a[k] = 100 + v
	}

	fmt.Println(a)
}

func Test3(t *testing.T) {
	type str struct {
		A string
	}
	b := new(map[int]int)
	c := new(str)
	d := new([]int)
	e := new([3]int)
	fmt.Println(b, c, d, e)
}

func Test2(t *testing.T) {

	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
		}
		a[k] = 100 + v
	}

	fmt.Println(a)
}

func TestCat(t *testing.T) {

	var wg sync.WaitGroup

	chcat := make(chan bool)
	chdog := make(chan bool)
	chfish := make(chan bool)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for _ = range chcat {
			fmt.Println("cat")
			chdog <- true
		}
		close(chdog)
	}()

	go func() {
		defer wg.Done()

		for _ = range chdog {
			fmt.Println("dog")
			chfish <- true
		}
		close(chfish)

	}()

	go func() {
		defer wg.Done()

		i := 0
		for _ = range chfish {
			i++
			fmt.Println("fish")
			if i < 10 {
				chcat <- true
			} else {
				close(chcat)
			}
		}
	}()

	chcat <- true
	wg.Wait()
}

func TestAlign(t *testing.T) {
	type S struct {
		a int64  // 8
		b string // 16
		c int32  // 4
		d byte   // 1
	}
	//
	////01234567  16+8=    24+4 28
	//
	var s S
	//
	fmt.Println(unsafe.Alignof(s))
	fmt.Println(unsafe.Sizeof(s))
}

func TestAppend1(t *testing.T) {
	ss := []int{1, 2} // 2=>4
	fmt.Println(len(ss), cap(ss))
	ss = append(ss, 3, 4, 5)
	fmt.Println(len(ss), cap(ss))
}

func TestAppend2(t *testing.T) {
	ss := []string{"my", "name", "is"}
	fmt.Println(len(ss), cap(ss)) // 3,3
	ss = append(ss, "egg")        // newcap 3 double 6 newlen 4 > 6 => 6//
	fmt.Println(len(ss), cap(ss))
	// 0-16-32-48-64-80-96
}

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(1)
	}()
}
