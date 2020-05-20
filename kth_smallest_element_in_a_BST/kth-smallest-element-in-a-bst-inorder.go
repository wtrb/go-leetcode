package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time complexity: O(N) to build a traversal.
// Space complexity: O(N) to keep an inorder traversal.
func kthSmallestInorder(root *TreeNode, k int) int {
	arr := []int{}
	inorder(root, &arr)
	return arr[k-1]
}

func inorder(node *TreeNode, arr *[]int) {
	if node.Left != nil {
		inorder(node.Left, arr)
	}
	*arr = append(*arr, node.Val)
	if node.Right != nil {
		inorder(node.Right, arr)
	}
}

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

// tags: bst, binary search tree, tree, in-order traversal
