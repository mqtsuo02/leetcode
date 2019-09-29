package leetcode

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hl, nl := len(haystack), len(needle)
	for n := 0; n < hl-nl+1; n++ {
		if haystack[n:n+nl] == needle {
			return n
		}
	}
	return -1
}
