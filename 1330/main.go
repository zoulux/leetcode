package main

import (
	"fmt"
	"math"
)

// 输入：nums = [2,3,1,5,4]
// 输出：10
// 解释：通过翻转子数组 [3,1,5] ，数组变成 [2,5,1,3,4] ，数组值为 10 。

func main() {

	fmt.Println(maxValueAfterReverse([]int{
		2, 3, 1, 5, 4,
	}))
}

func maxValueAfterReverse(nums []int) int {
	n := len(nums)
	a, b := nums[0], nums[n-1]
	premax1 := math.MinInt / 2 //  ai + ai-1 -|ai - ai-1|  前缀和
	premax2 := math.MinInt / 2 //  ai - ai-1 -|ai - ai-1|  前缀和
	premax3 := math.MinInt / 2 // -ai + ai-1 -|ai - ai-1|  前缀和
	premax4 := math.MinInt / 2 // -ai - ai-1 -|ai - ai-1|  前缀和

	sum, ans := 0, 0 // sum 为原来的和， ans 为增量
	for i := 1; i < n; i++ {
		y, x := nums[i-1], nums[i] // 可以表示 ai-1(y) , ai(x) 和 aj(y)、aj+1(x)
		d := abs(x - y)
		sum += d

		ans = max(ans, max(
			abs(y-b),    //  abs(ai-1 - an-1) => abs(y-b)
			abs(x-a),    //  abs(ai - a0 )    => abs(x-a)
			premax1-x-y, // -aj+1 -aj -|aj+1 - aj| + ai + ai-1 -|ai - ai-1|
			premax2-x+y, // -aj+1 +aj -|aj+1 - aj| + ai - ai-1 -|ai - ai-1|
			premax3+x-y, //  aj+1 -aj -|aj+1 - aj| + -ai + ai-1 -|ai - ai-1|
			premax4+x+y, //  aj+1 +aj -|aj+1 - aj| + -ai - ai-1 -|ai - ai-1|
		)-d)

		premax1 = max(premax1, x+y-d)
		premax2 = max(premax2, x-y-d)
		premax3 = max(premax3, -x+y-d)
		premax4 = max(premax4, -x-y-d)
	}

	return sum + ans
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(arr ...int) int {
	x := arr[0]
	for _, v := range arr[1:] {
		if v > x {
			x = v
		}
	}
	return x
}
