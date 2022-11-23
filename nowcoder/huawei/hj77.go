package huawei

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//HJ77 火车进站
func HJ77() {
	//获取数据
	var n int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums[i] = num
	}

	//存结果
	m := make(map[string]struct{})

	//闭包函数
	var dfs func(i int, s string, stack []int)
	dfs = func(i int, s string, stack []int) {
		l := len(stack)
		if i == n {
			if l > 0 {
				for i := len(stack) - 1; i >= 0; i-- {
					s += strconv.Itoa(stack[i]) + " "
				}
			}
			s = strings.TrimSpace(s)
			m[s] = struct{}{}
			return
		}
		//入站并出站
		s_tmp1 := s
		s_tmp1 += strconv.Itoa(nums[i]) + " "
		stack_tmp1 := make([]int, l)
		copy(stack_tmp1, stack)
		dfs(i+1, s_tmp1, stack_tmp1)
		//入站不出
		stack_tmp2 := make([]int, l)
		copy(stack_tmp2, stack)
		stack_tmp2 = append(stack_tmp2, nums[i])
		dfs(i+1, s, stack_tmp2)
		//站里的车出来
		for k := 1; k <= l; k++ {
			s_tmp3 := s
			stack_tmp3 := make([]int, l)
			copy(stack_tmp3, stack)
			for j := l - 1; j >= l-k; j-- {
				s_tmp3 += strconv.Itoa(stack_tmp3[j]) + " "
			}
			stack_tmp3 = stack_tmp3[0 : l-k]
			dfs(i, s_tmp3, stack_tmp3)
		}
	}

	stack := make([]int, 0)
	dfs(0, "", stack)

	res := make([]string, 0)
	for k := range m {
		res = append(res, k)
	}
	sort.Strings(res)
	for _, v := range res {
		fmt.Println(v)
	}
}
