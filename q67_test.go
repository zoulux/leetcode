package leetcode

import "testing"

func addBinary(a string, b string) string {

	na, nb := len(a), len(b)
	maxn := na
	if nb > maxn {
		maxn = nb
	}

	ans := []byte{'0'}
	flag := byte('0')
	for i := 0; i < maxn; i++ {

		sa := byte('0')
		if na-i-1 >= 0 {
			sa = a[na-i-1]
		}

		sb := byte('0')
		if nb-i-1 >= 0 {
			sb = b[nb-i-1]
		}

		tmp := (sa - '0') + (sb - '0') + (flag - '0')

		copy(ans[1:], ans[0:])

		if tmp == 0 || tmp == 1 {
			flag = byte('0')
			ans[0] = '0' + tmp
			// ans = append([]byte{'0' + tmp}, ans...)
		} else if tmp == 2 {
			flag = byte('1')

			ans[0] = '0'

			// ans = append([]byte{'0'}, ans...)
		} else if tmp == 3 {
			flag = byte('1')
			ans[0] = '1'
			// ans = append([]byte{'1'}, ans...)
		}
	}
	if flag == byte('1') {
		ans = append([]byte{'1'}, ans...)
	}
	return string(ans)
}

func TestAddBinary(t *testing.T) {
	t.Log(addBinary("0", "0"))
	t.Log(addBinary("1", "0"))
	t.Log(addBinary("1", "1"))
	t.Log(addBinary("11", "1"))
	t.Log(addBinary("1010", "1011"))
}
