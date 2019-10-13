package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"reflect"
	"strconv"
)

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lvs, rvs := []string{}, []string{}
	if root.Left != nil {
		recursiveLeftBranch(root.Left, &lvs)
	}
	if root.Right != nil {
		recursiveRightBranch(root.Right, &rvs)
	}
	return reflect.DeepEqual(lvs, rvs)
}

func recursiveLeftBranch(node *TreeNode, vsp *[]string) {
	*vsp = append(*vsp, strconv.Itoa(node.Val))
	if node.Left != nil {
		recursiveLeftBranch(node.Left, vsp)
	} else {
		*vsp = append(*vsp, "")
	}
	if node.Right != nil {
		recursiveLeftBranch(node.Right, vsp)
	} else {
		*vsp = append(*vsp, "")
	}
}

func recursiveRightBranch(node *TreeNode, vsp *[]string) {
	*vsp = append(*vsp, strconv.Itoa(node.Val))
	if node.Right != nil {
		recursiveRightBranch(node.Right, vsp)
	} else {
		*vsp = append(*vsp, "")
	}
	if node.Left != nil {
		recursiveRightBranch(node.Left, vsp)
	} else {
		*vsp = append(*vsp, "")
	}
}
