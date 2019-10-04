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

func TestInorderTraversal(t *testing.T) {
	res := inorderTraversal(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}})
	if !reflect.DeepEqual(res, []int{1, 3, 2}) {
		t.Fatal("failed", res)
	}
}

func TestPostorderTraversal(t *testing.T) {
	res := postorderTraversal(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}})
	if !reflect.DeepEqual(res, []int{3, 2, 1}) {
		t.Fatal("failed", res)
	}
}
