package pow

func myPow(x float64, n int) float64 {
	switch {
	case n == 0:
		return 1.0
	case n == 1:
		return x
	case n < 0:
		return 1 / myPow(x, -n)
	}

	result := 1.0
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}
	return result
}

// Pow(x, n)
// https://leetcode.com/problems/powx-n/

// https://leetcode.com/problems/powx-n/discuss/659996/Binary-exponential-indepth-explanation

/*
	a^n =
	1,					if n == 0
	a*a^(n-1),			if n is odd
	a^(n/2) * a^(n/2),	if n is even
*/

// tags: bit
