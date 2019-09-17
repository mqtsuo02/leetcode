package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if mi, ok := m[n]; ok {
			return []int{mi, i}
		}
		m[target-n] = i
	}
	return []int{}
}
