package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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
