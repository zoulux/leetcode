package main

func main() {

}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	mc := make(map[int]bool)

	for _, v := range m {
		if _, ok := mc[v]; ok {
			return false
		}
		mc[v] = true
	}
	return true
}
