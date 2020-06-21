package bookcase

func minHeightShelves(books [][]int, shelfWidth int) int {
	// return recursion(books, 0, 0, 0, shelfWidth)

	memo := make([][]int, len(books))
	for i := 0; i < len(books); i++ {
		memo[i] = make([]int, shelfWidth+1)
	}
	return memoize(books, 0, 0, 0, shelfWidth, memo)

	// return dypr(books, shelfWidth)
}

func recursion(books [][]int, i int, maxHeightCurrLevel int, remainingWidth, shelfWidth int) int {
	if i == len(books) {
		return 0
	}

	// try placing this i-th book in new shelf
	minH := books[i][1] + recursion(books, i+1, books[i][1], shelfWidth-books[i][0], shelfWidth)

	// OR

	// if there is still room, try placing this i-th book in current shelf
	if books[i][0] <= remainingWidth {
		newMaxHeight := maxInts(maxHeightCurrLevel, books[i][1])
		minH = minInts(
			minH,
			newMaxHeight-maxHeightCurrLevel+
				recursion(books, i+1, newMaxHeight, remainingWidth-books[i][0], shelfWidth),
		)
	}

	return minH
}

func memoize(books [][]int, i int, maxHeightCurrLevel int, remainingWidth, shelfWidth int, memo [][]int) int {
	if i == len(books) {
		return 0
	}

	if memo[i][remainingWidth] != 0 {
		return memo[i][remainingWidth]
	}

	// try placing this i-th book in new shelf
	minH := books[i][1] + memoize(books, i+1, books[i][1], shelfWidth-books[i][0], shelfWidth, memo)

	// OR

	// if there is still room, try placing this i-th book in current shelf
	if books[i][0] <= remainingWidth {
		newMaxHeight := maxInts(maxHeightCurrLevel, books[i][1])
		minH = minInts(
			minH,
			newMaxHeight-maxHeightCurrLevel+
				memoize(books, i+1, newMaxHeight, remainingWidth-books[i][0], shelfWidth, memo),
		)
	}

	memo[i][remainingWidth] = minH

	return memo[i][remainingWidth]
}

func dypr(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = 1<<63 - 1
	}

	for i := 1; i < n+1; i++ {
		curHeight := 0
		curWidth := 0
		j := 1
		for j <= i {
			curHeight = maxInts(curHeight, books[i-j][1])
			curWidth += books[i-j][0]
			if curWidth > shelfWidth {
				break
			}

			dp[i] = minInts(dp[i], dp[i-j]+curHeight)
			j++
		}
	}

	return dp[n]
}

// func minHeightShelves(books [][]int, shelfWidth int) int {
// 	n := len(books)
// 	dp := make([]int, n+1)
// 	dp[0] = 0
// 	for i := 1; i < n+1; i++ {
// 		dp[i] = 1<<63 - 1
// 	}

// 	for i := 1; i < n+1; i++ {
// 		h := 0
// 		w := shelfWidth

// 		for j := i - 1; j >= 0; j-- {
// 			w -= books[j][0]
// 			h = maxInts(h, books[j][1])

// 			if w >= 0 {
// 				dp[i] = minInts(dp[i], dp[j]+h)
// 			}
// 		}
// 	}

// 	return dp[n]
// }

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Filling Bookcase Shelves
// https://leetcode.com/problems/filling-bookcase-shelves/

// https://leetcode.com/problems/filling-bookcase-shelves/discuss/341629/Intuitive-Java-Solution-With-Clear-Explanation
// https://leetcode.com/problems/filling-bookcase-shelves/discuss/600696/C%2B%2B-DP-Solution-with-very-good-explanation

// tags: dp, dynamic programming
