package main

import (
	"container/list"
)

func main() {

}

const pageSize = 1024

type QuickList struct {
	data *list.List
	size int
}

type iterator struct {
	node   *list.Element
	offset int
	ql     *QuickList
}

func (iter *iterator) get() interface{} {
	return iter.page()[iter.offset]
}

func (iter *iterator) page() []interface{} {
	return iter.node.Value.([]interface{})
}

func (iter *iterator) next() bool {
	page := iter.page()
	if iter.offset < len(page)-1 {
		iter.offset++
		return true
	}

	if iter.node == iter.ql.data.Back() {
		iter.offset = len(page)
		return false
	}
	iter.offset = 0
	iter.node = iter.node.Next()
	return true
}

func (iter *iterator) prev() bool {
	if iter.offset > 0 {
		iter.offset--
		return true
	}
	if iter.node == iter.ql.data.Front() {
		iter.offset = -1
		return false
	}
	iter.node = iter.node.Prev()
	prevPage := iter.node.Value.([]interface{})
	iter.offset = len(prevPage) - 1
	return true
}

func (iter *iterator) remove() interface{} {
	page := iter.page()      // 当前页
	val := page[iter.offset] // 当前值
	page = append(page[:iter.offset], page[iter.offset+1:]...)
	if len(page) > 0 {
		iter.node.Value = page
		if iter.offset == len(page) {
			if iter.node != iter.ql.data.Back() {
				iter.node = iter.node.Next()
				iter.offset = 0
			}
		}
	} else {
		if iter.node == iter.ql.data.Back() {
			iter.ql.data.Remove(iter.node)
			iter.node = nil
			iter.offset = 0
		} else {
			nextNode := iter.node.Next()
			iter.ql.data.Remove(iter.node)
			iter.node = nextNode
			iter.offset = 0
		}
	}
	iter.ql.size--
	return val
}

func (ql *QuickList) Find(index int) *iterator {
	if ql == nil {
		panic("")
	}
	if index < 0 || index >= ql.size {
		panic("index out of bound")
	}

	var n *list.Element
	var page []interface{}
	var pageBeg int

	if index < ql.size/2 {
		// index 离左端点更近一些，从左端点开始向后查找
		n = ql.data.Front()
		pageBeg = 0
		for {
			page = n.Value.([]interface{})
			if pageBeg+len(page) > index {
				break
			}
			pageBeg += len(page)
			n = n.Next()
		}
	} else {
		// index 离右端更近一些，从后向前查找
		n = ql.data.Back()
		pageBeg = ql.size
		for {
			page = n.Value.([]interface{})
			pageBeg -= len(page)
			if pageBeg <= index {
				break
			}
			n = n.Prev()
		}
	}

	// pageBeg 统计了前面页总和
	pageOffset := index - pageBeg // 相对于当前页的偏移量
	return &iterator{
		node:   n,
		offset: pageOffset,
		ql:     ql,
	}
}
func (ql *QuickList) Add(val interface{}) {
	ql.size++

	if ql.data.Len() == 0 {
		// 长度为 0，则加一页
		page := make([]interface{}, 0, pageSize)
		page = append(page, val)
		ql.data.PushBack(page)
		return
	}
	backNode := ql.data.Back()
	backPage := backNode.Value.([]interface{})
	if len(backPage) == cap(backPage) {
		// 末尾页已经满了，则加一页
		page := make([]interface{}, 0, pageSize)
		page = append(page, val)
		ql.data.PushBack(page)
		return
	}

	backNode.Value = append(backPage, val) // 在尾页上添加一个数据,其实也不会变化
}

func (ql *QuickList) Insert(index int, val interface{}) {
	if index == ql.size {
		ql.Add(val)
		return
	}

	iter := ql.Find(index)
	page := iter.node.Value.([]interface{})
	if len(page) < pageSize {
		page = append(page[:iter.offset+1], page[iter.offset:]...) // 将 offset 位置空出
		page[iter.offset] = val                                    // 将val 赋值到 offset 位置
		iter.node.Value = page                                     // 页赋值到双向列表
		ql.size++
		return
	}

	// 切半主要是为了优化性能，因为可能后面还要插入数据，这个时候不用页分裂

	var nextPage []interface{}
	nextPage = append(nextPage, page[pageSize/2:]...) // 准备将后 512 个数移动到下一页去
	page = page[:pageSize/2]                          //当前页拥有前 512 个数
	if iter.offset < len(page) {
		// 想插入到当前页
		page = append(page[:iter.offset+1], page[iter.offset:]...) // 移动 ，空出 offset
		page[iter.offset] = val
	} else {
		// 想插入到下一页
		i := iter.offset - pageSize/2                      // 在下一页的一个偏移量
		nextPage = append(nextPage[:i+1], nextPage[i:]...) // 空出来 i 的位置
		nextPage[i] = val
	}

	iter.node.Value = page                   // 当前页赋值
	ql.data.InsertAfter(nextPage, iter.node) // 将下一页插入当前页的后面
	ql.size++
}

func (ql *QuickList) Len() int {
	return ql.size
}

type Consumer func(i int, v interface{}) bool

func (ql *QuickList) ForEach(consumer Consumer) {

	if ql == nil {
		panic("list is nil")
	}
	if ql.Len() == 0 {
		return
	}
	iter := ql.Find(0)
	i := 0
	for iter.next() {
		goNext := consumer(i, iter.get())
		if !goNext {
			break
		}
		i++
	}
}
