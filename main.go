package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	DefWordsFile = "./dictionary/words_alpha.txt"
)

func main() {
	opts, err := NewOpts()
	if err != nil {
		errExit(err, 1, true)
	}

	dict, err := NewDictionary(opts)
	if err != nil {
		errExit(err, 2, false)
	}

	if opts.addWords != nil {
		count, errAdd := dict.addWords(opts.addWords)
		if count > 0 {
			if err := dict.save(opts.wordsFile); err != nil {
				errExit(fmt.Errorf("error saving %s: %w", opts.wordsFile, err), 3, false)
			} else {
				fmt.Printf("Saved dictionary %s\n%d words added.", opts.wordsFile, count)
			}
		}
		if errAdd != nil {
			fmt.Fprintln(os.Stderr, errAdd)
		}
		return
	}

	if opts.disableWords != nil {
		count, err := dict.disableWords(opts.disableWords)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := dict.save(opts.wordsFile); err != nil {
			errExit(fmt.Errorf("error saving %s: %w", opts.wordsFile, err), 4, false)
		} else {
			fmt.Printf("Saved dictionary %s\n%d words removed.", opts.wordsFile, count)
		}
		return
	}

	wordsFound := FindWords(opts, dict)
	if len(wordsFound) > 0 {
		printWords(opts, wordsFound)
	} else {
		fmt.Println("No words found.")
	}

	//printNode(words.root, "")
}

func errExit(err error, code int, usage bool) {
	fmt.Fprintln(os.Stderr, err)
	if usage {
		flag.Usage()
	}
	os.Exit(code)
}

func printWords(opts *Opts, words []string) {
	// group the words by length
	buckets := make([][]string, len(opts.letters)+1)
	for _, w := range words {
		pos := len(w)
		list := buckets[pos]
		list = append(list, w)
		buckets[pos] = list
	}

	for i, list := range buckets {
		if len(list) > 0 {
			fmt.Printf("\n%d:\n%v\n", i, list)
		}
	}
	fmt.Println("")
}
