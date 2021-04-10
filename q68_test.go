package leetcode

import (
	"fmt"
	"testing"
)

func fullJustify(words []string, maxWidth int) (ans []string) {

	n := len(words)
	idx := 0
	for idx < n {

		tmp := 0
		end := 0
		for i := idx; i < n; i++ {
			end = i
			tmp += len(words[i])
			tmp += 1
			if tmp > maxWidth {
				break
			}
		}

		fmt.Println(idx, end-1)
		idx = end
	}

	return
}

func TestFullJustify(t *testing.T) {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}

	t.Log(fullJustify(words, 16))

}
