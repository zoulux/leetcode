package main

func main() {

}

type Trie struct {
	end  bool
	next map[rune]*Trie
}

func NewTrie() *Trie {
	return &Trie{next: map[rune]*Trie{}}
}

func Constructor() Trie {
	return Trie{next: map[rune]*Trie{}}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		v -= 'a'
		if node.next[v] == nil {
			node.next[v] = NewTrie()
		}
		node = node.next[v]
	}

	node.end = true
}

func (this *Trie) Search(word string) bool {
	node := this.searchNode(word)
	return node != nil && node.end
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

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchNode(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
