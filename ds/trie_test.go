package ds

import "testing"

func TestTrie_Insert(t1 *testing.T) {
	type fields struct {
		root *TrieNode
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				root: tt.fields.root,
			}
			t.Insert(tt.args.word)
		})
	}
}
