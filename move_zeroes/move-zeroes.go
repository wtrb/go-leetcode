package move_zeroes_283

// Time Complexity: O(n^2)
// Space Complexity : O(1)
func moveZeroes1(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
			}
		}
	}
}

// Time Complexity: O(n)
// Space Complexity : O(1)
func moveZeroes2(nums []int) {
	zeroPtr := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zeroPtr], nums[i] = nums[i], nums[zeroPtr]

			zeroPtr++
		}
	}
}

// Time Complexity: O(n)
// Space Complexity : O(1)
func moveZeroes(nums []int) {
	zeroPtr := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zeroPtr] = nums[i]
			zeroPtr++
		}
	}

	for i := zeroPtr; i < len(nums); i++ {
		nums[i] = 0
	}
}

// Move Zeroes
// https://leetcode.com/problems/move-zeroes/

// tags: pointer
