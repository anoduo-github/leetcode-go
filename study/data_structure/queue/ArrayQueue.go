package queue

import (
	"errors"
	"sync"
)

//ArrayQueue 数组队列
type ArrayQueue struct {
	array []string   //字符切片
	size  int        //大小
	lock  sync.Mutex //锁
}

//NewArrayQueue 初始化创建一个数组队列
func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		array: make([]string, 0),
		size:  0,
		lock:  sync.Mutex{},
	}
}

//Add 入队
func (a *ArrayQueue) Add(v string) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.array = append(a.array, v)
	a.size++
}

//Remove 出队
func (a *ArrayQueue) Remove() (string, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.size == 0 {
		return "", errors.New("队列中元素已空")
	}

	res := a.array[0]

	temp := make([]string, a.size-1)
	for i := 1; i < a.size; i++ {
		temp[i-1] = a.array[i]
	}
	a.array = temp
	a.size--

	return res, nil
}

//Size 大小
func (a *ArrayQueue) Size() int {
	return a.size
}

//IsEmpty 是否为空
func (a *ArrayQueue) IsEmpty() bool {
	return a.size == 0
}
