package main

import (
	"errors"

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
func (s *Stack) pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("Stack underflow")
	}
	l := len(s.data) - 1
	r := s.data[l]
	s.data = s.data[:l]
	return r, nil
}

func (s *Stack) add() {
	n1, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	n2, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	s.push(n1 + n2)
}

func (s *Stack) sub() {
	n1, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	n2, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	s.push(n2 - n1)
}

func (s *Stack) mul() {
	n1, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	n2, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	s.push(n1 * n2)
}

func (s *Stack) div() {
	n1, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	n2, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
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
	v, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	s.push(v)
	s.push(v)
}

// swap the top two value
func (s *Stack) swap() {
	n1, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	n2, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	s.push(n1)
	s.push(n2)
}

// 0 = trye
// -1 = false

func (s *Stack) isEqual() {
	n1, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	n2, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	if n1 == n2 {
		s.push(0)
	} else {
		s.push(-1)
	}
}

func (s *Stack) isGreater() {
	n1, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	n2, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	if n1 > n2 {
		s.push(0)
	} else {
		s.push(-1)
	}
}

func (s *Stack) isLess() {
	n1, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	n2, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	if n1 < n2 {
		s.push(0)
	} else {
		s.push(-1)
	}
}

func (s *Stack) isNot() {
	n, err := s.pop()
	if err != nil {
		println(err.Error())
		return
	}
	if n == 0 {
		s.push(-1)
	} else {
		s.push(0)
	}
}
