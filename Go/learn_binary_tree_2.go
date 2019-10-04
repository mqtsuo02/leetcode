package leetcode

import (
	"reflect"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	recursiveForMaxDepth(root, 1, &ans)
	return ans
}

func recursiveForMaxDepth(node *TreeNode, depth int, ansp *int) {
	if *ansp < depth {
		*ansp = depth
	}
	if node.Left != nil {
		recursiveForMaxDepth(node.Left, depth+1, ansp)
	}
	if node.Right != nil {
		recursiveForMaxDepth(node.Right, depth+1, ansp)
	}
}

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

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	ans := false
	recursiveForHasPathSum(root, 0, sum, &ans)
	return ans
}

func recursiveForHasPathSum(node *TreeNode, n, sum int, ansp *bool) {
	if node.Left == nil && node.Right == nil {
		if n+node.Val == sum {
			*ansp = true
		}
		return
	}
	if node.Left != nil {
		recursiveForHasPathSum(node.Left, n+node.Val, sum, ansp)
	}
	if node.Right != nil {
		recursiveForHasPathSum(node.Right, n+node.Val, sum, ansp)
	}
}
