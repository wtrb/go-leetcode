package twosum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	return useTraverse(root, k, m)

	// return useStack(root, k)
}

// Time complexity: O(n)
// Space complexity: O(n)
func useTraverse(root *TreeNode, k int, m map[int]struct{}) bool {
	if root == nil {
		return false
	}

	if _, ok := m[k-root.Val]; ok && root.Val != k-root.Val {
		return true

	} else {
		m[root.Val] = struct{}{}
	}

	return useTraverse(root.Left, k, m) || useTraverse(root.Right, k, m)
}

// Time complexity: O(n)
// Space complexity: O(n)
func useStack(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if _, ok := m[k-root.Val]; ok && root.Val != k-root.Val {
			return true

		} else {
			m[root.Val] = struct{}{}
		}

		root = root.Right
	}

	return false
}

// Two Sum IV - Input is a BST
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

// tags: bst, tree, stack, traverse
