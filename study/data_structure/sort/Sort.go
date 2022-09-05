package sort

import (
	"fmt"
	tree "leetcode-go/study/data_structure/tree/sort"
)

//BubbleSort 冒泡排序
func BubbleSort(list []int) {
	length := len(list)
	//判断符，用于判断是否发生了交换，如果一直为true则表示没有发生一次交换，已经是排好的，直接返回
	tag := true
	for length > 0 {
		for i := 0; i < length-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i] //左右交换，然后i+1继续比较，保证每次循环完后右边的值是多大的
				tag = false                             //表示发生了交换
			}
		}
		if tag {
			break
		}
		length-- //因为第一次循环完，列表的最右边已经最大了，所以再比较它之前的数
	}
	//相当于从左往右取相邻两数比较，小的放左边大的放右边，保证每一次循环完后最小的放最左边，然后省略这个数，减小范围继续比较
}

//SelectSort 选择排序
func SelectSort(list []int) {
	length := len(list)
	for i := 0; i < length-1; i++ { //从左往右取待比较的数
		for j := i + 1; j < length; j++ { //从左往右取待比较的数的右边的第一个数
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i] //如果左边大于右边，交换，然后继续循环
			}
		}
		//循环完后，保证最左边的已经是最小了，所以下一次时，i++
	}
	//相当于从左往右依次取数，与其右边的所有数依次比较，找出最小的数放在该位置
}

//InsertSort 插入排序
func InsertSort(list []int) {
	length := len(list)
	//从右往左比较
	for i := 1; i < length; i++ {
		temp := list[i]     //待插入的数
		j := i - 1          //相邻的数,temp从右往左的第一个数
		if list[j] > temp { //如果左边的数大于它
			for j >= 0 && list[j] > temp {
				list[j+1] = list[j] //将左边的数右移一位
				j--                 //再与左边的左边进行比较,一直循环比较到第一个或者不满足条件为止
			}
			list[j+1] = temp //此时j+1的位置就是temp改插入的地方
			//为什么是j+1，因为list[j]<temp<list[j+1],自然要插入到j+1的位置
			//而且由于list[j+1] = list[j]这个操作，实际上满足上述条件时，list[j+2]=list[j+1],所以给list[j+1]重新赋值
		}
	}
}

//ShellSort 希尔排序
func ShellSort(list []int) {
	length := len(list)
	//分组插入，每次取列表长度的一半，之后依次减半，直到为1
	for d := length / 2; d >= 1; d = d / 2 {
		//接下来和插入排序一样
		for i := d; i < length; i += d {
			temp := list[i]
			j := i - d
			if list[j] > temp {
				for j >= 0 && list[j] > temp {
					list[j+d] = list[j]
					j = j - d
				}
				list[j+d] = temp
			}
		}
	}
}

//MergeSort1 自顶向下的归并排序(递归)
func MergeSort1(array []int, begin, end int) {
	if end-begin > 1 {
		mid := (end-begin+1)/2 + begin
		//左边归并
		MergeSort1(array, begin, mid)
		//右边归并
		MergeSort1(array, mid, end)
		//合并
		merge(array, begin, mid, end)
	}
}

//MergeSort2 自底向上的归并排序(非递归)
func MergeSort2(array []int, begin, end int) {
	step := 1 //步长，先一个个的合并，然后再两个两个合并，再四个四个...
	for end-begin > step {
		//开始切分数组
		for i := begin; i < end; i = i + step*2 {
			left := i           //左下标
			mid := i + step     //中间下标
			right := i + 2*step //右下标
			//相当于以此隔离出两个数组进行归并
			//[left,mid),[mid,right)
			if mid > end {
				//表示不存在第二个数组了,中止这次循环
				break
			}
			if right > end {
				//防止right超界
				right = end
			}
			//归并
			merge(array, left, mid, right)
		}
		//循环完一次，步长加倍
		step *= 2
	}
}

//merge 合并操作，注意：[begin,mid) [mid,end)
func merge(array []int, begin, mid, end int) {
	//根据begin,mid,end，先模拟出两个要比较的数组长度
	left := mid - begin   //左边数组长度
	right := end - mid    //右边数组长度
	length := end - begin //总长
	//新建辅助数组
	temp := make([]int, 0, length)
	//新建左右比较数组的下标
	l, r := 0, 0
	//开始比较
	for l < left && r < right {
		if array[begin+l] < array[mid+r] {
			temp = append(temp, array[begin+l])
			l++
		} else {
			temp = append(temp, array[mid+r])
			r++
		}
	}
	//将剩余的添加进去
	if l == left {
		//表示左边的已经都添加进temp了,右边的还有剩的
		temp = append(temp, array[mid+r:end]...)
	} else {
		//表示右边的已经都添加进temp了,左边的还有剩的
		temp = append(temp, array[begin+l:mid]...)
	}
	//赋值到原数组
	for i := 0; i < length; i++ {
		array[begin+i] = temp[i]
	}
}

//HeapSort1 堆排序1
func HeapSort1(array []int) {
	h := tree.NewHeap(array)
	for _, v := range array {
		h.Push(v)
	}
	for range array {
		h.Pop()
	}
}

//HeapSort 堆排序2 自底向上
func HeapSort2(array []int) {
	//堆的大小
	count := len(array)

	//最底层的叶子节点的下标
	start := count/2 + 1

	//结束点
	end := count - 1

	//从最底层开始交换
	for start >= 0 {
		sift(array, start, count)
		//向左偏移一个节点，如果该层没有就表示是上一层最右边的节点
		start--
	}

	//排序
	for end > 0 {
		//将堆顶元素与堆尾元素互换
		array[end], array[0] = array[0], array[end]
		//重新交换
		sift(array, 0, end)
		end--
	}
}

//sift 下沉操作
func sift(array []int, start, count int) {
	//父节点
	root := start

	//左子节点
	left := root*2 + 1

	//交换
	for left < count {
		//如果右子节点比左子节点大，那么指向右子节点，count-left>1表示不要超过范围.
		//left,right再[start,count-1]
		if count-left > 1 && array[left] < array[left+1] {
			left++
		}

		//如果父节点小于子节点，交换
		if array[root] < array[left] {
			array[root], array[left] = array[left], array[root]
			//之后父节点变为子节点，继续比较
			root = left
			left = root*2 + 1
		} else {
			return
		}
	}
}

//QuickSort 快速排序
func QuickSort(list []int) {
	quickSort(list, 0, len(list)-1)
}

//quickSort 快排递归
func quickSort(list []int, begin, end int) {
	if begin < end {
		//先分区
		temp := part(list, begin, end)
		//左边递归
		quickSort(list, begin, temp-1)
		//右边递归
		quickSort(list, temp+1, end)
	}
}

//part 分区
func part(list []int, begin, end int) int {
	//begin是基准数
	i := begin + 1 //第一个[]，也就是基准数右边第一个
	j := end       //第二个[],也就是最后一个

	for i < j {
		//当第一个[]大于基准数，就与第二个[]交换
		if list[i] > list[begin] {
			list[i], list[j] = list[j], list[i]
			j-- //第二个[]左移
		} else {
			i++ //第一个[]右移
		}
	}

	if list[i] >= list[begin] {
		//表示此时i所处的值不小于基准数，而i的左边的值是小于基准数的
		i--
	}
	list[begin], list[i] = list[i], list[begin]
	return i
}

//RadixSort 基数排序（桶排序）
func RadixSort(list []int) {
	//新建一个二维数组
	radix := make([][]int, 10)
	for i := 0; i < 10; i++ {
		radix[i] = make([]int, 0)
	}
	//找出最大位数
	max := 0
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	//从个位开始遍历，直到达到最大位为止
	div := 1
	for max/div > 0 {
		//遍历列表
		for _, v := range list {
			//算当前位的大小
			temp := v / div % 10
			//放入对应桶中
			radix[temp] = append(radix[temp], v)
		}
		//遍历桶，取出数据
		index := 0
		for i := 0; i < 10; i++ {
			for _, v := range radix[i] {
				list[index] = v
				index++
			}
			//同时将桶的数据清空
			radix[i] = make([]int, 0)
		}
		//放大div,从下一位开始
		div *= 10
	}
}

//TestSort 测试排序
func TestSort() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}
	//BubbleSort(list)
	//SelectSort(list)
	//InsertSort(list)
	//ShellSort(list)
	//MergeSort1(list, 0, len(list))
	//MergeSort2(list, 0, len(list))
	//HeapSort1(list)
	//HeapSort2(list)
	//QuickSort(list)
	RadixSort(list)
	fmt.Println(list)
}
