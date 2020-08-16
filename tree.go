package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

type Node struct {
	r        rune
	word     bool
	children map[rune]*Node
}

func NewNode(r rune) *Node {
	return &Node{
		r:        r,
		word:     false,
		children: make(map[rune]*Node),
	}
}

func addWordToTree(root *Node, word string) {
	node := root
	for _, r := range word {
		fnode, ok := node.children[r]
		if !ok {
			fnode = NewNode(r)
			node.children[r] = fnode
		}
		node = fnode
	}
	node.word = true
}

// searchTree walks the tree checking if each item in arr is a child of the
// previous. If all items of arr exist in the tree, then found=true. If the
// last item of arr happens to be the last character of a word, then word=true.
func searchTree(root *Node, arr []rune) (found bool, word bool) {
	if len(arr) == 0 || root == nil {
		return false, false
	}

	node := root
	for _, r := range arr {
		// find this rune in the children.
		fnode, ok := node.children[r]
		if !ok {
			return false, false
		}
		node = fnode
	}
	return true, node.word
}

func printNode(n *Node, prefix string) {
	buf := make([]byte, utf8.UTFMax)
	var b []byte
	for k := range n.children {
		utf8.EncodeRune(buf, k)
		b = append(b, buf...)
	}
	fmt.Fprintf(os.Stdout, "prefix: %s  rune: %s  word: %t  children: %s\n", prefix, string(n.r), n.word, string(b))

	r := []rune(prefix)
	r = append(r, n.r)
	prefix = string(r)
	for _, child := range n.children {
		printNode(child, prefix)
	}
}
