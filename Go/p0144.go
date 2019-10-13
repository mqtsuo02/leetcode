package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// itaration pattern
func preorderTraversal(root *TreeNode) []int {
	stack, output := []*TreeNode{root}, []int{}
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root != nil {
			output = append(output, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
		}
	}
	return output
}

/*
// recursive pattern
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root != nil {
		recursiveForPreorderTraversal(root, &ans)
	}
	return ans
}

func recursiveForPreorderTraversal(node *TreeNode, ansp *[]int) {
	*ansp = append(*ansp, node.Val)
	if node.Left != nil {
		recursiveForPreorderTraversal(node.Left, ansp)
	}
	if node.Right != nil {
		recursiveForPreorderTraversal(node.Right, ansp)
	}
}
*/
