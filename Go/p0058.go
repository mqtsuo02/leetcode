package leetcode

func lengthOfLastWord(s string) int {
	length, tail := 0, len(s)-1
	for tail >= 0 && s[tail:tail+1] == " " {
		tail--
	}
	for tail >= 0 && s[tail:tail+1] != " " {
		tail--
		length++
	}
	return length
}
