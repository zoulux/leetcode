package main

import "fmt"

func main() {

	sc := Constructor([]string{
		"ab", "ba", "aaab", "abab", "baa",
	})

	res := []bool{false, false, false, false, false, true, true,
		true, true, true, false, false, true, true, true, true, false, false, false, true, true, true, true, true, true, false, true, true, true, false}
	query := "aaaaaba"
	//query := "aaaaabababbbababbbbababaaabaaa"

	fmt.Println(len(res))
	fmt.Println(len(query))

	for i := 0; i < len(query); i++ {
		fmt.Println(res[i] == sc.Query(query[i]))
	}

}

type StreamChecker struct {
	words  []string // 原来的
	words2 []string // 可能存在的
	s      string
}

func Constructor(words []string) StreamChecker {
	return StreamChecker{
		words:  words,
		words2: make([]string, 0),
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	tmp := make([]string, 0)

	ret := false

	if len(this.words2) != 0 {
		for i := 0; i < len(this.words2); i++ {
			cur := this.words2[i]
			if len(this.s) < len(cur) && cur[len(this.s)] == letter {
				tmp = append(tmp, cur)
				if cur == this.s+string(letter) {
					ret = true
				}
			}
		}
	} else {
		for i := 0; i < len(this.words); i++ {
			cur := this.words[i]
			if len(this.s) < len(cur) && cur[len(this.s)] == letter {
				tmp = append(tmp, cur)
				if cur == this.s+string(letter) {
					ret = true
				}
			}
		}
	}

	if len(tmp) > 0 {
		this.s += string(letter)
	} else {
		this.s = ""
	}

	this.words2 = tmp
	return ret
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
