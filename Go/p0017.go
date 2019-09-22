package leetcode

var m = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var ans []string

func combine(s, ds string) {
	if len(ds) == 0 {
		ans = append(ans, s)
		return
	}
	as := m[ds[:1]]
	for i := 0; i < len(as); i++ {
		combine(s+as[i:i+1], ds[1:])
	}
}

func letterCombinations(digits string) []string {
	ans = []string{}
	if len(digits) > 0 {
		combine("", digits)
	}
	return ans
}
