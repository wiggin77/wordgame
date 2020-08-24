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

	disableWords []string
	addWords     []string

	showDisabled bool
	verbose      bool
}

func stringToRunes(s string) []rune {
	out := make([]rune, 0, len(s))
	for _, ch := range s {
		out = append(out, ch)
	}
	return out
}

func (o *Opts) setLetters(letters string) error {
	if len(letters) < o.minLength {
		return fmt.Errorf("not enough letters; min=%d", o.minLength)
	}

	letters = strings.ToLower(letters)
	arrLetters := stringToRunes(letters)

	for _, ch := range arrLetters {
		if !unicode.IsLetter(ch) {
			return fmt.Errorf("invalid letter (%v)", ch)
		}
	}
	o.letters = arrLetters
	return nil
}

func NewOpts() (*Opts, error) {
	opts := &Opts{}
	var disableWords bool
	var addWords bool

	flag.StringVar(&opts.wordsFile, "f", DefWordsFile, "Optional filespec to list of words.")
	flag.IntVar(&opts.minLength, "m", 3, "Minimum word length to find.")

	flag.BoolVar(&disableWords, "d", false, "Disable all words in args; space separated.")
	flag.BoolVar(&addWords, "a", false, "Add all words in args; space separated.")

	flag.BoolVar(&opts.showDisabled, "s", false, "Show disabled words.")

	flag.BoolVar(&opts.verbose, "v", false, "Verbose output.")
	flag.Parse()

	var err error

	args := flag.Args()
	argsCount := len(args)

	if addWords && disableWords {
		return nil, errors.New("cannot used -a and -d together.")
	}

	if addWords {
		if argsCount == 0 {
			return nil, errors.New("not enough args; supply words to add")
		}
		opts.addWords = args
		return opts, nil
	}

	if disableWords {
		if argsCount == 0 {
			return nil, errors.New("not enough args; supply words to disable")
		}
		opts.disableWords = args
		return opts, nil
	}

	switch argsCount {
	case 0:
		return nil, errors.New("not enough args; supply letters")
	case 1:
		err = opts.setLetters(args[0])
	default:
		return nil, errors.New("too many args")
	}
	return opts, err
}
