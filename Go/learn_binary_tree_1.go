package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// preorderTraversal itaration pattern
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
// preorderTraversal recursice pattern
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root != nil {
		recursiveForPreorderTraversal(root, &ans)
	}
	return ans
}
*/

func recursiveForPreorderTraversal(node *TreeNode, ansp *[]int) {
	*ansp = append(*ansp, node.Val)
	if node.Left != nil {
		recursiveForPreorderTraversal(node.Left, ansp)
	}
	if node.Right != nil {
		recursiveForPreorderTraversal(node.Right, ansp)
	}
}

/*
// inorderTraversal iteration pattern
1. currがnilになるまでstackにaddして、Leftに移動
2. nilだったらstackをpopしてoutputにaddしてRightに移動
3. 1に戻る
4. 最終的にstackがない状態でRightに移動後、currがnilの場合にfor文が終了する
*/
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
// inorderTraversal recursice pattern
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root != nil {
		recursiveForInorderTraversal(root, &ans)
	}
	return ans
}
*/

func recursiveForInorderTraversal(node *TreeNode, ansp *[]int) {
	if node.Left != nil {
		recursiveForInorderTraversal(node.Left, ansp)
	}
	*ansp = append(*ansp, node.Val)
	if node.Right != nil {
		recursiveForInorderTraversal(node.Right, ansp)
	}
}

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
