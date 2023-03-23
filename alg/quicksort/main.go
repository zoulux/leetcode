package main

import (
	"fmt"
	"math/rand"
)

func main() {

	arr := rand.Perm(5)
	fmt.Println(arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
func quickSort(arr []int, low, hight int) {
	if low >= hight {
		return
	}

	left, right := low, hight
	pivot := arr[(low+hight)/2] // 这里的经验值取的是中间数，经过 Benchmark 测试，确实比较优秀

	for left <= right {
		// 从左边开始迭代

		// 左边的数如果比 pivot 小，那么就应该将他放在左边，继续向右滑动，遇到一个比他大的为止
		for arr[left] < pivot {
			left++
		}

		// 右边的数如果比 pivot 大，那么就应该将他放在右边，继续向左滑动，遇到一个比他小的为止
		for arr[right] > pivot {
			right--
		}

		// 这里进行一次交换，将上面碰到的大数和小数交换一次
		//left 继续右走，right 继续左走 注意这里还不一定相遇，去继续执行上面的逻辑
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// 【 xxx[xxxxx]xxxxxx】
	// 【 xxxxxx][xxxxxxxx】
	// [] => left,right
	// 【】 => low,hight
	quickSort(arr, low, right)
	quickSort(arr, left, hight)
}
