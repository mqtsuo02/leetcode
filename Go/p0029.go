package leetcode

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	n := 0
	if 0 < dividend && divisor < 0 { // +1 -1
		for 0 <= dividend {
			dividend += divisor
			n--
		}
		return n + 1
	}
	if dividend < 0 && 0 < divisor { // -1 1
		for dividend <= 0 {
			dividend += divisor
			n--
		}
		return n + 1
	}
	if dividend < 0 && divisor < 0 { // -1 -1
		for dividend <= divisor {
			dividend -= divisor
			n++
		}
		return n
	}
	for divisor <= dividend {
		dividend -= divisor
		n++
	}
	return n
}
