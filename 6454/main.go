package main

func main() {

}

func makeSmallestPalindrome(s string) string {
	bs := []byte(s)
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if bs[left] != bs[right] {
			if bs[left] < bs[right] {
				bs[right] = bs[left]
			} else {
				bs[left] = bs[right]
			}
		}
	}
	return string(bs)
}
