package ds

import (
	"fmt"
)

// TrieNode represents each node in the Trie
type TrieNode struct {
	children  map[rune]*TrieNode
	isEnd     bool
	wordCount int // New field to keep track of the number of words passing through or ending at this node
}

// Trie represents the Trie data structure
type Trie struct {
	root *TrieNode
}

// NewTrieNode creates a new Trie node
func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// Insert inserts a word into the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
		node.wordCount++ // Increment word count at each node
	}
	node.isEnd = true
}

// Search searches for a word in the Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

// StartsWith checks if there is any word in the Trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return true
}

// Delete deletes a word from the Trie
func (t *Trie) Delete(word string) bool {
	if t.delete(t.root, word, 0) {
		node := t.root
		for _, char := range word {
			if _, exists := node.children[char]; exists {
				node = node.children[char]
				node.wordCount-- // Decrement word count at each node
			}
		}
		return true
	}
	return false
}

// delete is a helper function for deleting a word from the Trie
func (t *Trie) delete(node *TrieNode, word string, depth int) bool {
	if depth == len(word) {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		return len(node.children) == 0
	}
	char := rune(word[depth])
	childNode, exists := node.children[char]
	if !exists {
		return false
	}
	if t.delete(childNode, word, depth+1) {
		delete(node.children, char)
		return len(node.children) == 0 && !node.isEnd
	}
	return false
}

// CountWordsWithPrefix counts the number of words starting with the given prefix
func (t *Trie) CountWordsWithPrefix(prefix string) int {
	node := t.root
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return 0
		}
		node = node.children[char]
	}
	return node.wordCount
}

// WordsWithPrefix returns a list of words starting with the given prefix
func (t *Trie) WordsWithPrefix(prefix string) []string {
	node := t.root
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return nil
		}
		node = node.children[char]
	}
	var words []string
	t.collectWords(node, prefix, &words)
	return words
}

// collectWords is a helper function for collecting words in a subtree
func (t *Trie) collectWords(node *TrieNode, prefix string, words *[]string) {
	if node.isEnd {
		*words = append(*words, prefix)
	}
	for char, child := range node.children {
		t.collectWords(child, prefix+string(char), words)
	}
}

func main() {
	trie := NewTrie()

	// Insert words
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("apricot")
	trie.Insert("bat")
	trie.Insert("ball")
	trie.Insert("banana")

	// Search for words
	fmt.Println("Search 'apple':", trie.Search("apple"))     // true
	fmt.Println("Search 'app':", trie.Search("app"))         // true
	fmt.Println("Search 'apricot':", trie.Search("apricot")) // true
	fmt.Println("Search 'bat':", trie.Search("bat"))         // true
	fmt.Println("Search 'ball':", trie.Search("ball"))       // true
	fmt.Println("Search 'banana':", trie.Search("banana"))   // true
	fmt.Println("Search 'grape':", trie.Search("grape"))     // false

	// Check for prefixes
	fmt.Println("StartsWith 'app':", trie.StartsWith("app"))     // true
	fmt.Println("StartsWith 'ba':", trie.StartsWith("ba"))       // true
	fmt.Println("StartsWith 'grape':", trie.StartsWith("grape")) // false

	// Delete words
	fmt.Println("Delete 'app':", trie.Delete("app"))                // true
	fmt.Println("Search 'app' after deletion:", trie.Search("app")) // false

	// Count words with prefix
	fmt.Println("Count words with prefix 'app':", trie.CountWordsWithPrefix("app"))     // 1
	fmt.Println("Count words with prefix 'b':", trie.CountWordsWithPrefix("b"))         // 2
	fmt.Println("Count words with prefix 'grape':", trie.CountWordsWithPrefix("grape")) // 0

	// List words with prefix
	fmt.Println("Words with prefix 'app':", trie.WordsWithPrefix("app"))     // [apple apricot]
	fmt.Println("Words with prefix 'b':", trie.WordsWithPrefix("b"))         // [bat ball banana]
	fmt.Println("Words with prefix 'grape':", trie.WordsWithPrefix("grape")) // []
}
