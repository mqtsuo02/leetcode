package leetcode

func generateParenthesis(n int) []string {
	ansp := &[]string{}
	writeParenthesis(0, 0, n, "", ansp)
	return *ansp
}

func writeParenthesis(open, close, max int, s string, ansp *[]string) {
	if open+close == max*2 {
		*ansp = append(*ansp, s)
		return
	}
	if open < max {
		writeParenthesis(open+1, close, max, s+"(", ansp)
	}
	if close < open {
		writeParenthesis(open, close+1, max, s+")", ansp)
	}
}
