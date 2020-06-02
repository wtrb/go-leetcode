package binarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	right := len(preorder)

	for i := range preorder {
		if preorder[i] > val {
			right = i
			break
		}
	}

	return &TreeNode{
		Val:   val,
		Left:  bstFromPreorder(preorder[1:right]),
		Right: bstFromPreorder(preorder[right:]),
	}
}

// Construct Binary Search Tree from Preorder Traversal
// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal
// https://leetcode.com/submissions/detail/327941155/?from=/explore/featured/card/30-day-leetcoding-challenge/530/week-3/3305/

// tags: tree, binary search tree, bst, preorder traversal
