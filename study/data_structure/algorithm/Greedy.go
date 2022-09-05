package algorithm

import (
	"fmt"
)

//set
type set struct {
	m map[string]struct{}
	//sync.RWMutex 简单的set,算了
	size int
}

//newSet
func newSet() *set {
	return &set{
		m:    make(map[string]struct{}, 0),
		size: 0,
	}
}

//add
func (s *set) add(str string) {
	if _, ok := s.m[str]; !ok {
		s.m[str] = struct{}{}
		s.size++
	}

}

//addAll
func (s *set) addAll(t *set) {
	for k, _ := range t.m {
		s.add(k)
	}
}

//getSize
func (s *set) getSize() int {
	return s.size
}

//retainAll 求交集
func (s *set) retainAll(target *set) {
	for k, _ := range s.m {
		if _, ok := target.m[k]; !ok {
			delete(s.m, k)
			s.size--
		}
	}
}

//removeAll 删除指定的key
func (s *set) removeAll(target *set) {
	for k, _ := range target.m {
		delete(s.m, k)
		s.size--
	}
}

//broadcast
type broadcast struct {
	elem map[string]*set //数据
	list []string        //键排序
}

//TestGreedy 测试贪心算法
func TestGreedy() {
	//创建广播台
	b := new(broadcast)
	b.elem = make(map[string]*set, 0)
	b.list = make([]string, 0)

	//将各个电台放入广播台中
	set1 := newSet()
	set1.add("北京")
	set1.add("上海")
	set1.add("天津")

	set2 := newSet()
	set2.add("北京")
	set2.add("广州")
	set2.add("深圳")

	set3 := newSet()
	set3.add("成都")
	set3.add("上海")
	set3.add("杭州")

	set4 := newSet()
	set4.add("天津")
	set4.add("上海")

	set5 := newSet()
	set5.add("大连")
	set5.add("杭州")

	b.elem["K1"] = set1
	b.elem["K2"] = set2
	b.elem["K3"] = set3
	b.elem["K4"] = set4
	b.elem["K5"] = set5
	b.list = append(b.list, "K1", "K2", "K3", "K4", "K5")

	//存放所有地区
	allSet := newSet()
	allSet.add("北京")
	allSet.add("上海")
	allSet.add("天津")
	allSet.add("广州")
	allSet.add("深圳")
	allSet.add("成都")
	allSet.add("杭州")
	allSet.add("大连")

	//存放选择的电台集合
	res := make([]string, 0)

	//遍历
	for allSet.getSize() != 0 {
		maxKey := ""
		for _, v := range b.list {
			tempSet := b.elem[v]
			//取交集
			tempSet.retainAll(allSet)
			if tempSet.getSize() > 0 && (maxKey == "" || tempSet.getSize() > b.elem[maxKey].getSize()) {
				maxKey = v
			}
		}
		if maxKey != "" {
			res = append(res, maxKey)
			allSet.removeAll(b.elem[maxKey])
		}
	}

	fmt.Println(res)
}
