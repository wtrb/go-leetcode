package strings

// Time complexity: O(N)
// Space complexity: O(1)
func findAnagrams(s string, p string) []int {
	result := []int{}
	sLen := len(s)
	pLen := len(p)

	if sLen < pLen {
		return []int{}
	}

	pMap := [26]int{}
	for _, ch := range p {
		pMap[ch-'a']++
	}

	sMap := [26]int{}
	for i := 0; i < pLen; i++ {
		sMap[s[i]-'a']++
	}

	if matches(sMap, pMap) {
		result = append(result, 0)
	}

	for i := pLen; i < sLen; i++ {
		sMap[s[i]-'a']++
		sMap[s[i-pLen]-'a']--

		if matches(sMap, pMap) {
			result = append(result, i-pLen+1)
		}
	}

	return result
}

func matches(sMap, pMap [26]int) bool {
	for i := range pMap {
		if sMap[i] != pMap[i] {
			return false
		}
	}
	return true
}

// Find All Anagrams in a String

// https://leetcode.com/problems/find-all-anagrams-in-a-string/
// https://leetcode.com/submissions/detail/341030680/?from=/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3332/

// https://www.youtube.com/watch?v=PinhuDr_Q9c

// tags: map, hash map, anagram, permutation, kmp, sliding window
