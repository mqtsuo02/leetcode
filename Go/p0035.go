package leetcode

func searchInsert(nums []int, target int) int {
	li, hi := 0, len(nums)-1
	for li <= hi {
		mi := (li + hi) / 2
		if nums[mi] == target {
			return mi
		}
		if target < nums[mi] {
			hi = mi - 1
			continue
		}
		li = mi + 1
	}
	return li
}
