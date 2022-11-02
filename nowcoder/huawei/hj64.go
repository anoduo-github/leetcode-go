package huawei

import "fmt"

//HJ64 MP3光标位置
func HJ64(n int, commond string) {
	cur := 1      //初始指向第一首，光标位置
	left := 1     //上边界
	var right int //下边界
	if n < 4 {
		right = n
	} else {
		right = 4
	}

	//拆分命令
	for _, v := range commond {
		if v == 'U' {
			//朝上
			cur = cur - 1
			if cur == 0 {
				//换页，换到最后一页
				cur = n
				right = n
				if n < 4 {
					left = 1
				} else {
					left = n - 3
				}
			} else if cur < left {
				//小于上边界，向上移页的位置
				left--
				right--
			}
		} else if v == 'D' {
			//向下移
			cur = cur + 1
			if cur > n {
				//换页，换到第一页
				cur = 1
				left = 1
				if n < 4 {
					right = n
				} else {
					right = 4
				}
			} else if cur > right {
				//大于下边界,向下移页的位置
				left++
				right++
			}
		}
	}

	//输出当前歌曲
	for i := left; i <= right; i++ {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Print("\n")
	//输出当前歌曲
	fmt.Println(cur)
}
