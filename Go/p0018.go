package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	ansp := &[][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSumForFourSum(nums[i], nums[i+1:], target, ansp)
	}
	return *ansp
}

func threeSumForFourSum(n int, ns []int, t int, ansp *[][]int) {
	for j := 0; j < len(ns)-2; j++ {
		if j > 0 && ns[j] == ns[j-1] {
			continue
		}
		k, l := j+1, len(ns)-1
		for k < l {
			if k > j+1 && ns[k] == ns[k-1] {
				k++
				continue
			}
			if l < len(ns)-1 && ns[l] == ns[l+1] {
				l--
				continue
			}
			diff := n + ns[j] + ns[k] + ns[l] - t
			if diff < 0 {
				k++
				continue
			}
			if diff > 0 {
				l--
				continue
			}
			*ansp = append(*ansp, []int{n, ns[j], ns[k], ns[l]})
			k++
			l--
		}
	}
}
