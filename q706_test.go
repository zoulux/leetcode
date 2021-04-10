package leetcode

import (
	"sort"
	"testing"
)

// type MyHashMap struct {
// 	buncket [][]int
// 	cap     int
// }

// /** Initialize your data structure here. */
// func Constructor3() MyHashMap {
// 	cap := 1021
// 	return MyHashMap{make([][]int, cap), cap}
// }
// func (this *MyHashMap) hash(key int) int {
// 	return key % this.cap
// }

// /** value will always be non-negative. */
// func (this *MyHashMap) Put(key int, value int) {
// 	b := this.hash(key)
// 	lenght := len(this.buncket[b])
// 	if lenght == 0 {
// 		this.buncket[b] = []int{key, value}
// 	} else {
// 		for i := 0; i < lenght; i += 2 {
// 			if this.buncket[b][i] == key {
// 				this.buncket[b][i+1] = value
// 				return
// 			}
// 		}
// 		this.buncket[b] = append(this.buncket[b], key, value)
// 	}
// }

// /** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
// func (this *MyHashMap) Get(key int) int {
// 	b := this.hash(key)
// 	lenght := len(this.buncket[b])

// 	for i := 0; i < lenght; i += 2 {
// 		if this.buncket[b][i] == key {
// 			return this.buncket[b][i+1]
// 		}
// 	}
// 	return -1
// }

// /** Removes the mapping of the specified value key if this map contains a mapping for the key */
// func (this *MyHashMap) Remove(key int) {

// 	b := this.hash(key)
// 	lenght := len(this.buncket[b])

// 	for i := 0; i < lenght; i += 2 {
// 		if this.buncket[b][i] == key {

// 			this.buncket[b] = append(this.buncket[b][:i], this.buncket[b][i+2:]...)

// 			return
// 		}
// 	}
// }

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

// func TestMyHashMap(t *testing.T) {

// 	obj := Constructor3()
// 	obj.Put(1, 2)
// 	obj.Put(1, 3)
// 	obj.Put(2, 3)
// 	obj.Put(1023+2, 4)
// 	t.Log(obj.Get(1))
// 	obj.Remove(2)
// 	t.Log(obj.Get(1))
// 	t.Log(obj.Get(2))

// }

// type MyHashMap struct {
// 	buncket []list.List
// 	cap     int
// }

// type MyHashMapNode struct {
// 	key   int
// 	value int
// }

// /** Initialize your data structure here. */
// func Constructor3() MyHashMap {
// 	cap := 1021
// 	return MyHashMap{make([]list.List, cap), cap}
// }
// func (this *MyHashMap) hash(key int) int {
// 	return key % this.cap
// }

// /** value will always be non-negative. */
// func (this *MyHashMap) Put(key int, value int) {
// 	b := this.hash(key)
// 	for front := this.buncket[b].Front(); front != nil; front = front.Next() {
// 		if front.Value.(*MyHashMapNode).key == key {
// 			front.Value = &MyHashMapNode{key, value}
// 			return
// 		}
// 	}
// 	this.buncket[b].PushBack(&MyHashMapNode{key, value})
// }

// /** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
// func (this *MyHashMap) Get(key int) int {
// 	b := this.hash(key)

// 	for front := this.buncket[b].Front(); front != nil; front = front.Next() {
// 		if front.Value.(*MyHashMapNode).key == key {
// 			return front.Value.(*MyHashMapNode).value
// 		}
// 	}
// 	return -1
// }

// /** Removes the mapping of the specified value key if this map contains a mapping for the key */
// func (this *MyHashMap) Remove(key int) {

// 	b := this.hash(key)

// 	for front := this.buncket[b].Front(); front != nil; front = front.Next() {
// 		if front.Value.(*MyHashMapNode).key == key {
// 			this.buncket[b].Remove(front)
// 			return
// 		}
// 	}
// }

func TestMyHashMap(t *testing.T) {

	obj := Constructor3()
	obj.Put(1, 2)
	obj.Put(1, 3)
	obj.Put(1021+2, 444)
	obj.Put(2, 3)
	t.Log(obj.Get(1))
	obj.Remove(2)
	t.Log(obj.Get(1))
	t.Log(obj.Get(2))
	t.Log(obj.Get(1021 + 2))
	t.Log(obj.Get(1021 + 3))
}

type MyHashMap struct {
	buncketKeys   [][]int
	buncketValues [][]int
	cap           int
}

/** Initialize your data structure here. */
func Constructor3() MyHashMap {
	cap := 1021
	return MyHashMap{make([][]int, cap), make([][]int, cap), cap}
}
func (this *MyHashMap) hash(key int) int {
	return key % this.cap
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	b := this.hash(key)

	keyBuncket := this.buncketKeys[b]

	// 1 101 1001

	keyIdx := sort.SearchInts(keyBuncket, key) // 二分查找
	lenght := len(keyBuncket)
	if keyIdx < lenght && keyBuncket[keyIdx] == key {
		// key 存在
		this.buncketValues[b][keyIdx] = value
	} else {

		if keyIdx >= lenght {
			this.buncketKeys[b] = append(this.buncketKeys[b], key)
			this.buncketValues[b] = append(this.buncketValues[b], value)

		} else {

			this.buncketKeys[b] = append(this.buncketKeys[b], 0)
			this.buncketValues[b] = append(this.buncketValues[b], 0)

			copy(this.buncketKeys[b][keyIdx+1:], this.buncketKeys[b][keyIdx:])
			copy(this.buncketValues[b][keyIdx+1:], this.buncketValues[b][keyIdx:])
			this.buncketKeys[b][keyIdx] = key
			this.buncketValues[b][keyIdx] = value
		}
	}

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	b := this.hash(key)

	buncket := this.buncketKeys[b]

	keyIdx := sort.SearchInts(buncket, key) // 二分查找

	if keyIdx < len(this.buncketKeys[b]) && buncket[keyIdx] == key {
		// key 存在
		return this.buncketValues[b][keyIdx]
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	b := this.hash(key)

	buncket := this.buncketKeys[b]

	// 1 101 1001

	keyIdx := sort.SearchInts(buncket, key) // 二分查找

	if keyIdx < len(this.buncketKeys[b]) && buncket[keyIdx] == key {

		this.buncketKeys[b] = append(this.buncketKeys[b][:keyIdx], this.buncketKeys[b][keyIdx+1:]...)
		this.buncketValues[b] = append(this.buncketValues[b][:keyIdx], this.buncketValues[b][keyIdx+1:]...)

		return
	}
}
