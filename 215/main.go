package main

import (
	"fmt"
)

func main() {

	//fmt.Println(findKthLargest([]int{
	//	//3, 2, 1, 5, 6, 4,
	//	//3, 2, 3, 1, 2, 4, 5, 5, 6,
	//	-1, 2, 0,
	//}, 3))

	//fmt.Println(findKthLargest([]int{3, 1, 2, 4}, 2) == 3)
	//fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 2) == 10)
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 20) == 2)
	//fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2) == 5)
	//	[3 2 1 5 6 4]
	//	[1 2 3 5 6 4]
	//	[1 2 3 4 6 5]
	//	[1 2 3 4 5 6]

	//fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 3))
	//
	//fmt.Println(findKthLargest([]int{-1, 2, 0}, 3))
	//
	//fmt.Println(findKthLargest([]int{
	//	5, 2, 4, 1, 3, 6, 0,
	//}, 4))

}

func findKthLargest(nums []int, k int) int {

	var quick func(nums []int, k int) int
	quick = func(nums []int, k int) int {
		idx := (len(nums)) / 2
		provit := nums[idx]
		larr, rarr := make([]int, 0), make([]int, 0)

		for i, v := range nums {
			if i == idx {
				continue
			}
			if v < provit {
				larr = append(larr, v)
			} else {
				rarr = append(rarr, v)
			}
		}

		if len(larr) >= k {
			return quick(larr, k)
		} else if len(larr) == k-1 {
			return provit
		} else {
			return quick(rarr, k-len(larr)-1)
		}

	}
	return quick(nums, len(nums)-k)
}

func findKthLargest4(nums []int, k int) int {

	var quick func(low, height, k int) int

	quick = func(l, r, k int) int {
		fmt.Println(nums)
		if l >= r {
			return nums[k]
		}

		i, j := l, r
		//i, j := l-1, r+1
		p := nums[(i+j)/2]

		for i <= j {
			//i++
			//j--
			for nums[i] < p {
				i++
			}
			for nums[j] > p {
				j--
			}

			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}

		if i-1 == j+1 {
			j++
		}

		if j < k {
			return quick(j+1, r, k)
		}
		return quick(l, j, k)
	}
	return quick(0, len(nums)-1, len(nums)-k)
}

func findKthLargest3(nums []int, k int) int {

	var quick func(low, height, k int) int

	quick = func(l, r, k int) int {
		if l >= r {
			return nums[k]
		}

		i, j := l-1, r+1
		p := nums[(i+j)/2]

		for i < j {
			i++
			j--
			for nums[i] < p {
				i++
			}
			for nums[j] > p {
				j--
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}

		if j < k {
			return quick(j+1, r, k)
		}
		return quick(l, j, k)
	}
	return quick(0, len(nums)-1, len(nums)-k)
}

func findKthLargest2(nums []int, k int) int {

	list := make([]int, 0)
	for j := 0; j < len(nums); j++ {
		for len(list) >= k && nums[j] > list[0] {
			list = list[1:]
		}

		i := 0
		for i < len(list) && nums[j] > list[i] {
			i++
		}

		l1 := list[:i]
		l2 := list[i:]

		list = append([]int{}, l1...)
		list = append(list, nums[j])
		list = append(list, l2...)

		if len(list) >= k {
			list = list[len(list)-k:]
		}
	}

	return list[0]
}
