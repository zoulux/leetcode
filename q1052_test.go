package leetcode

import "testing"

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)

	leftSum, winSum, rightSum := 0, 0, 0

	for i := X; i < n; i++ {
		if grumpy[i] == 0 {
			rightSum += customers[i]
		}
	}

	for i := 0; i < X; i++ {
		winSum += customers[i]
	}

	left, right := 0, X

	maxSum := leftSum + winSum + rightSum

	for right < n {
		if grumpy[left] == 0 {
			leftSum += customers[left]
		}

		if grumpy[right] == 0 {
			rightSum -= customers[right]
		}
		winSum = winSum - customers[left] + customers[right]

		tmp := leftSum + winSum + rightSum

		if tmp > maxSum {
			maxSum = tmp
		}
		left++
		right++
	}

	return maxSum
}

func TestMaxSatisfied(t *testing.T) {
	t.Log(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}
