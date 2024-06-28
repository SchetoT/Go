/*
Diseña e implementa una estructura de datos en Go para un árbol de búsqueda de prefijos (Trie) que almacene un conjunto de palabras.
Proporciona funciones para insertar palabras en el árbol y buscar palabras por prefijo.
*/
package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func newTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) SearchPrefix(prefix string) []string {
	node := t.root
	for _, char := range prefix {
		if _, ok := node.children[char]; !ok {
			return nil
		}
		node = node.children[char]
	}
	return t.collectWords(node, prefix)
}

func (t *Trie) collectWords(node *TrieNode, prefix string) []string {
	var words []string
	if node.isEnd {
		words = append(words, prefix)
	}
	for char, childNode := range node.children {
		words = append(words, t.collectWords(childNode, prefix+string(char))...)
	}
	return words
}

func main() {
	trie := newTrie()
	words := []string{"apple", "app", "apricot", "banana", "bat", "batch", "batman"}
	for _, word := range words {
		trie.Insert(word)
	}

	prefix := "app"
	fmt.Printf("Words with prefix '%s': %v\n", prefix, trie.SearchPrefix(prefix))

	prefix = "bat"
	fmt.Printf("Words with prefix '%s' : %v\n", prefix, trie.SearchPrefix(prefix))
}
