package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(5 % -2)
	fmt.Println(baseNeg2(0))
	fmt.Println(baseNeg2(6))
	fmt.Println(baseNeg2(3))
	fmt.Println(baseNeg2(4))
}

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	var ans string
	for n != 0 {
		if n%2 == 0 {
			// 偶数
			ans += "0"
			n /= -2
		} else {
			ans += "1"
			if n < 0 {
				// -5/-2 = 3...1  => -2*3+1= -5 => (5+1)/2 => 3
				n = (-n + 1) / 2
			} else {
				n /= -2
			}
		}
	}
	return reverse(ans)
}
func reverse(ans string) string {
	tmp := []byte(ans)
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return string(tmp)
}

func baseNeg22(n int) string {

	if n == 0 {
		return "0"
	}
	var ans []byte
	k := 1
	for n != 0 {
		if n%2 != 0 {
			ans = append(ans, '1')
			n -= k
		} else {
			ans = append(ans, '0')
		}
		k *= -1
		n /= 2
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)

}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

//输入：n = 2
// 10
//输出："110"
//解释：(-2)2 + (-2)1 = 2

//输入：n = 3
// 11
//输出："111"
//解释：(-2)2 + (-2)1 + (-2)0 = 3

//输入：n = 4
// 100
//输出："100"
//解释：(-2)2 = 4

//输入：n = 5
// 101
//输出："101"
//解释：(-2)2 + (-2)0  = 4

//输入：n = 6
// 110
//输出："11010"
//解释：(-2)4 +(-2)3 + (-2)1  = 6
