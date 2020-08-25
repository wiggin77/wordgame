package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/wiggin77/merror"
)

type Dictionary struct {
	root  *Node
	words []string
}

func NewDictionary(opts *Opts) (*Dictionary, error) {
	words := &Dictionary{root: NewNode(0)}

	if err := words.load(opts.wordsFile, opts.showDisabled); err != nil {
		return nil, err
	}

	return words, nil
}

func (d *Dictionary) load(file string, showDisabled bool) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToLower(line)
		line = strings.TrimSpace(line)

		d.words = append(d.words, line)

		disabled := strings.HasPrefix(line, "!") || !hasVowel(line)

		// strip annotations
		line = strings.Trim(line, "!@#$%^&*()-+")

		addWordToTree(d.root, line, disabled)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if len(d.words) > 0 {
		sort.Strings(d.words)
	}
	return nil
}

func (d *Dictionary) save(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	defer f.Close()

	errs := merror.New()
	sort.Strings(d.words)

	for _, w := range d.words {
		if _, err := fmt.Fprintln(writer, w); err != nil {
			errs.Append(err)
		}
	}

	if err := writer.Flush(); err != nil {
		errs.Append(err)
	}
	return errs.ErrorOrNil()
}

func (d *Dictionary) disableWords(words []string) (int, error) {
	errs := merror.New()
	sort.Strings(d.words)
	var count int

	for _, w := range words {
		idx, found := binarySearch(d.words, w)
		if found {
			d.words[idx] = "!" + w
			count++
		} else {
			if _, found := binarySearch(d.words, "!"+w); found {
				errs.Append(fmt.Errorf("'%s' already disabled", w))
			} else {
				errs.Append(fmt.Errorf("'%s' not found", w))
			}
		}
	}
	return count, errs.ErrorOrNil()
}

func (d *Dictionary) addWords(words []string) (int, error) {
	if d.words == nil {
		d.words = make([]string, 0, len(words))
	}

	// can't just append the slice since we must not allow duplicates.
	var count int
	errs := merror.New()
	sort.Strings(d.words)
	for _, w := range words {
		_, found := binarySearch(d.words, w)
		if !found {
			d.words = append(d.words, w)
			count++
		} else {
			errs.Append(fmt.Errorf("'%s' already in dictionary", w))
		}
	}
	sort.Strings(d.words)
	return count, errs.ErrorOrNil()
}

func binarySearch(arr []string, s string) (int, bool) {
	sort.Strings(arr)
	idx := sort.SearchStrings(arr, s)
	found := idx != len(arr) && arr[idx] == s
	return idx, found
}

func hasVowel(s string) bool {
	for _, ch := range s {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'y':
			return true
		}
	}
	return false
}
