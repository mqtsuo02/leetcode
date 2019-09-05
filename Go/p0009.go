package leetcode

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revx := 0
	for x > revx {
		revx = revx*10 + x%10
		x /= 10
	}
	return x == revx || x == revx/10
}
