package main

func main() {

}

type Trie struct {
	end  bool
	next [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

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
