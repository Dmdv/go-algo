# Trie

```go
package main

import "github.com/dmdv/go-algo/ds"

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

```
