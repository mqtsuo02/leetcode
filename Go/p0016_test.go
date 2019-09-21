package leetcode

import (
	"testing"
)

func TestP0016(t *testing.T) {
	ans := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	if ans != 2 {
		t.Fatal("failed", ans)
	}
}
