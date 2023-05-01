package main

import "fmt"

func main() {

	fmt.Println(minReorder(5, [][]int{
		{4, 3},
		{2, 3},
		{1, 2},
		{1, 0},
	}))
}

type UnionSet struct {
	Father []int
}

// 初始化
func (u UnionSet) Init(n int) []int {
	//var father [math.MaxInt]int
	for i := 1; i < n; i++ {
		u.Father[i] = i
	}
	return u.Father
}

// 合并
func (u UnionSet) Union(i int, j int) {
	iFather := u.Father[i]      //找到i的祖先节点
	jFather := u.Father[j]      //找到j的祖先节点
	u.Father[jFather] = iFather //i的祖先指向j的祖先
}

// 查找
func (u UnionSet) Find(i int) int {
	if i == u.Father[i] {
		return i
	} else {
		u.Father[i] = u.Find(u.Father[i]) //进行路径压缩
		return u.Father[i]                //返回父节点
	}
}

func minReorder(n int, connections [][]int) int {
	set := UnionSet{
		Father: make([]int, n),
	}
	ret := 0
	set.Init(n)
	for _, v := range connections {
		if set.Find(v[0]) == 0 && set.Find(v[1]) != 0 {
			set.Union(v[0], v[1])
			ret++
		} else if set.Find(v[1]) == 0 {
			set.Union(v[1], v[0])
		}
	}
	for i := len(connections) - 1; i >= 0; i-- {
		v := connections[i]
		if set.Find(v[0]) == 0 && set.Find(v[1]) != 0 {
			set.Union(v[0], v[1])
			ret++
		} else if set.Find(v[1]) == 0 {
			set.Union(v[1], v[0])
		}
	}
	return ret
}
