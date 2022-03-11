package main

import (
	"fmt"
	"os"
	"strconv"
)

// interpret check the words to execute it
func interpret(w string, gw *Words) {
	// is number
	if n, err := strconv.Atoi(w); err == nil {
		stack.push(n)
	} else {
		// if in the dictionary
		if v, ok := dictionary[w]; ok {
			for _, a := range v {
				interpret(a, gw)
			}
		}
		// then words
		if w == "bye" {
			os.Exit(0)
		} else if w == "+" {
			stack.add()
		} else if w == "-" {
			stack.sub()
		} else if w == "*" {
			stack.mul()
		} else if w == "/" {
			stack.div()
		} else if w == "drop" {
			stack.pop()
		} else if w == "dup" {
			stack.dup()
		} else if w == "." {
			fmt.Println(stack.pop())
		} else if w == "show" {
			fmt.Println(stack.data)
		} else if w == "key" {
			stack.key()
		} else if w == "emit" {
			fmt.Printf("%c", stack.pop())
		} else if w == "cr" {
			println()
		} else if w == ":" { // add word to dictionary
			var word []string
			for {
				gw.next()
				w = gw.current
				if w == ";" {
					break
				}
				word = append(word, w)
			}
			dictionary[word[0]] = word[1:]
		} else if w == "words" {
			fmt.Println(dictionary)
		}
	}
}

func main() {
	var w Words
	if len(os.Args) < 2 {
		w = setup(os.Stdin)
	} else {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Couldn't open the file")
			return
		}
		w = setup(f)
		w.isFile = true
	}
	w.run()
}
