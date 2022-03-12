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
			old := gw.word
			new := v
			gw.word = append(new, old...)
			// interpret(gw.word[1], gw)
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
		} else if w == "=" {
			stack.isEqual()
		} else if w == ">" {
			stack.isGreater()
		} else if w == "<" {
			stack.isLess()
		} else if w == "!" {
			stack.isNot()
		} else if w == "drop" {
			_, err := stack.pop()
			if err != nil {
				println(err.Error())
			}
		} else if w == "dup" {
			stack.dup()
		} else if w == "swap" {
			stack.swap()
		} else if w == "." {
			n, err := stack.pop()
			if err != nil {
				println(err.Error())
				return
			}
			fmt.Println(n)
		} else if w == "print" {
			fmt.Println(stack.data)
		} else if w == "key" {
			stack.key()
		} else if w == "emit" {
			c, err := stack.pop()
			if err != nil {
				println(err.Error())
				return
			}
			fmt.Printf("%c", c)
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
		} else if w == `."` { // print string
			str := ""
			for {
				gw.next()
				w = gw.current
				for c := range w {
					if w[c] == '"' {
						fmt.Print(str)
						return
					}
					str += string(w[c])
				}
				str += " "
			}
		} else if w == "branch" {
			// TODO implement branch
			// simple jump
			return
		} else if w == "branch?" {
			b, err := stack.pop()
			if err != nil {
				println(err.Error())
				return
			}
			// if not false
			if b != -1 {
				for {
					gw.next()
					w = gw.current
					if w == "then" {
						break
					}
					interpret(w, gw)
				}
			} else {
				for {
					gw.next()
					w = gw.current
					if w == "then" {
						break
					}
					// do not interpret
				}
			}
			return
		}
	}
}

func main() {
	var w Words
	if len(os.Args) < 2 {
		w = setup(os.Stdin)
		println("Welcome to gForth\ntype 'bye' to exit")
	} else {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Couldn't open the file")
			return
		}
		w = setup(f)
	}
	w.run()
}
