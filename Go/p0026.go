package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] == nums[j] {
			continue
		}
		i++
		nums[i] = nums[j]
	}
	nums = nums[:i+1]
	return i + 1
}
