package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ansHead := &ListNode{Val: 0, Next: nil}
	c := ansHead
	for l1 != nil || l2 != nil {
		if l1 == nil {
			c.Next = &ListNode{Val: l2.Val, Next: nil}
			c = c.Next
			l2 = l2.Next
			continue
		}
		if l2 == nil {
			c.Next = &ListNode{Val: l1.Val, Next: nil}
			c = c.Next
			l1 = l1.Next
			continue
		}
		if l1.Val < l2.Val {
			c.Next = &ListNode{Val: l1.Val, Next: nil}
			c = c.Next
			l1 = l1.Next
			continue
		}
		c.Next = &ListNode{Val: l2.Val, Next: nil}
		c = c.Next
		l2 = l2.Next
	}
	return ansHead.Next
}
