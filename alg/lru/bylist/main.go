package bylist

import "container/list"

type LRUCache struct {
	cap   int
	cache map[int]int
	keys  *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]int),
		keys:  list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.cache[key]; ok {
		// 将 key 移动到链表尾部表示最近访问
		this.makeRecently(key)
		return val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		// 修改 key 的值，将 key 移动到链表尾部表示最近访问
		this.cache[key] = value
		this.makeRecently(key)
		return
	}

	if len(this.cache) >= this.cap {
		// 链表头部就是最久未使用的 key
		this.removeOldest()
	}
	// 将新的 key-value 添加链表尾部
	this.cache[key] = value
	this.keys.PushBack(key)
}

func (this *LRUCache) makeRecently(key int) {
	// 将 key 移动到链表尾部表示最近访问

	for cur := this.keys.Front(); cur != this.keys.Back(); cur = cur.Next() {
		if k := cur.Value.(int); k == key {
			this.keys.MoveToFront(cur)
			break
		}

	}

}

func (this *LRUCache) removeOldest() {
	// 删除链表头部表示最久未使用的 key
	if this.keys.Len() > 0 {
		k0 := this.keys.Front()
		this.keys.Remove(k0)
		delete(this.cache, k0.Value.(int))
	}
}
