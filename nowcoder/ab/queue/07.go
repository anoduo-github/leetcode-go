package queue

import (
	"errors"
	"sync"
)

//Queue 队列
type Queue struct {
	arr  []int
	size int
	lock sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		arr:  make([]int, 0),
		size: 0,
		lock: sync.Mutex{},
	}
}

func (q *Queue) Push(x int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.arr = append(q.arr, x)
	q.size++
}

func (q *Queue) Pop() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		return 0, errors.New("error")
	}
	res := q.arr[0]
	q.arr = q.arr[1:]
	q.size--
	return res, nil
}

func (q *Queue) Front() (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		return 0, errors.New("error")
	}
	return q.arr[0], nil
}
