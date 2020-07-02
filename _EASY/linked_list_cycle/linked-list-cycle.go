package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// return useMap(head)
	return useTwoPointer(head)
}

// Time complexity: O(n)
// Space complexity: O(n)
func useMap(head *ListNode) bool {
	m := make(map[*ListNode]struct{})

	cur := head
	for cur != nil {
		if _, ok := m[cur]; ok {
			return true

		} else {
			m[cur] = struct{}{}
		}

		cur = cur.Next
	}

	return false
}

// Time complexity: O(n)
// Space complexity: O(1)
func useTwoPointer(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/

// tags: linked list, two pointer, map, cycle detection
