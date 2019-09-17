package leetcode

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	rs := []rune(s)
	if len(rs) < 2 {
		return len(rs)
	}
	answer := 0
	count := 0
	m := make(map[rune]int)
	for i, r := range rs {
		if _, ok := m[r]; !ok {
			count++
		} else {
			count = int(math.Min(float64(i-m[r]), float64(count+1)))
		}
		answer = int(math.Max(float64(answer), float64(count)))
		m[r] = i
	}
	return answer
}
