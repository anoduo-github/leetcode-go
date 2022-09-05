package stack

import (
	"errors"
	"sync"
)

//使用数组模拟栈

//ArrayStack 栈
type ArrayStack struct {
	stack []int      //栈
	size  int        //栈的大小
	lock  sync.Mutex //锁（并发安全）
}

//NewArrayStack 初始化创建一个栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		stack: make([]int, 0),
		size:  0,
		lock:  sync.Mutex{},
	}
}

//Push 入栈
func (s *ArrayStack) Push(e int) {
	s.lock.Lock()                //先上锁
	defer s.lock.Unlock()        //最后解锁
	s.stack = append(s.stack, e) //入栈，放入末尾
	s.size++                     //大小加一
}

//Pop 出栈
func (s *ArrayStack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		return 0, errors.New("栈中元素已空")
	}

	res := s.stack[s.size-1]

	//切片收缩
	//直接s.Stack[0:s.Size-1]并不好，因为底层数组大小未变，也就是多次入栈，数组只会越来越大，占用的空间也是

	//创建一个新的数组,长度和容量一致
	//这个就是要遍历原数组，耗时间
	temp := make([]int, s.size-1)
	for i := 0; i < s.size-1; i++ {
		temp[i] = s.stack[i]
	}
	s.stack = temp

	s.size--

	return res, nil
}

//Peek 获取栈顶元素值
func (s *ArrayStack) Peek() (int, error) {
	if s.size == 0 {
		return 0, errors.New("栈中元素已空")
	}

	return s.stack[s.size-1], nil
}

//Size 栈大小
func (s *ArrayStack) Size() int {
	return s.size
}

//IsEmpty 栈是否为空
func (s *ArrayStack) IsEmpty() bool {
	return s.size == 0
}
