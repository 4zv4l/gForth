package main

import (
	"bufio"
	"io"
	"strings"
)

type Words struct {
	scan    *bufio.Reader
	word    []string
	current string
}

// get words from user input
// return each word in an array
func (w Words) get() []string {
	words, err := w.scan.ReadString('\n')
	if err != nil {
		return []string{"bye"}
	}
	words = strings.Trim(words, "\n")
	// println("  ok")
	return strings.Split(words, " ")
}

// next set the current word as the next word in the array
// if no word is left then wait for user input
func (w *Words) next() {
	if len(w.word) > 0 {
		w.current = w.word[0]
		w.word = w.word[1:]
	} else {
		// if no word left
		// then wait for user input
		w.word = w.get()
		w.current = w.word[0]
		w.word = w.word[1:]
	}
}

// show each words until reach "bye"
func (w *Words) run() {
	for {
		w.next()
		interpret(w.current, w)
	}
}

func setup(r io.Reader) Words {
	w := Words{}
	w.scan = bufio.NewReader(r)
	return w
}
