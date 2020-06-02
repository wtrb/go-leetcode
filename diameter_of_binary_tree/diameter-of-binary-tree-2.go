package btree

func diameterOfBinaryTree2(root *TreeNode) int {
	var diameter = 0
	calculateMaxHeight(root, &diameter)
	return diameter
}

func calculateMaxHeight(node *TreeNode, diameter *int) int {
	var maxHeight = 0

	if node != nil {
		var maxHeightLeftNode = calculateMaxHeight(node.Left, diameter)
		var maxHeightRightNode = calculateMaxHeight(node.Right, diameter)

		*diameter = maxInts(*diameter, maxHeightLeftNode+maxHeightRightNode)
		maxHeight = maxInts(maxHeightLeftNode, maxHeightRightNode) + 1
	}

	return maxHeight
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Diameter of Binary Tree
// https://leetcode.com/problems/diameter-of-binary-tree
