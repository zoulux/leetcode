package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func kthGrammar(N int, K int) int {

	if N == 1 {
		return 0

	}

	if N == 2 {
		if K == 1 {
			return 0
		}
		return 1
	}

	str := "01"

	for i := 3; i <= N; i++ {
		var sb strings.Builder

		sb.WriteString(str)
		for j := len(str) - 1; j >= 0; j-- {
			sb.WriteByte(str[j])
		}
		str = sb.String()
	}

	fmt.Println("N", N, str)

	return int(str[K-1] - '0')
}

func TestKthGrammar(t *testing.T) {
	// t.Log(kthGrammar(1, 1))
	// t.Log(kthGrammar(2, 1))
	// t.Log(kthGrammar(2, 2))
	// t.Log(kthGrammar(4, 5))
	// t.Log(kthGrammar(30, 434991989))

	t.Log(kthGrammar(1, 1))
	t.Log(kthGrammar(2, 1))
	t.Log(kthGrammar(3, 2))
	t.Log(kthGrammar(4, 5))
	t.Log(kthGrammar(5, 5))
	t.Log(kthGrammar(6, 5))
	t.Log(kthGrammar(7, 5))

	// 	30
	// 434991989
}

// 第一行: 0
// 第二行: 01
// 第三行: 0110
// 第四行: 01101001
// 第五行: 0110100110010110
// 第六行: 01101001100101101001011001101001
// 第七行: 0110100110010110100101100110100110010110011010010110100110010110
