package stack

import (
	"fmt"
	"sync"
)

const (
	expandFactor = 2
)

type Stack[T any] struct {
	sync.RWMutex
	data []T
	size int
	cap  int
}

func NewStack[T any](cap int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, cap),
		size: 0,
		cap:  cap,
	}
}

func (s *Stack[T]) Push(data T) {
	s.Lock()
	defer s.Unlock()
	if s.size >= s.cap {
		// 需要扩容
		s.expand()
	}
	s.data[s.size] = data
	s.size++
}

func (s *Stack[T]) Pop() (val T) {
	s.Lock()
	defer s.Unlock()
	if s.size <= 0 {
		return
	}
	val = s.data[s.size-1]
	var empty T
	s.data[s.size-1] = empty
	s.size--
	return
}

func (s *Stack[T]) Top() (val T) {
	s.RLock()
	defer s.RUnlock()
	if s.size <= 0 {
		return
	}
	return s.data[s.size-1]
}

func (s *Stack[T]) Len() int {
	s.RLock()
	defer s.RUnlock()
	return s.size
}

func (s *Stack[T]) Cap() int {
	s.RLock()
	defer s.RUnlock()
	return s.cap
}

func (s *Stack[T]) Empty() bool {
	s.RLock()
	defer s.RUnlock()
	return s.Len() == 0
}

func (s *Stack[T]) expand() {
	newCap := s.cap * expandFactor
	newData := make([]T, newCap)
	copy(newData, s.data)
	s.cap = newCap
	s.data = newData
}

func (s *Stack[T]) PrintStack() {
	s.RLock()
	defer s.RUnlock()
	for i := s.size - 1; i >= 0; i-- {
		fmt.Printf("| %02v |\n", s.data[i])
	}
}
