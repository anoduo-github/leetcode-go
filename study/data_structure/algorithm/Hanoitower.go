package algorithm

import "fmt"

//分治算法
//汉诺塔问题

//HanoiTower 解决汉诺塔问题,a,b,c指的是盘经过的柱子的顺序
func HanoiTower(num int, a, b, c byte) {
	if num == 1 { //只有一个盘的情况下
		fmt.Printf("第1个盘从 %c->%c\n", a, c)
	} else {
		//大于一个盘，把问题分成两部分，最下面的一个盘和上面的盘，
		//也就是说上面的盘先移动到b,然后将下面的盘移动到c,最后将b的盘移动到c,完成

		//首先移动上面的盘到b
		HanoiTower(num-1, a, c, b)
		//再将下面的盘移动到c
		fmt.Printf("第%d个盘从 %c->%c\n", num, a, c)
		//最后将b的盘移动到c
		HanoiTower(num-1, b, a, c)
	}
}
