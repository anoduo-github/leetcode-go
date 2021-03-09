package question

//1512. 好数对的数目
func numIdenticalPairs(nums []int) int {
	sum := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				sum++
			}
		}
	}
	return sum
}
