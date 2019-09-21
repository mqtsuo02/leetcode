package leetcode

import (
	"sort"
	"strconv"
)

func threeSumLow(nums []int) [][]int {
	ns := nums
	sort.Ints(ns)
	answer := [][]int{}
	length := len(ns)
	m := make(map[string]bool)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				if ns[i]+ns[j]+ns[k] == 0 {
					l := []int{nums[i], nums[j], nums[k]}
					ls := nsToString(l)
					if _, ok := m[ls]; !ok {
						answer = append(answer, l)
						m[ls] = true
					}
				}
			}
		}
	}
	return answer
}

func nsToString(ns []int) string {
	s := ""
	d := ""
	for _, n := range ns {
		s = s + strconv.Itoa(n) + d
		d = ","
	}
	return s
}
