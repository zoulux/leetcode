package main

func main() {

}

func minimumOperations(nums []int) int {
	mm := make(map[int]struct{})
	for _, v := range nums {
		if v > 0 {
			mm[v] = struct{}{}
		}
	}
	return len(mm)
}
