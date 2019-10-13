package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// iteration pattern
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, output := []*TreeNode{root}, []int{}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if len(output) == 0 {
			output = append(output, curr.Val)
			continue
		}
		output = append(output[0:1], output...)
		output[0] = curr.Val
	}
	return output
}

/*
// recursive pattern
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root != nil {
		recursiveForPostorderTraversal(root, &ans)
	}
	return ans
}

func recursiveForPostorderTraversal(node *TreeNode, ansp *[]int) {
	if node.Left != nil {
		recursiveForPostorderTraversal(node.Left, ansp)
	}
	if node.Right != nil {
		recursiveForPostorderTraversal(node.Right, ansp)
	}
	*ansp = append(*ansp, node.Val)
}
*/
