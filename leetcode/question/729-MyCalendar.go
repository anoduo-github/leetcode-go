package question

type MyCalendar struct {
	date [][]int //行程时间的区间
	size int     //行程数
}

func Constructor2() MyCalendar {
	return MyCalendar{
		date: make([][]int, 0),
		size: 0,
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for i := 0; i < this.size; i++ {
		if this.date[i][0] < end && this.date[i][1] > start {
			return false
		}
	}
	this.date = append(this.date, []int{start, end})
	this.size++
	return true
}
