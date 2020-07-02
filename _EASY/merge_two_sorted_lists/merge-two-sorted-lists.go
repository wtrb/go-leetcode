package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode
	var cur *ListNode

	if l1.Val < l2.Val {
		cur = l1
		l1 = l1.Next
	} else {
		cur = l2
		l2 = l2.Next
	}
	head = cur

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next

		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return head
}

// Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

// tags: linked list
