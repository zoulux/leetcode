package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sl := New()
	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	//sl.Delete(1)
	//sl.Delete(2)
	//sl.Delete(3)
	sl.Add(4)
	sl.Add(5)
	sl.Add(6)
	sl.Add(7)

	fmt.Println(sl.Search(1))
	fmt.Println(sl.Search(4))

}

func main2() {

	sl := New()
	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	sl.Add(3)
	sl.Add(10)
	sl.Add(90)
	sl.Add(100)
	sl.Add(50)
	sl.Add(51)

	fmt.Println(sl.Search(1))
	fmt.Println(sl.Search(50))
	fmt.Println(sl.Search(49))
	fmt.Println(sl.Search(1000))
	fmt.Println(sl.Search(51))
	fmt.Println(sl.Search(52))
	sl.Delete(3)
	sl.Delete(50)
	fmt.Println(sl.Search(3))
	fmt.Println(sl.Search(4))
	fmt.Println(sl.Search(50))
	fmt.Println(sl.Search(49))

}

const (
	maxLevel = 32
	p        = 0.25
)

type SkipNode struct {
	Val     int
	forward []*SkipNode
}

type SkipList struct {
	header *SkipNode
	level  int
}

func New() *SkipList {
	return &SkipList{
		header: &SkipNode{
			forward: make([]*SkipNode, maxLevel),
		},
	}
}

func (sl *SkipList) randomLevel() int {
	l := 1
	for rand.Float64() < p && l < maxLevel {
		l++
	}
	return l
}

func (sl *SkipList) Search(val int) bool {
	cur := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].Val < val {
			cur = cur.forward[i]
		}
		if cur.forward[i] != nil && cur.forward[i].Val == val {
			return true
		}
	}
	return false
}

func (sl *SkipList) Add(val int) {
	updateNodes := make([]*SkipNode, maxLevel)

	for i := range updateNodes {
		updateNodes[i] = sl.header
	}

	cur := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].Val < val {
			cur = cur.forward[i]
		}
		updateNodes[i] = cur
	}

	sl.level = max(sl.level, sl.randomLevel())

	newNode := &SkipNode{
		Val:     val,
		forward: make([]*SkipNode, sl.level),
	}

	for i := 0; i < sl.level; i++ {
		newNode.forward[i] = updateNodes[i].forward[i]
		updateNodes[i].forward[i] = newNode
	}
}

func (sl *SkipList) Delete(num int) bool {
	cur := sl.header
	//updateNode := make([]*SkipNode, maxLevel)
	newLevel := sl.level
	flag := 0
	for i := sl.level; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].Val < num {
			cur = cur.forward[i]
		}
		if cur.forward[i] != nil && cur.forward[i].Val == num {
			cur.forward[i] = cur.forward[i].forward[i]
			flag = 1
		}
		if sl.header.forward[i] == nil {
			newLevel--
		}
	}
	sl.level = newLevel
	return flag == 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
