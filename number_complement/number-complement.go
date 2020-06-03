package bitwise

func findComplement(num int) int {
	result := 0
	power := 1

	for num > 0 {
		result += (num%2 ^ 1) * power

		power <<= 1 // power *= 2
		num >>= 1   // num /= 2
	}

	return result
}

// Number Complement
// https://leetcode.com/problems/number-complement/
// https://leetcode.com/problems/complement-of-base-10-integer/

// tags: bit manipulation
