package leetcode

import "strconv"

func countAndSay(n int) string {
	output := "1"
	for i := 1; i < n; i++ {
		output = readAndSay(output)
	}
	return output
}

func readAndSay(in string) string {
	out := ""
	cache, count := in[0], 1
	for n := 1; n < len(in); n++ {
		if cache == in[n] {
			count++
			continue
		}
		out += (strconv.Itoa(count) + string(cache))
		cache = in[n]
		count = 1
	}
	out += (strconv.Itoa(count) + string(cache))
	return out
}
