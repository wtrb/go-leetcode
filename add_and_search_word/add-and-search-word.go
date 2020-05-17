package word

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

// Constructor initializes your data structure here
func Constructor() WordDictionary {
	return WordDictionary{}
}

// AddWord adds a word into the data structure
func (wd *WordDictionary) AddWord(word string) {
	cur := wd
	for _, ch := range word {
		idx := ch - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &WordDictionary{}
		}
		cur = cur.children[idx]
	}
	cur.isEnd = true
}

// Search returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter.
func (wd *WordDictionary) Search(word string) bool {
	cur := wd
	for i, ch := range word {
		if ch == '.' {
			for _, child := range cur.children {
				if child != nil && child.Search(word[i+1:]) {
					return true
				}
			}
			return false

		} else {
			idx := ch - 'a'
			if cur.children[idx] == nil {
				return false
			}
			cur = cur.children[idx]
		}
	}
	return cur.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

// https://leetcode.com/problems/add-and-search-word-data-structure-design/

// tags: trie, digital tree, prefix tree, tree
