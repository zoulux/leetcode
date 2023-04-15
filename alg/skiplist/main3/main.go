package main

import (
	"fmt"
	"math/rand"
)

func main() {

	sl := Constructor()
	sl.Add(1)
	sl.Add(3)
	sl.Add(10)
	sl.Add(15)

	fmt.Println(sl.Search(10))
	fmt.Println(sl.Search(11))
}

const maxLevel = 32
const pFactor = 0.25

type SkiplistNode struct {
	val     int
	forward []*SkiplistNode
}

type Skiplist struct {
	head  *SkiplistNode
	level int
}

func Constructor() Skiplist {
	return Skiplist{&SkiplistNode{-1, make([]*SkiplistNode, maxLevel)}, 0}
}

func (Skiplist) randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

func (s *Skiplist) Search(target int) bool {
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 target 的元素
		for curr.forward[i] != nil && curr.forward[i].val < target {
			curr = curr.forward[i]
		}
	}
	curr = curr.forward[0]
	// 检测当前元素的值是否等于 target
	return curr != nil && curr.val == target
}

func (s *Skiplist) Add(num int) {
	newLevel := rand.Intn(maxLevel) + 1
	newNode := &SkiplistNode{num, make([]*SkiplistNode, newLevel)}
	cur := s.head

	for i := newLevel - 1; i >= 0; i-- {
		for cur.forward[i] != nil {
			if num > cur.forward[i].val {
				cur = cur.forward[i]
			} else {
				break
			}
		}
		newNode.forward[i] = cur.forward[i]
		cur.forward[i] = newNode
	}
}

func (s *Skiplist) Erase(num int) bool {
	update := make([]*SkiplistNode, maxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	curr = curr.forward[0]
	// 如果值不存在则返回 false
	if curr == nil || curr.val != num {
		return false
	}
	for i := 0; i < s.level && update[i].forward[i] == curr; i++ {
		// 对第 i 层的状态进行更新，将 forward 指向被删除节点的下一跳
		update[i].forward[i] = curr.forward[i]
	}
	// 更新当前的 level
	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}
	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
