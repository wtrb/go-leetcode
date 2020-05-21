package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time complexity: O(N) since we visit each node exactly once.
// Space complexity: O(N) since we keep up to the entire tree.
func isValidBST(root *TreeNode) bool {
	// return validate(root, math.MinInt64, math.MaxInt64)
	return validate(root, -1<<63, 1<<63-1)
}

func validate(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
}

// Validate Binary Search Tree

// https://leetcode.com/problems/validate-binary-search-tree/

// tags: bst, binary search tree, tree
