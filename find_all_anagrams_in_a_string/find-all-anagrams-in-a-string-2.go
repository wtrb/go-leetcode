package strings

func findAnagrams2(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	result := []int{}

	chars := [26]int{}
	for _, ch := range p {
		chars[ch-'a']++
	}

	start := 0
	end := 0
	count := len(p)
	for end < len(s) {
		//move right everytime, if the character exists in p's hash, decrease the count
		//current hash value >= 1 means the character is existing in p
		if chars[s[end]-'a'] >= 1 {
			count--
		}
		chars[s[end]-'a']--
		end++

		//when the count is down to 0, means we found the right anagram
		//then add window's left to result list
		if count == 0 {
			result = append(result, start)
		}

		//if we find the window's size equals to p, then we have to move left (narrow the window) to find the new match window
		//++ to reset the hash because we kicked out the left
		//only increase the count if the character is in p
		//the count >= 0 indicate it was original in the hash, cuz it won't go below 0
		if (end - start) == len(p) {
			if chars[s[start]-'a'] >= 0 {
				count++
			}
			chars[s[start]-'a']++
			start++
		}
	}

	return result
}

// Find All Anagrams in a String

// https://leetcode.com/problems/find-all-anagrams-in-a-string/
// https://leetcode.com/submissions/detail/340642478/?from=/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3332/

// tags: map, hash map, anagram, permutation
