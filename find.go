package main

import (
	"fmt"
	"sort"
)

func FindWords(opts *Opts, dict *Dictionary) []string {

	wordsFound := make(map[string]struct{})

	var permutations func([]rune, []rune)

	permutations = func(prefix []rune, arr []rune) {
		if len(prefix) > 0 {
			found, word := searchTree(dict.root, prefix)
			if !found {
				verbosePrintln(opts, "prefix not found: ", string(prefix))
				return
			}

			if word && len(prefix) >= opts.minLength {
				w := string(prefix)
				verbosePrintln(opts, "word found: ", w)
				wordsFound[w] = struct{}{}
			}
		}

		n := len(arr)
		if n == 0 {
			verbosePrintln(opts, "combo: ", string(prefix))
		} else {
			for i := 0; i < n; i++ {
				permutations(runeAppend(prefix, arr[i]), runeAppend(arr[:i], arr[i+1:]...))
			}
		}
	}

	permutations(nil, opts.letters)

	list := make([]string, 0, len(wordsFound))
	for k := range wordsFound {
		list = append(list, k)
	}
	sort.Strings(list)
	return list
}

func runeAppend(arr1 []rune, arr2 ...rune) []rune {
	out := make([]rune, 0, len(arr1)+len(arr2))

	out = append(out, arr1...)
	out = append(out, arr2...)
	return out
}

func verbosePrintln(opts *Opts, args ...interface{}) {
	if opts.verbose {
		fmt.Println(args...)
	}
}
