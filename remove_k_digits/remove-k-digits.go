package digits

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	stack := make([]rune, 0)

	for _, c := range num {
		for k > 0 && len(stack) > 0 && c < stack[len(stack)-1] {
			stack = stack[0 : len(stack)-1]
			k--
		}

		stack = append(stack, c)
	}

	stack = stack[:len(stack)-k] // increasing numbers. Ex: 112, 1234,...

	leadingZeros := 0
	for _, c := range stack {
		if c != '0' {
			break
		}
		leadingZeros++
	}

	if len(stack) == leadingZeros {
		return "0"

	} else {
		return string(stack[leadingZeros:])
	}
}

// Remove K Digits
// https://leetcode.com/problems/remove-k-digits/
// https://leetcode.com/submissions/detail/339185385/?from=/explore/featured/card/may-leetcoding-challenge/535/week-2-may-8th-may-14th/3328/

// tags: stack
