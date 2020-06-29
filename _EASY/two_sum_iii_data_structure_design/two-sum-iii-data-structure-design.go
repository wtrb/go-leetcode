package twosum

type TwoSum struct {
	data map[int]struct{}
}

func New() *TwoSum {
	return &TwoSum{
		data: make(map[int]struct{}),
	}
}

// Add the number to the internal data structure.
// Time complexity: O(1)
func (t *TwoSum) Add(num int) {
	t.data[num] = struct{}{}
}

// Find if there exists any pair of numbers which sum is equal to the value.
// Time complexity: O(n)
func (t *TwoSum) Find(target int) bool {
	for x := range t.data {
		if _, ok := t.data[target-x]; ok && x != target-x {
			return true
		}
	}
	return false
}

// Two Sum III - Data structure design
// https://leetcode.com/problems/two-sum-iii-data-structure-design/

/*
	Design and implement a TwoSum class. It should support the following operations: add and find.
		add - Add the number to an internal data structure.
		find - Find if there exists any pair of numbers which sum is equal to the value.

	Example 1:
		add(1); add(3); add(5);
		find(4) -> true
		find(7) -> false
*/

// tags: map
