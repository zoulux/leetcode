package main

import "strings"

func main() {

}

func mergeAlternately(word1 string, word2 string) string {
	if word1 == "" {
		return word2
	}
	if word2 == "" {
		return word1
	}

	var ans strings.Builder
	ans.WriteByte(word1[0])
	ans.WriteByte(word2[0])
	ans.WriteString(mergeAlternately(word1[1:], word2[1:]))
	return ans.String()
}
