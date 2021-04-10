package leetcode

import (
	"reflect"
	"testing"
)

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordCount := len(words)
	wordsTotalLen := wordLen * wordCount

	mm := make(map[string]int, wordCount)
	for i := 0; i < len(words); i++ {
		mm[words[i]]++
	}

	ans := []int{}
	for i := 0; i < len(s)-wordsTotalLen+1; i++ {
		mmTmp := make(map[string]int, len(mm))

		breakFlag := false
		for j := i; j < i+wordsTotalLen; j += wordLen {
			subsub := s[j : j+wordLen]
			if _, ok := mm[subsub]; !ok {
				breakFlag = true
				break
			} else {
				mmTmp[subsub]++
				if mmTmp[subsub] > mm[subsub] {
					breakFlag = true
					break
				}
			}
		}

		if !breakFlag {
			if !reflect.DeepEqual(mm, mmTmp) {
				breakFlag = true
			}
		}

		if !breakFlag {
			ans = append(ans, i)
		}
	}
	return ans
}

func TestFindSubstring(t *testing.T) {
	t.Log(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	t.Log(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	t.Log(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	t.Log(findSubstring("aaaaaaaaaaaaaa", []string{"aa", "aa"}))

}
