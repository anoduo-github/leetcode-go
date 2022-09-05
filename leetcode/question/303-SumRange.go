package question

//NumArray 数组
type NumArray struct {
	Nums []int
}

//Constructor 结构参数
func Constructor(nums []int) NumArray {
	numsArray := new(NumArray)
	numsArray.Nums = nums
	return *numsArray
}

//SumRange 求和
func (na *NumArray) SumRange(i int, j int) int {
	sum := 0
	for i <= j {
		sum = sum + na.Nums[i]
		i++
	}
	return sum
}
