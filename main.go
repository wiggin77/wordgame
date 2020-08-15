package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	DefWordsFile = "./words_alpha"
)

func main() {
	opts, err := NewOpts()
	if err != nil {
		errExit(err, 1, true)
	}

	words, err := NewWords(opts)
	if err != nil {
		errExit(err, 2, false)
	}

	if err := FindWords(opts, words); err != nil {
		errExit(err, 3, false)
	}
}

func errExit(err error, code int, usage bool) {
	fmt.Fprintln(os.Stderr, err)
	if usage {
		flag.Usage()
	}
	os.Exit(code)
}
