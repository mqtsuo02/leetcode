package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ansdiff := int(^uint(0) >> 1)
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			diff := nums[i] + nums[j] + nums[k] - target
			if math.Abs(float64(diff)) < math.Abs(float64(ansdiff)) {
				ansdiff = diff
			}
			if diff > 0 {
				k--
				continue
			}
			if diff < 0 {
				j++
				continue
			}
			return target
		}
	}
	return target + ansdiff
}
