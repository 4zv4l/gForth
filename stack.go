package main

import (
	"github.com/eiannone/keyboard"
)

type Stack struct {
	data []int
}

var stack Stack

// push add a value on top of the stack
func (s *Stack) push(n int) {
	s.data = append(s.data, n)
}

// pop a number from the stack
// return that number
func (s *Stack) pop() int {
	l := len(s.data) - 1
	r := s.data[l]
	s.data = s.data[:l]
	return r
}

// all builtin command here
var builtin = []string{"+", "-", "*", "/", "drop", "dup", ".", "show", "key", "emit", "cr", "bye"}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s *Stack) add() {
	s.push(s.pop() + s.pop())
}

func (s *Stack) sub() {
	n1 := s.pop()
	n2 := s.pop()
	s.push(n1 - n2)
}

func (s *Stack) mul() {
	s.push(s.pop() * s.pop())
}

func (s *Stack) div() {
	n1 := s.pop()
	n2 := s.pop()
	if n1 == 0 {
		println("Divide by zero")
		return
	}
	s.push(n2 / n1)
}

// get a char from stdin
// put the ascii code on the stack
func (s *Stack) key() {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	s.push(int(char))
}

// dup dupliate the top of the stack
func (s *Stack) dup() {
	v := s.pop()
	s.push(v)
	s.push(v)
}
