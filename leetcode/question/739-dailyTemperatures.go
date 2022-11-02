package question

//739. 每日温度
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, 0)
Label:
	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res = append(res, j-i)
				continue Label
			}
		}
		res = append(res, 0)
	}
	return res
}

/* func dailyTemperatures(temperatures []int) []int {
    length := len(temperatures)
    ans := make([]int, length)
    stack := []int{}
    for i := 0; i < length; i++ {
        temperature := temperatures[i]
        for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
            prevIndex := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans[prevIndex] = i - prevIndex
        }
        stack = append(stack, i)
    }
    return ans
} */
