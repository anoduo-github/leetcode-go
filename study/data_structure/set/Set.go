package set

import (
	"fmt"
	"sync"
)

//Set 集合
type Set struct {
	m            map[string]struct{} //map实现set
	length       int                 //长度
	sync.RWMutex                     //锁
}

//NewSet 创建一个集合
func NewSet() *Set {
	return &Set{
		m:      make(map[string]struct{}, 0),
		length: 0,
	}
}

//Add 添加一个值
func (s *Set) Add(e string) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.m[e]; !ok {
		s.m[e] = struct{}{}
		s.length++
	}
}

//Delete 删除一个值
func (s *Set) Delete(e string) {
	s.Lock()
	defer s.Unlock()

	if s.length == 0 {
		return
	}

	if _, ok := s.m[e]; ok {
		delete(s.m, e)
		s.length--
	}
}

//Exist 是否存在一个值
func (s *Set) Exist(e string) bool {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.m[e]
	return ok
}

//Size 集合大小
func (s *Set) Size() int {
	return s.length
}

//Clear 清空集合
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()

	s.m = make(map[string]struct{}, 0)
	s.length = 0
}

//IsEmpty 集合是否为空
func (s *Set) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()

	return s.length == 0
}

//List 集合转为无序list
func (s *Set) List() []string {
	s.Lock()
	defer s.Unlock()

	res := make([]string, 0)
	for k := range s.m {
		res = append(res, k)
	}

	return res
}

//TestSet
func TestSet() {
	s := NewSet()
	s.Add("1")
	s.Add("1")
	s.Add("2")
	s.Add("3")
	s.Add("4")
	s.Add("5")

	fmt.Println(s.Size())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Exist("1"))

	s.Delete("yes")
	s.Delete("2")

	fmt.Println(s.Exist("2"))
	fmt.Println(s.Size())

	list := s.List()
	fmt.Println(list)

	s.Clear()
	fmt.Print(s.IsEmpty())
}
