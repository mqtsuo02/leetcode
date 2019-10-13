package leetcode

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	res := inorderTraversal(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}})
	if !reflect.DeepEqual(res, []int{1, 3, 2}) {
		t.Fatal("failed", res)
	}
}
