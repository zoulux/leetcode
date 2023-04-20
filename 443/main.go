package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []byte{
		'a', 'b', 'c', 'c', 'c', 'c', 'c', 'c',
	}

	fmt.Println(compress(a))

	fmt.Println(string(a))

	a = []byte{
		'a', 'a', 'b', 'b', 'c', 'c', 'c',
	}

	fmt.Println(compress(a))
	//fmt.Println(a)

	a = []byte{
		'a',
	}
	fmt.Println(compress(a))

	//fmt.Println(a)

	a = []byte{
		'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b',
	}
	fmt.Println(compress(a))
	//fmt.Println(a)
	// aab
}

// leetcode submit region begin(Prohibit modification and deletion)
func compress(chars []byte) int {
	chars = append(chars, 0) // 可以加一个非法字符，在后面

	idx := 0 // 用于修改原数组的下标
	handle := func(l, r int) {
		chars[idx] = chars[l]
		idx++
		d := r - l
		if d == 1 {
			return
		}
		ds := strconv.Itoa(d)
		for i := 0; i < len(ds); i++ {
			chars[idx] = ds[i]
			idx++
		}
	}

	l, r := 0, 1
	for r < len(chars) {
		if chars[l] == chars[r] {
			r++
		} else {
			handle(l, r)
			l = r
		}
	}

	//handle(l, r)
	return idx
}

func compress2(chars []byte) int {

	idx := 0 // 用于修改原数组的下标
	handle := func(l, r int) {
		chars[idx] = chars[l]
		idx++
		d := r - l
		if d == 1 {
			return
		}
		ds := strconv.Itoa(d)
		for i := 0; i < len(ds); i++ {
			chars[idx] = ds[i]
			idx++
		}
	}

	l, r := 0, 1
	for r < len(chars) {
		if chars[l] == chars[r] {
			r++
		} else {
			handle(l, r)
			l = r
		}
	}

	handle(l, r)
	return idx
}

//leetcode submit region end(Prohibit modification and deletion)
