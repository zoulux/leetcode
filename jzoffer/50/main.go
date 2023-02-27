package main

func main() {

}

func firstUniqChar(s string) byte {
	mm := make(map[rune]int)
	for _, v := range s {
		mm[v]++
	}

	for _, v := range s {
		if mm[v] == 1 {
			return byte(v)
		}
	}
	return ' '
}
