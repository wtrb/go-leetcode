package parentheses

// Time complexity: O(n)
// Space complexity: O(n)
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	closes := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}
	for _, c := range s {
		if open, ok := closes[c]; ok {
			if len(stack) == 0 || open != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, c)
		}
	}

	// for _, c := range s {
	// 	switch c {
	// 	case '(', '{', '[':
	// 		stack = append(stack, c)

	// 	case ')', '}', ']':
	// 		if len(stack) == 0 {
	// 			return false
	// 		}

	// 		pair := string([]rune{stack[len(stack)-1], c})
	// 		if pair == "()" || pair == "{}" || pair == "[]" {
	// 			stack = stack[:len(stack)-1]

	// 		} else {
	// 			return false
	// 		}

	// 	default:
	// 		return false
	// 	}
	// }

	return len(stack) == 0
}

// Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

// tags: stack, map
