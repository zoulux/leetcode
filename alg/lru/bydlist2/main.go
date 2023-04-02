package bydlist2

import "container/list"

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	list  *list.List
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*list.Element),
		list:  list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.cache[key]; ok {
		// 将 key 移动到链表尾部表示最近访问
		this.makeRecently(val)
		return val.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		// 修改 key 的值，将 key 移动到链表尾部表示最近访问
		//this.cache[key] = value
		v.Value.(*entry).value = value
		this.makeRecently(v)
		return
	}

	if len(this.cache) >= this.cap {
		// 链表头部就是最久未使用的 key
		this.removeOldest()
	}
	// 将新的 key-value 添加链表尾部

	this.list.PushBack(&entry{
		key:   key,
		value: value,
	})
	this.cache[key] = this.list.Back()
}

func (this *LRUCache) makeRecently(ele *list.Element) {
	this.list.MoveToBack(ele)
}

func (this *LRUCache) removeOldest() {
	// 删除链表头部表示最久未使用的 key
	if this.list.Len() > 0 {
		k0 := this.list.Front()
		this.list.Remove(k0)
		delete(this.cache, k0.Value.(*entry).key)
	}
}
