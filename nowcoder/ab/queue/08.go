package queue

import (
	"errors"
	"sync"
)

//AB8 【模板】循环队列

//LoopQueue 循环队列
type LoopQueue struct {
	arr      []int
	capacity int //容量
	size     int //大小
	lock     sync.Mutex
}

func NewLoopQueue(n int) *LoopQueue {
	return &LoopQueue{
		arr:      make([]int, n),
		capacity: n,
		size:     0,
		lock:     sync.Mutex{},
	}
}

func (l *LoopQueue) Push(x int) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.size < l.capacity {
		l.arr[l.size] = x
		l.size++
		return nil
	} else {
		return errors.New("full")
	}
}

func (l *LoopQueue) Front() (int, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.size == 0 {
		return 0, errors.New("empty")
	} else {
		return l.arr[0], nil
	}
}

func (l *LoopQueue) Pop() (int, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.size == 0 {
		return 0, errors.New("empty")
	} else {
		temp := l.arr[0]
		l.size--
		for i := 0; i < l.size; i++ {
			l.arr[i] = l.arr[i+1]
		}
		return temp, nil
	}
}
