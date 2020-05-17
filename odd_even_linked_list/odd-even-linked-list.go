package linkedlist

// ListNode defines singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(N)
// Space complexity: O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	evenHead := head.Next
	odd := head
	even := head.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}

// https://leetcode.com/problems/odd-even-linked-list/
// https://leetcode.com/submissions/detail/340589998/?from=/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3331/

// tags: linked list
