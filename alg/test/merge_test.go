package test

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMerge(t *testing.T) {

	arr := rand.Perm(10)
	t.Log(arr)
	arr = mergeSort(arr)
	t.Log(arr)

}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))

	for len(a) != 0 && len(b) != 0 {
		if a[0] < b[0] {
			result = append(result, a[0])
			a = a[1:]
		} else {
			result = append(result, b[0])
			b = b[1:]
		}
	}

	result = append(result, a...)
	result = append(result, b...)

	return result
}

func TestQuick(t *testing.T) {
	arr := []int{2, 2, 2}
	//arr := rand.Perm(10)
	t.Log(arr)
	quickSort(arr, 0, len(arr)-1)
	t.Log(arr)
}

func quickSort(arr []int, low, height int) {
	if low >= height {
		return
	}

	left := low
	right := height

	p := arr[(left+right)/2]
	for left <= right {
		for arr[left] < p {
			left++
		}

		for arr[right] > p {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			fmt.Println("---")
			left++
			right--
		}
	}
	quickSort(arr, low, right)
	quickSort(arr, left, height)
}
