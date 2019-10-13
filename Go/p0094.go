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
func inorderTraversal(root *TreeNode) []int {
	curr, stack, output := root, []*TreeNode{}, []int{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		output = append(output, curr.Val)
		curr = curr.Right
	}
	return output
}

/*
1. currがnilになるまでstackにaddして、Leftに移動
2. nilだったらstackをpopしてoutputにaddしてRightに移動
3. 1に戻る
4. 最終的にstackがない状態でRightに移動後、currがnilの場合にfor文が終了する
*/

/*
// recursive pattern
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root != nil {
		recursiveForInorderTraversal(root, &ans)
	}
	return ans
}

func recursiveForInorderTraversal(node *TreeNode, ansp *[]int) {
	if node.Left != nil {
		recursiveForInorderTraversal(node.Left, ansp)
	}
	*ansp = append(*ansp, node.Val)
	if node.Right != nil {
		recursiveForInorderTraversal(node.Right, ansp)
	}
}
*/
