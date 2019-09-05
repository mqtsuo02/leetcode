package leetcode

import (
	"math"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbersWrong(l1 *ListNode, l2 *ListNode) *ListNode {
	return int2ListNode(listNode2Int(l1) + listNode2Int(l2))
}

func listNode2Int(l *ListNode) int {
	n := 0
	for i := 0; l != nil; i++ {
		n += l.Val * int(math.Pow(10, float64(i)))
		l = l.Next
	}
	return n
}

func int2ListNode(n int) *ListNode {
	if n == 0 {
		return &ListNode{Val: 0, Next: nil}
	}

	var l *ListNode
	rs := []rune(strconv.Itoa(n))

	for i := 0; i < len(rs); i++ {
		v, _ := strconv.Atoi(string(rs[i]))
		l = &ListNode{Val: v, Next: l}
	}
	return l
}
