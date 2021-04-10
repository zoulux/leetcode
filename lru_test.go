package leetcode

import "container/list"

type LRUCache struct {
	Cap  int
	Keys map[int]*list.Element
	List *list.List
}

type pair struct {
	K, V int
}

func Contructor(cap int) LRUCache {
	return LRUCache{
		Cap:  cap,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if el, ok := c.Keys[key]; ok {
		c.List.MoveToFront(el)
		return el.Value.(pair).V
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if el, ok := c.Keys[key]; ok {
		el.Value = pair{K: key, V: value}
		c.List.MoveToFront(el)
	} else {
		el := c.List.PushFront(pair{K: key, V: value})
		c.Keys[key] = el
	}

	if c.List.Len() > c.Cap {
		el := c.List.Back()
		c.List.Remove(el)
		delete(c.Keys, el.Value.(pair).K)
	}
}
