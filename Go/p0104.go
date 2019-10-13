package leetcode

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
