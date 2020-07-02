package roman

// Time complexity: O(n)
// Space complexity: O(1)
func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		v := dict[s[i]]
		for i > 0 && dict[s[i-1]] < dict[s[i]] {
			v -= dict[s[i-1]]
			i--
		}

		sum += v
	}
	return sum
}

// Roman to Integer
// https://leetcode.com/problems/roman-to-integer/

// tags: map, loop
