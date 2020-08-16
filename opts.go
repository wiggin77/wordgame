package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"unicode"
)

type Opts struct {
	wordsFile string
	letters   []rune
	minLength int
	verbose   bool
}

func (o *Opts) setLetters(letters string) error {
	if len(letters) < o.minLength {
		return fmt.Errorf("not enough letters; min=%d", o.minLength)
	}

	letters = strings.ToLower(letters)

	for _, ch := range letters {
		if !unicode.IsLetter(ch) {
			return fmt.Errorf("invalid letter (%v)", ch)
		}
		o.letters = append(o.letters, ch)
	}
	return nil
}

func NewOpts() (*Opts, error) {
	opts := &Opts{}

	flag.StringVar(&opts.wordsFile, "f", DefWordsFile, "Optional filespec to list of words.")
	flag.IntVar(&opts.minLength, "m", 3, "Minimum word length to find.")
	flag.BoolVar(&opts.verbose, "v", false, "Verbose output.")
	flag.Parse()

	var err error

	args := flag.Args()
	switch len(args) {
	case 0:
		return nil, errors.New("not enough args; supply letters")
	case 1:
		err = opts.setLetters(args[0])
	default:
		return nil, errors.New("too many args")
	}
	return opts, err
}
