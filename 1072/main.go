package main

import (
	"fmt"
	"strings"
)

type a interface {
	aa()
}

func handle(a a) {
	a.aa()
}

type bb struct {
}

func (b bb) aa() {
	fmt.Println("bb")
}

func main() {
	var b *bb
	//b.aa()

	handle(b)

	//var d int = 2
	//fmt.Printf("%08b\n", d)
	//fmt.Printf("%08b\n", ^d)
	//fmt.Printf("%08b\n", ^0) // 11111101
	//fmt.Println(maxEqualRowsAfterFlips([][]int{
	//	{0, 0, 0},
	//	{0, 0, 1},
	//	{1, 1, 0},
	//}))

	//fmt.Println(0111 ^ 0111)
	//fmt.Println(0111 | 0111)
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	mm := make(map[string]int)
	for i := range matrix {
		var s []byte
		for _, v := range matrix[i] {
			s = append(s, byte('0'+1^v^matrix[i][0]))
		}
		mm[string(s)]++
	}

	max := 0
	for _, v := range mm {
		if v > max {
			max = v
		}
	}
	return max
}

func maxEqualRowsAfterFlips3(matrix [][]int) int {
	max := 0
	mm := make(map[string]int)
	for i := range matrix {
		var s []byte
		var rs []byte

		for _, v := range matrix[i] {
			s = append(s, byte('0'+v))
			rs = append(rs, byte('0'+1-v))
		}

		mm[string(s)]++
		mm[string(rs)]++

		if mm[string(s)] > max {
			max = mm[string(s)]
		}
	}
	return max
}

func maxEqualRowsAfterFlips2(matrix [][]int) int {

	max := 0
	mm := make(map[string]int)
	for i := range matrix {
		firstZero := false
		if matrix[i][0] == 0 {
			firstZero = true
		}

		var sb strings.Builder
		//sb.Write([]byte(matrix[i]))
		for j := range matrix[i] {
			if firstZero {
				sb.WriteByte(byte('0' + matrix[i][j]))
			} else {
				sb.WriteByte(byte('0' + matrix[i][j] ^ 1))
			}
		}

		s := sb.String()
		mm[s]++

		if mm[s] > max {
			max = mm[s]
		}
	}
	return max
}
