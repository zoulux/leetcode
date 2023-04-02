package main

import "fmt"

func main() {

	//fmt.Println(addStrings("11", "123"))
	//fmt.Println(addStrings("456", "77"))
	//fmt.Println(addStrings("0", "0"))
	//fmt.Println(addStrings("9133", "0"))
	fmt.Println(addStrings("408", "5"))
}

// leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	if num2 == "0" {
		return num1
	}

	if num1 == "0" {
		return num2
	}

	s1 := len(num1) - 1
	s2 := len(num2) - 1

	e := uint8(0)
	ans := ""
	for s1 >= 0 && s2 >= 0 {
		x := num1[s1] - '0' + num2[s2] - '0' + e
		t := x % 10
		e = x / 10
		ans = string(t+'0') + ans
		s1--
		s2--
	}

	if s1 >= 0 {
		ans = addStrings(string(e+'0'), num1[:s1+1]) + ans
	} else if s2 >= 0 {
		ans = addStrings(string(e+'0'), num2[:s2+1]) + ans
	} else if e != 0 {
		ans = string(e+'0') + ans
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
