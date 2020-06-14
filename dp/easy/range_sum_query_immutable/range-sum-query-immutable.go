package rangesum

// Space complexity: O(N)
type NumArray struct {
	sums []int
}

// Time complexity: O(N)
func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	return NumArray{
		sums: sums,
	}
}

// Time complexity: O(1)
func (this *NumArray) SumRange(i int, j int) int {
	return this.sums[j+1] - this.sums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

// Range Sum Query - Immutable
// https://leetcode.com/problems/range-sum-query-immutable/

// tags: prefix sums, sum of slice, sum of subarray, dp, dynamic programming
