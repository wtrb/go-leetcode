package robber

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time complexity: O()
// Space complexity: O()
func rob(root *TreeNode) int {
	// return recursion(root)

	return maxInts(robb(root))
}

func recursion(node *TreeNode) int {
	if node == nil {
		return 0
	}

	a := node.Val
	if node.Left != nil {
		a += recursion(node.Left.Left) + recursion(node.Left.Right)
	}
	if node.Right != nil {
		a += recursion(node.Right.Left) + recursion(node.Right.Right)
	}

	b := recursion(node.Left) + recursion(node.Right)

	return maxInts(a, b)
}

// Time complexity: O(2^N)
// Space complexity: O(1)
func robb(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	inclusive := node.Val
	exclusive := 0

	leftIn, leftEx := robb(node.Left)
	rightIn, rightEx := robb(node.Right)

	inclusive += leftEx + rightEx
	exclusive += maxInts(leftIn, leftEx) + maxInts(rightIn, rightEx)

	return inclusive, exclusive
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// House Robber II
// https://leetcode.com/problems/house-robber-ii/

// tags: dp, dynamic programming
