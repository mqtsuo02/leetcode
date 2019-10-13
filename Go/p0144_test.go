package leetcode

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	res := preorderTraversal(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}})
	if !reflect.DeepEqual(res, []int{1, 2, 3}) {
		t.Fatal("failed", res)
	}
}
