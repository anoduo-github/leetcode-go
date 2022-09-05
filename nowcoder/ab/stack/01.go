package stack

import (
	"errors"
	"sync"
)

//AB1 栈

//stack 栈
type Stack struct {
	arr  []int
	size int
	lock sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		arr:  make([]int, 0),
		size: 0,
		lock: sync.Mutex{},
	}
}

func (s *Stack) Push(x int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.arr = append(s.arr, x)
	s.size++
}

func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		return 0, errors.New("error")
	}
	res := s.arr[s.size-1]
	s.arr = s.arr[:s.size-1]
	s.size--
	return res, nil
}

func (s *Stack) Top() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		return 0, errors.New("error")
	}

	return s.arr[s.size-1], nil
}
