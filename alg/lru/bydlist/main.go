package bydlist

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*EntryNode
	head, tail *EntryNode
}

type EntryNode struct {
	key, value int
	pre, next  *EntryNode
}

func initEntryNode(key, value int) *EntryNode {
	return &EntryNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:    make(map[int]*EntryNode),
		head:     initEntryNode(-1, -1),
		tail:     initEntryNode(-1, -1),
		capacity: capacity,
	}
	//设置前驱后驱
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok == false {
		node := initEntryNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) moveToHead(node *EntryNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 指定节点插入头部
func (this *LRUCache) addToHead(node *EntryNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

// 从双向链表中删除指定节点
func (this *LRUCache) removeNode(node *EntryNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	//断开
	node.next = nil
	node.pre = nil
}

// 删除并返回最后一个节点
func (this *LRUCache) removeTail() *EntryNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
