package byslices

type LRUCache struct {
	cap   int
	cache map[int]int
	keys  []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]int),
		keys:  make([]int, 0),
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
	this.keys = append(this.keys, key)
}

func (this *LRUCache) makeRecently(key int) {
	// 将 key 移动到链表尾部表示最近访问

	for i, k := range this.keys {
		if k == key {
			left := this.keys[:i]
			right := this.keys[i+1:]
			this.keys = append([]int{}, left...)
			this.keys = append(this.keys, right...)
			this.keys = append(this.keys, key)
			return
		}
	}

}

func (this *LRUCache) removeOldest() {
	// 删除链表头部表示最久未使用的 key
	if len(this.keys) > 0 {
		k0 := this.keys[0]
		this.keys = this.keys[1:]
		delete(this.cache, k0)
	}
}
