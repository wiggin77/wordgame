package main

import (
	"bufio"
	"os"
	"strings"
)

type Dictionary struct {
	root *Node
}

func NewDictionary(opts *Opts) (*Dictionary, error) {
	words := &Dictionary{root: NewNode(0)}

	if err := words.load(opts.wordsFile); err != nil {
		return nil, err
	}

	return words, nil
}

func (d *Dictionary) load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToLower(line)

		// strip annotations
		line = strings.Trim(line, "!@#$%^&*()-+")

		addWordToTree(d.root, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
