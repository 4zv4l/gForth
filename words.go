package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

type Words struct {
	scan    *bufio.Reader
	word    []string
	current string
	isFile  bool
}

// get words from user input
// return each word in an array
func (w Words) get() []string {
	// get input from file
	if w.isFile {
		words, err := w.scan.ReadString('\n')
		if err != nil {
			return []string{"bye"}
		}
		words = strings.Trim(words, "\n")
		return strings.Split(words, " ")
	} else { // get input from keyboard
		var word string
		var l int
		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			if key == keyboard.KeyEnter {
				_ = keyboard.Close()
				words := strings.Split(word, " ")
				if _, err := strconv.Atoi(words[len(words)-1]); err == nil {
					println("  ok")
				} else if contains(builtin, word) {
					println("  ok")
				} else {
					println("  ?")
				}
				println("words:", word)
				return words
			}
			if key == keyboard.KeySpace {
				print(" ")
				word += " "
			}
			// handle backspace
			if key == keyboard.KeyBackspace || key == keyboard.KeyBackspace2 {
				if l > 0 {
					word = word[:l-1]
					fmt.Print("\b \b")
					l -= 2
				}
			}
			word += string(char)
			l += 1
			print(string(char))
		}

	}
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
