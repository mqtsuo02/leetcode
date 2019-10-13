package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	recursiveForCountUnivalSubtrees(root, &count)
	return count
}

func recursiveForCountUnivalSubtrees(tn *TreeNode, c *int) bool {
	l, r := true, true
	if tn.Left != nil {
		l = recursiveForCountUnivalSubtrees(tn.Left, c) && tn.Val == tn.Left.Val
	}

	if tn.Right != nil {
		r = recursiveForCountUnivalSubtrees(tn.Right, c) && tn.Val == tn.Right.Val
	}

	if l && r {
		*c++
		return true
	}
	return false
}
