package main

type Trie struct {
	children [26]*Trie
	wordEnd  bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	current := this

	for _, r := range word {
		index := r - 'a'

		if current.children[index] == nil {
			node := Constructor()
			current.children[index] = &node
		}

		current = current.children[index]
	}

	current.wordEnd = true
}

func (this *Trie) Search(word string) bool {
	current := this

	for _, r := range word {
		index := r - 'a'

		if current.children[index] == nil {
			return false
		}

		current = current.children[index]
	}

	return current.wordEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this

	for _, r := range prefix {
		index := r - 'a'

		if current.children[index] == nil {
			return false
		}

		current = current.children[index]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

