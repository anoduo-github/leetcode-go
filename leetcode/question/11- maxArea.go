package question

//11. 盛最多水的容器
func maxArea(height []int) int {
	n := len(height)
	if n == 0 || n == 1 {
		return 0
	}
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		if height[i] > height[j] {
			temp := height[j] * (j - i)
			if max < temp {
				max = temp
			}
			j--
		} else {
			temp := height[i] * (j - i)
			if max < temp {
				max = temp
			}
			i++
		}
	}
	return max
}
