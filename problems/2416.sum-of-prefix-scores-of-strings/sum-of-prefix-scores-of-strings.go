package main

import "fmt"

type Trie struct {
	children [26]*Trie
	count    int
}

func Constructor() *Trie {
	return &Trie{}
}

func (trie *Trie) Insert(word string) {
	current := trie

	for _, r := range word {
		index := r - 'a'

		if current.children[index] == nil {
			current.children[index] = Constructor()
		}

		current = current.children[index]
		current.count += 1
	}
}

func (trie *Trie) CountPrefixes(word string) int {
	current, count := trie, 0

	for _, r := range word {
		index := r - 'a'

		if current.children[index] == nil {
			return count
		}
		current = current.children[index]
		count += current.count
	}

	return count
}

func sumPrefixScores(words []string) []int {
	tree := Constructor()
	for _, word := range words {
		tree.Insert(word)
	}

	result := make([]int, len(words))
	for idx, word := range words {
		result[idx] = tree.CountPrefixes(word)
	}

	return result
}

func main() {
	fmt.Println(sumPrefixScores([]string{"abc", "ab", "bc", "b"}))
}

