package paint

func paintFence(n, k int) int {
	if n == 0 || k == 0 {
		return 0
	}

	if n == 1 {
		return k
	}

	same := k
	diff := k * (k - 1)
	for i := 3; i <= n; i++ {
		same, diff = diff*1, (same+diff)*(k-1)
	}

	return same + diff
}

// Paint fence
// https://leetcode.com/problems/paint-fence/

/*
There is a fence with n posts, each post can be painted with one of the k colors. You have to paint all the posts such that no more than two adjacent fence posts have the same color. Return the total number of ways you can paint the fence.
Note: n and k are non-negative integers.
*/

// https://www.youtube.com/watch?v=deh7UpSRaEY
/*
posts:	1	2			3
--
same:	k	k*1			prev diff * 1
diff:		k*(k-1)		prev same * (k-1) + prev diff * (k-1)
*/

// tags: dp, dynamic programming
