package main

func main() {

}

type UnionFindSets struct {
	Parent []int
}

func NewUnionFind(n int) *UnionFindSets {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFindSets{Parent: parent}
}

func (uf *UnionFindSets) Find(i int) int {
	if uf.Parent[i] != i {
		uf.Parent[i] = uf.Find(uf.Parent[i])
	}
	return uf.Parent[i]
}
func (uf *UnionFindSets) Union(i, j int) {
	uf.Parent[uf.Find(i)] = uf.Find(j)
}
