package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//输入：s = "25525511135" 输出：["255.255.11.135","255.255.111.35"]

	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("101023"))

}

// leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {

	var ans []string
	var dfs func([]string, int)
	dfs = func(res []string, start int) {
		if start > len(s) {
			return
		}

		if len(res) == 4 && start == len(s) {
			ans = append(ans, strings.Join(res, "."))
			return
		}

		for ln := 1; ln <= 3; ln++ {
			if start+ln > len(s) {
				return
			}
			t := s[start : start+ln]
			if ti, _ := strconv.Atoi(t); ti > 255 {
				return
			}
			if len(t) > 1 && strings.HasPrefix(t, "0") {
				return
			}
			dfs(append(res, t), start+ln)
		}
	}

	dfs([]string{}, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
