package main

import "fmt"

func main() {
	fmt.Println(smallestBeautifulString("dbgh", 8) == "dbha")
	fmt.Println(smallestBeautifulString("abdc", 4) == "acba")
	fmt.Println(smallestBeautifulString("cd", 4) == "da")
	fmt.Println(smallestBeautifulString("ced", 6) == "cef")
	fmt.Println(smallestBeautifulString("abcz", 26) == "abda")
	fmt.Println(smallestBeautifulString("dc", 4) == "")
	fmt.Println(smallestBeautifulString("b", 6) == "c")
}

func smallestBeautifulString(s string, k int) string {
	bs := []byte(s)
	n := len(bs) - 1

	bs[n]++

	for i := n; i >= 0 && i <= n; {
		if bs[i] == byte('a'+k) {
			// 需要进位处理
			if i == 0 {
				// 前面也没法处理进位
				return ""
			}
			bs[i] = 'a' // 当前位置0
			bs[i-1]++   // 上一位 +1
			i--         // 现在处理上一位
		} else if (i-1 >= 0 && bs[i] == bs[i-1]) || (i-2 >= 0 && bs[i] == bs[i-2]) {
			// 不需要进位处理，但是需要判断回文
			// 如果有回文，只能将当前的值 +1，这时候不能移动 i，因为 +1 之后可能还是回文，也可能需要进位
			bs[i]++ //
		} else {
			// 不需要进位处理,也没有回文，说明这一位暂时是稳住了，可以向后走了
			i++
		}
	}
	return string(bs)
}

func h(s []byte) bool {
	if len(s) == 1 {
		return false
	}
	//所以只需要保证答案不包含长度为 2 或者长度为 3 的回文串。

	for i := len(s) - 1; i-1 >= 0; i-- {
		if s[i] == s[i-1] {
			return true
		}

		if i-2 >= 0 && s[i] == s[i-2] {
			return true
		}
	}

	return false
}

func addOne(bs []byte, end byte) []byte {
	if len(bs) == 0 {
		return nil
	}
	n := len(bs) - 1
	bs[n]++
	if bs[n] >= 'a' && bs[n] < end {
		return bs
	}

	if addOne(bs[:n], end) == nil {
		return nil
	}
	bs[n] = 'a'
	return bs
}

func smallestBeautifulStringErr2(s string, k int) string {

	end := byte('a' + k)
	// 处理 s 是不是回文
	bs := []byte(s)
	for {
		bs = addOne(bs, end)
		if bs == nil {
			return ""
		}

		if !h(bs) {
			return string(bs)
		}
	}
	return ""
}

func smallestBeautifulStringErr(s string, k int) string {

	// 处理 s 是不是回文
	h := func(s []byte) bool {
		if len(s) == 1 {
			return false
		}
		dp := make([][]bool, len(s))
		for i := range dp {
			dp[i] = make([]bool, len(s))
			dp[i][i] = true
		}

		for i := len(s) - 1; i >= 0; i-- {
			for j := i + 1; j < len(s); j++ {
				if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
					return true
				}
			}
		}

		return false
	}

	end := byte('a' + k - 1)
	bs := []byte(s)

	var dfs func(bs []byte) string
	dfs = func(bs []byte) string {
		n := len(bs) - 1
		for j := bs[n]; j <= end; j++ {
			bs[n] = j // 最后一位不断+1 测试
			if !h(bs) {
				return string(bs) // 找到非回文的，那就返回 bs
			}
		}

		bs[n] = end // 最后一位等于k的最大值
		for n-1 >= 0 && bs[n] >= end {
			bs[n-1] = bs[n-1] + 1 // 如果最后一位等于 end 了，那么前一位 +1
			bs[n] = 'a'           // 当前位置可以调整到第 0 位，因为前面的 j是 +1，所以这里 -1 处理一下
			n--                   // 继续向前处理 n-1 位
		}

		if bs[len(bs)-1] >= end || bs[0] > end { // 到最终情况了只能返回 空
			return ""
		}

		return dfs(bs) // 继续找 bs
	}

	bs[len(bs)-1]++
	return dfs(bs)
}
