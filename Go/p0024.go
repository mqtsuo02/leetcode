package leetcode

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func swapPairs(head *ListNode) *ListNode {
	ansHead := &ListNode{Val: 0, Next: nil}
	cur := ansHead
	n := 0
	cache := 0
	for head != nil {
		if n%2 == 1 {
			cur.Next = &ListNode{Val: head.Val, Next: &ListNode{Val: cache, Next: nil}}
			cur = cur.Next.Next
			head = head.Next
			n++
			continue
		}
		if head.Next == nil {
			cur.Next = &ListNode{Val: head.Val, Next: nil}
			break
		}
		cache = head.Val
		head = head.Next
		n++
	}
	return ansHead.Next
}
