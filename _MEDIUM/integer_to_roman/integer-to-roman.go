package roman

import (
	"strings"
)

func intToRoman(num int) string {
	values := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	dict := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	var result strings.Builder
	for _, v := range values {
		for num >= v {
			result.WriteString(dict[v])
			num -= v
		}
	}
	return result.String()
}

// Integer to Roman
// https://leetcode.com/problems/integer-to-roman/

// tags:
