package main

import "math/rand"

// 跳表

func main() {

}

const (
	maxLevel = 16
	p        = 0.25
)

type Node struct {
	Val  int
	Next []*Node
}
type Skiplist struct {
	level  int
	header *Node // 头节点
}

func Constructor() Skiplist {
	return Skiplist{
		level: 0, // 默认层为0
		header: &Node{
			Next: make([]*Node, maxLevel), // 头节点有 maxLevel 个指针
		},
	}
}

func (this *Skiplist) Search(target int) bool {
	cur := this.header // 从头开始
	for i := this.level - 1; i >= 0; i-- {
		// 从最高层开始，依次向下一层查找
		for cur.Next[i] != nil && cur.Next[i].Val < target {
			// 如果当前层的值比目标值小，则继续在当前层前进
			cur = cur.Next[i]
		}

		// 如果当前层找到的这个最小值的下一个值不为nil，并且等于目标值，直接返回 true
		if cur.Next[i] != nil && cur.Next[i].Val == target {
			return true
		}
	}

	// 全局扫一遍都没找到就返回 false
	return false
}

func (this *Skiplist) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (this *Skiplist) randLevel() int {
	l := 1
	// 默认 1 层，每增加一层概率缩小
	for rand.Float64() < p && l < maxLevel {
		l++
	}
	return l
}

func (this *Skiplist) Add(num int) {

	//其实这个变量思考了很久，不知道他在干什么
	// 1. updateNode 其实需要记录着每一层对于 num 这个节点找到的最接近的值
	// 2. 初始化为 header ,是因为方便简化下面的一个新节点加入的时候逻辑
	// 3. 新加入的节点，它可能不会增加新的层，那么他的前驱就是第 1 点描述
	// 4. 新加入的节点，它可能会增加新的层，新层的话这个新节点的前驱节点就是 header，第 2点描述
	updateNode := make([]*Node, maxLevel)
	for i := range updateNode {
		updateNode[i] = this.header
	}

	// 从当前节点开始
	cur := this.header
	for i := this.level - 1; i >= 0; i-- {
		for cur.Next[i] != nil && cur.Next[i].Val < num {
			cur = cur.Next[i]
		}
		updateNode[i] = cur // 记录本层，最接近 num 的节点
	}

	lv := this.randLevel() // 新节点的层数
	newNode := &Node{
		Val:  num,               // num 赋值，这里不判断重复值，重复的依然插入
		Next: make([]*Node, lv), // 有几层，就有几个指针
	}

	for i := 0; i < lv; i++ {
		// 从第 0 层开始，对于每一层都要赋值
		// 新节点的在第 i 层的下一个节点等于 updateNode 所指向的下一个节点
		// updateNode 的下一个节点执行 newNode 这个新节点
		newNode.Next[i], updateNode[i].Next[i] = updateNode[i].Next[i], newNode
	}

	// 跳表最大的层级等于之前的层级和此节点的层级最大值
	this.level = this.max(this.level, lv)
}

func (this *Skiplist) Erase(num int) bool {
	flag := 0              // 是否触发删除，触发删除逻辑则为 1
	newLevel := this.level // 删除节点可能会触发层级降低，因为空的层级没啥意义
	cur := this.header
	// 依旧是遍历
	for i := this.level - 1; i >= 0; i-- {
		// 逻辑类似
		for cur.Next[i] != nil && cur.Next[i].Val < num {
			cur = cur.Next[i]
		}

		// 我这里的实现和大多数人不一样，但是我觉得这样简单一些
		// 这里和 search 实现有点类似，如果找到下一个值和目标值相等
		if cur.Next[i] != nil && cur.Next[i].Val == num {
			cur.Next[i] = cur.Next[i].Next[i] // 当前节点指向的下一个节点等于当前节点的下下个节点，那么下个节点就相当于删除了，可以思考一下指针
			flag = 1                          // 修改标志位
			if this.header.Next[i] == nil {
				// 这里其实是缩层逻辑，如果在这一层，头指针指向了一个空节点，这一层其实就不要了
				newLevel--
			}
		}
	}
	this.level = newLevel
	return flag == 1
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
