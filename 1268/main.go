package main

import (
	"fmt"
)

func main() {
	fmt.Println(suggestedProducts([]string{
		"mobile", "mouse", "moneypot", "monitor", "mousepad",
	}, "mouse"))

}

func suggestedProducts(products []string, searchWord string) [][]string {
	t := Constructor()
	for _, p := range products {
		t.Insert(p)
	}
	var ans [][]string
	for i := range searchWord {
		ans = append(ans, t.Search(searchWord[:i+1]))
	}
	return ans
}

type Trie struct {
	end  bool
	next [26]*Trie
}

func Constructor() *Trie {
	return &Trie{}
}

// 常规写法
func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		v -= 'a'
		if node.next[v] == nil {
			node.next[v] = new(Trie)
		}
		node = node.next[v]
	}
	node.end = true
}

func (this *Trie) Search(word string) []string {
	node := this.searchNode(word) // 找到word的结尾
	if node == nil {
		// 没找到
		return nil
	}

	var ans []string
	var dfs func(node *Trie, arr []byte)
	dfs = func(node *Trie, arr []byte) {
		if len(ans) == 3 {
			// 找齐了，不需要继续去找
			return
		}

		if node.end {
			// 当前是一个字符串的结尾，那就加进来
			ans = append(ans, string(arr))
		}

		for i, v := range node.next {
			if v != nil {
				// 当前字符不为空，那就根据当前字符索引继续搜
				dfs(v, append(arr, 'a'+byte(i)))
			}
		}
	}

	dfs(node, []byte(word)) // 当前路由的前缀加上,，开始搜
	return ans
}

func (this *Trie) searchNode(word string) *Trie {
	node := this
	for _, v := range word {
		v -= 'a'
		if node.next[v] == nil {
			return nil
		}
		node = node.next[v]
	}
	return node
}
