package bst

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
	for i := 0; i < len(preorder); i++ {
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
// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/

// https://leetcode.com/submissions/detail/343920872/?from=/explore/featured/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3339/

// tags: bst, binary search tree, tree, traverse, preorder
