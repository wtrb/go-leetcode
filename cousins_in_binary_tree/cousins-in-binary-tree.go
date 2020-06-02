package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	queue := []*TreeNode{root}
	genPopCounter := 1

	for len(queue) != 0 {
		if genPopCounter == 0 {
			genPopCounter = len(queue)
		}

		node := queue[0]
		queue = queue[1:]
		genPopCounter--

		target := -1
		if node.Left != nil {
			if node.Left.Val == x {
				target = y

			} else if node.Left.Val == y {
				target = x

			} else {
				queue = append(queue, node.Left)
			}
		}
		if node.Right != nil {
			// if node.Right.Val == target { //sibling check
			// 	return false
			// }

			if node.Right.Val == x {
				target = y

			} else if node.Right.Val == y {
				target = x

			} else {
				queue = append(queue, node.Right)
			}
		}
		if target != -1 {
			for genPopCounter > 0 { // search for the cousin
				node := queue[0]
				queue = queue[1:]
				genPopCounter--

				if node.Left != nil && node.Left.Val == target {
					return true
				}
				if node.Right != nil && node.Right.Val == target {
					return true
				}
			}
			return false
		}
	}

	return false
}

// Cousins in Binary Tree
// https://leetcode.com/problems/cousins-in-binary-tree/
// https://leetcode.com/submissions/detail/338066276/?from=/explore/featured/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3322/

// https://www.youtube.com/watch?v=nA3Fa05zIfg

// tags: tree, bfs
