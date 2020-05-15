package trie

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

// Time complexity : O(m), where m is the key length.
// Space complexity : O(m).
// In the worst case newly inserted key doesn't share a prefix with the the keys already inserted in the trie. We have to add mm new nodes, which takes us O(m) space.
func (this *Trie) Insert(word string) {
	curr := this

	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}

		curr = curr.children[idx]
	}
	curr.isEnd = true
}

// Time complexity : O(m)
// Space complexity : O(1)
func (this *Trie) Search(word string) bool {
	curr := this

	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}

		curr = curr.children[idx]
	}

	return curr.isEnd
}

// Time complexity : O(m)
// Space complexity : O(1)
func (this *Trie) StartsWith(prefix string) bool {
	curr := this

	for _, ch := range prefix {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}

		curr = curr.children[idx]
	}

	return true
}

/*
	Your Trie object will be instantiated and called as such:
	obj := Constructor();
	obj.Insert(word);
	param_2 := obj.Search(word);
	param_3 := obj.StartsWith(prefix);
*/

// Implement Trie (Prefix Tree)
// https://leetcode.com/problems/implement-trie-prefix-tree/
// https://leetcode.com/submissions/detail/339530407/?from=/explore/featured/card/may-leetcoding-challenge/535/week-2-may-8th-may-14th/3329/

// tags: trie, digital tree, prefix tree, tree
