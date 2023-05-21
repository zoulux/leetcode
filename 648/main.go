package main

import (
	"sort"
	"strings"
)

func main() {

}

func replaceWords(dictionary []string, sentence string) string {

	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	var ans []string
	arr := strings.Split(sentence, " ")
	for _, s := range arr {
		pre := ""
		for _, p := range dictionary {
			if strings.HasPrefix(s, p) {
				pre = p
				break
			}
		}
		if pre != "" {
			ans = append(ans, pre)
		} else {
			ans = append(ans, s)
		}
	}
	return strings.Join(ans, " ")
}
