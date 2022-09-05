package question

import "math/rand"

//398. 随机数索引

type Solution struct {
	m map[int][]int
}

func Constructor3(nums []int) Solution {
	res := Solution{
		m: make(map[int][]int, 0),
	}
	for i, v := range nums {
		if _, ok := res.m[v]; ok {
			res.m[v] = append(res.m[v], i)
		} else {
			res.m[v] = make([]int, 0)
			res.m[v] = append(res.m[v], i)
		}
	}
	return res
}

func (this *Solution) Pick(target int) int {
	res := this.m[target]
	index := rand.Intn(len(res))
	return res[index]
}
