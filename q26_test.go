package leetcode

import "testing"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	last := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == last {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
		last = num
	}

	return len(nums)
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	newNums := []int{} // 加一个新数组，放入不同德尔值
	for _, num := range nums {
		// 新数组为空的时候直接加入值就行
		// 新数组末尾的值和当前的值是不是相等 如果不相等那也插入新数组
		if len(newNums) == 0 || newNums[len(newNums)-1] != num {
			newNums = append(newNums, num)
		}
	}

	// 因为 题目要求改变 nums ，所以这里将 newNums 数据回写到 nums 中
	for idx, num := range newNums {
		nums[idx] = num
	}

	return len(newNums)
}

func removeDuplicates3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lens := 0 // 假如开了一个新数组，其实没有开，我们只是记录了下新数组的位置
	for i := 0; i < len(nums); i++ {
		// 新数组为 0 直接向右移动，向右移动表示占有这一段空间
		// 新数组最后一个值和 当前值进行比较，如果不相等，则将当前值移动到新数组后面
		if lens == 0 || nums[lens-1] != nums[i] {
			nums[lens], nums[i] = nums[i], nums[lens]
			lens++
		}
	}
	return lens // 直接返回长度即可
}

// func removeDuplicates(nums []int) int {
//     i := 0
//     for j := 1; j < len(nums); j++ {
//         if nums[i] != nums[j] {
//             i++
//             nums[i] = nums[j]
//         }
//     }
//     return i+1
// }

func TestRemoveDuplicates(t *testing.T) {
	// t.Log(removeDuplicates([]int{0, 0, 1, 1}))
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4}
	// t.Log(removeDuplicates(nums))
	// nn := len(nums) - 1
	// t.Log(nums[nn+1:], append(nums[:nn], nums[nn+1:]...))
}
