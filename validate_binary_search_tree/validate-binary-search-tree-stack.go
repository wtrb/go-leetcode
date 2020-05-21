package tree

// Time complexity: O(N).
// Space complexity: O(N) to keep stack.
func isValidBSTStack(root *TreeNode) bool {
	stack := []*TreeNode{}
	lastLeftVal := -1 << 63

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= lastLeftVal {
			return false
		}
		lastLeftVal = root.Val

		root = root.Right
	}

	return true
}

// Validate Binary Search Tree

// https://leetcode.com/problems/validate-binary-search-tree/

// tags: bst, binary search tree, tree, stack
