package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	ns := nums
	sort.Ints(ns)
	answer := [][]int{}
	length := len(ns)
	for i := 0; i < length-2; i++ {
		if i > 0 && ns[i] == ns[i-1] {
			continue
		}
		j, k := i+1, length-1
		for j < k {
			if j > i+1 && ns[j] == ns[j-1] {
				j++
				continue
			}
			if k < length-1 && ns[k] == ns[k+1] {
				k--
				continue
			}
			sum := ns[i] + ns[j] + ns[k]
			if sum == 0 {
				answer = append(answer, []int{ns[i], ns[j], ns[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return answer
}
