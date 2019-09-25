package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	c := &ListNode{Val: 0, Next: head}
	l := 0
	for c.Next != nil {
		c = c.Next
		l++
	}
	ti := l - n
	ansHead := &ListNode{Val: 0, Next: nil}
	c = ansHead
	for i := 0; i < l; i++ {
		if i != ti {
			c.Next = &ListNode{Val: head.Val, Next: nil}
			c = c.Next
		}
		head = head.Next
	}
	return ansHead.Next
}
