package bitree

import (
	"math"
)

var mps int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	mps = math.MinInt32

	pathSum(root)

	return mps
}

func pathSum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := maxInts(0, pathSum(node.Left))
	right := maxInts(0, pathSum(node.Right))

	mps = maxInts(mps, left+right+node.Val)

	return maxInts(left, right) + node.Val
}

func maxInts(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Binary Tree Maximum Path Sum
// https://leetcode.com/problems/binary-tree-maximum-path-sum/
// https://leetcode.com/submissions/detail/332067966/?from=/explore/featured/card/30-day-leetcoding-challenge/532/week-5/3314/

// tags: tree
