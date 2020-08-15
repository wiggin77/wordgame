package main

import (
	"bufio"
	"os"
	"sort"
)

type Words struct {
	list []string
}

func NewWords(opts *Opts) (*Words, error) {
	words := &Words{}

	if err := words.load(opts.wordsFile); err != nil {
		return nil, err
	}

	return words, nil
}

func (w *Words) load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	var list []string

	for scanner.Scan() {
		line := scanner.Text()
		list = append(list, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	sort.Strings(list)
	w.list = list
	return nil
}
