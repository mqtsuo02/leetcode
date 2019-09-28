package leetcode

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] == val {
			continue
		}
		nums[i] = nums[j]
		i++
	}
	nums = nums[:i]
	return i
}
