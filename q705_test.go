package leetcode

import "testing"

type MyHashSet struct {
	m map[int]bool
}

/** Initialize your data structure here. */
func Constructor2() MyHashSet {
	return MyHashSet{m:map[int]bool{}}
}

func (this *MyHashSet) Add(key int) {
	this.m[key] = true

}

func (this *MyHashSet) Remove(key int) {
	delete(this.m, key)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.m[key]
	return ok
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

// type MyHashSet struct {
// 	bucket [][]int
// 	cap    int
// }

// /** Initialize your data structure here. */
// func Constructor2() MyHashSet {
// 	cap := 1024
// 	return MyHashSet{make([][]int, cap), cap}
// }
// func (this *MyHashSet) hash(key int) int {
// 	return key % this.cap
// }

// func (this *MyHashSet) Add(key int) {
// 	// 10
// 	// 1010
// 	k := this.hash(key)
// 	if len(this.bucket[k]) == 0 {
// 		this.bucket[k] = []int{key}
// 	} else {
// 		for _, v := range this.bucket[k] {
// 			if v == key {
// 				//已经在里面就返回
// 				return
// 			}
// 		}
// 		this.bucket[k] = append(this.bucket[k], key)
// 	}
// }

// func (this *MyHashSet) Remove(key int) {
// 	k := this.hash(key)
// 	for idx, v := range this.bucket[k] {
// 		if v == key {
// 			this.bucket[k] = append(this.bucket[k][:idx], this.bucket[k][idx+1:]...)
// 			return
// 		}
// 	}

// }

// /** Returns true if this set contains the specified element */
// func (this *MyHashSet) Contains(key int) bool {
// 	k := this.hash(key)
// 	for _, v := range this.bucket[k] {
// 		if v == key {
// 			return true
// 		}
// 	}
// 	return false
// }

func TestMyHashSet(t *testing.T) {
	myHashSet := Constructor2()
	myHashSet.Add(1) // set = [1]
	myHashSet.Add(2) // set = [1, 2]

	t.Log(myHashSet.Contains(1)) // 返回 True
	t.Log(myHashSet.Contains(3)) // 返回 True
	myHashSet.Add(2)             // set = [1, 2]
	t.Log(myHashSet.Contains(2)) // 返回 True
	myHashSet.Remove(2)          // 返回 True
	t.Log(myHashSet.Contains(2)) // 返回 True
}
