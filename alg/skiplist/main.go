package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func main() {

	sl := New()
	//
	//sl.Insert(float64(100), "foo")
	//sl.Insert(float64(101), "foo")
	//sl.Insert(float64(102), "foo")
	//sl.Insert(float64(200), "foo")
	//sl.Insert(float64(50), "foo")
	//sl.Insert(float64(60), "foo")
	//sl.Insert(float64(12), "foo")
	//sl.Insert(float64(123), "foo")

	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		sl.Insert(float64(x), x)
	}

	s, _ := json.Marshal(sl)

	fmt.Println(string(s))

	//e, ok := sl.Search(float64(100))
	//fmt.Println(ok)
	//fmt.Println(e.Value)
	//e, ok = sl.Search(float64(200))
	//fmt.Println(ok)
	//fmt.Println(e)
	//
	//sl.Insert(50.0, "spam")
	//sl.Insert(20.0, 42)
	//
	//fmt.Println(sl.Len())
	//e = sl.Delete(50.0)
	//fmt.Println(e.Value)
	//fmt.Println(sl.Len())
	//
	//for e := sl.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}
}

const (
	maxLevel int     = 4 // Should be enough for 2^16 elements
	p        float32 = 0.25
)

// Element is an Element of a skiplist.
type Element struct {
	Score   float64
	Value   interface{}
	forward []*Element
}

func newElement(score float64, value interface{}, level int) *Element {
	return &Element{
		Score:   score,
		Value:   value,
		forward: make([]*Element, level),
	}
}

// SkipList represents a skiplist.
// The zero value from SkipList is an empty skiplist ready to use.
type SkipList struct {
	header *Element // header is a dummy element
	len    int      // current skiplist length，header not included
	level  int      // current skiplist level，header not included
}

// New returns a new empty SkipList.
func New() *SkipList {
	return &SkipList{
		header: &Element{forward: make([]*Element, maxLevel)},
	}
}

func randomLevel() int {
	level := 1
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}

// Front returns first element in the skiplist which maybe nil.
func (sl *SkipList) Front() *Element {
	return sl.header.forward[0]
}
func (sl *SkipList) Len() int {
	return sl.len
}

// Next returns first element after e.
func (e *Element) Next() *Element {
	if e != nil {
		return e.forward[0]
	}
	return nil
}

func (sl *SkipList) Search(score float64) (element *Element, ok bool) {
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		return x, true
	}
	return nil, false
}

// Insert (score, value) pair to the skiplist and returns pointer of element.
func (sl *SkipList) Insert(score float64, value interface{}) *Element {
	update := make([]*Element, maxLevel)
	cur := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].Score < score {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	cur = cur.forward[0]

	// Score already presents, replace with new value then return
	if cur != nil && cur.Score == score {
		cur.Value = value
		return cur
	}

	level := randomLevel()
	if level > sl.level {
		level = sl.level + 1
		update[sl.level] = sl.header
		sl.level = level
	}
	e := newElement(score, value, level)
	for i := 0; i < level; i++ {
		e.forward[i] = update[i].forward[i]
		update[i].forward[i] = e
	}
	sl.len++
	return e
}

// Delete remove and return element with given score, return nil if element not present
func (sl *SkipList) Delete(score float64) *Element {
	update := make([]*Element, maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]

	if x != nil && x.Score == score {
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] != x {
				return nil
			}
			update[i].forward[i] = x.forward[i]
		}
		sl.len--
	}
	return x
}
