package leetcode

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	res := postorderTraversal(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}})
	if !reflect.DeepEqual(res, []int{3, 2, 1}) {
		t.Fatal("failed", res)
	}
}
