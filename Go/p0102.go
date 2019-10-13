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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ts, output := []*TreeNode{root}, [][]int{}
	for len(ts) > 0 {
		tmpts := ts
		ts = []*TreeNode{}
		tmpns := []int{}
		for _, tmpt := range tmpts {
			tmpns = append(tmpns, tmpt.Val)
			if tmpt.Left != nil {
				ts = append(ts, tmpt.Left)
			}
			if tmpt.Right != nil {
				ts = append(ts, tmpt.Right)
			}
		}
		output = append(output, tmpns)
	}
	return output
}

/*
// recursive pattern
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root != nil {
		recursiveForLevelOrder(root, 0, &ans)
	}
	return ans
}

func recursiveForLevelOrder(node *TreeNode, level int, ansp *[][]int) {
	if len(*ansp)-1 < level {
		*ansp = append(*ansp, []int{})
	}
	(*ansp)[level] = append((*ansp)[level], node.Val)
	if node.Left != nil {
		recursiveForLevelOrder(node.Left, level+1, ansp)
	}
	if node.Right != nil {
		recursiveForLevelOrder(node.Right, level+1, ansp)
	}
}
*/
