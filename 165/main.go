package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("0.1", "1.1"))
}

// leetcode submit region begin(Prohibit modification and deletion)
func compareVersion(version1 string, version2 string) int {

	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	i := 0

	for {
		vi1, _ := strconv.Atoi(v1[i])
		vi2, _ := strconv.Atoi(v2[i])
		if vi1 < vi2 {
			return -1
		} else if vi1 > vi2 {
			return 1
		} else {
			i++
			if i == len(v1) && i == len(v2) {
				return 0
			}
			if i == len(v1) {
				v1 = append(v1, "0")
			}

			if i == len(v2) {
				v2 = append(v2, "0")
			}
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
