package bst

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}

}

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

// tags: bst, binary search tree, tree, stack
