package question

//1450. 在既定时间做作业的学生人数
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	max := 0
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			max++
		}
	}
	return max
}
