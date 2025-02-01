package main

import (
	"strconv"
)

type TrieNode struct {
	wordEnd  bool
	children []*TrieNode
}

func getNode() *TrieNode {
	node := &TrieNode{
		wordEnd:  false,
		children: make([]*TrieNode, 10),
	}

	return node
}

func (root *TrieNode) insert(key int) {
	current := root
	kstr := strconv.Itoa(key)
	for _, r := range kstr {
		index := int(r - 48)
		if current.children[index] == nil {
			current.children[index] = getNode()
		}

		current = current.children[index]
	}

	current.wordEnd = true
}

func (root *TrieNode) search(key int) int {
	current := root
	kstr, count := strconv.Itoa(key), 0
	for _, r := range kstr {
		index := int(r - 48)
		if current.children[index] == nil {
			return count
		}
		count += 1

		current = current.children[index]
	}

	return count
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	tree := getNode()
	for _, num := range arr1 {
		tree.insert(num)
	}

	result := 0

	for _, num := range arr2 {
		count := tree.search(num)
		if count > result {
			result = count
		}
	}

	return result
}

func main() {
	println(longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
	println(longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4}))
}
