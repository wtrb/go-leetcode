package string

// Time complexity : O(l1 + (l2 - l1)*(l1+26))
// Space complexity : O(1)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map := [26]int{}
	for _, ch := range s1 {
		s1Map[ch-'a']++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		s2Map := [26]int{}
		for j := 0; j < len(s1); j++ {
			s2Map[s2[i+j]-'a']++
		}

		if ok := matches(s1Map, s2Map); ok {
			return true
		}
	}

	return false
}

func matches(s1Map, s2Map [26]int) bool {
	for i := 0; i < 26; i++ {
		if s1Map[i] != s2Map[i] {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/permutation-in-string/
// https://leetcode.com/submissions/detail/341151853/?from=/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3333/

// tags: map, hash map, anagram, permutation
