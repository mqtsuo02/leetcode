package leetcode

var phone = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	ansp := &[]string{}
	if len(digits) > 0 {
		combine("", digits, ansp)
	}
	return *ansp
}

func combine(s, ds string, ansp *[]string) {
	if len(ds) == 0 {
		*ansp = append(*ansp, s)
		return
	}
	as := phone[ds[:1]]
	for i := 0; i < len(as); i++ {
		combine(s+as[i:i+1], ds[1:], ansp)
	}
}
