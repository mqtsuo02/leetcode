package leetcode

import "sort"

func nextPermutation(nums []int) {
	n := len(nums) - 1

	if n == 0 {
		return
	}

	for 0 < n {
		if nums[n-1] < nums[n] {
			break
		}
		n--
	}

	sort.Ints(nums[n:])

	if n < 1 {
		return
	}

	for j := n; j < len(nums); j++ {
		if nums[n-1] < nums[j] {
			nums[n-1], nums[j] = nums[j], nums[n-1]
			return
		}
	}
}
