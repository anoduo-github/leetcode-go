package huawei

import "fmt"

//HJ53 杨辉三角的变形
func HJ53() {
	var n int
	fmt.Scanln(&n)

	if n < 3 {
		fmt.Println(-1)
		return
	}

	temp := n % 4
	switch temp {
	case 0:
		fmt.Println(3)
	case 1:
		fmt.Println(2)
	case 2:
		fmt.Println(4)
	case 3:
		fmt.Println(2)
	}

	/* if n == 1 {
		fmt.Println(-1)
		return
	}

	length := 2*n - 1
	mid := length / 2
	temp := make([]int, length)
	temp[mid] = 1
	res := make([]int, length)
	for i := 1; i < n; i++ {
		left := mid - i  //左边界
		right := mid + i //右边界
		for j := left; j <= right; j++ {
			leftNum := 0     //左上边的数
			rightNum := 0    //右上边的数
			upNum := temp[j] //上面的数
			if j > 0 {
				leftNum = temp[j-1]
			}
			if j+1 < length {
				rightNum = temp[j+1]
			}
			res[j] = leftNum + rightNum + upNum
		}
		//复制
		copy(temp, res)
	}
	for i, v := range res {
		if v%2 == 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1) */
}
