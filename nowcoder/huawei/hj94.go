package huawei

import "fmt"

//HJ94 记票统计
func HJ94() {
	var n int
	fmt.Scanln(&n)
	m := make(map[string]int)
	names := make([]string, 0)
	for i := 0; i < n; i++ {
		var name string
		fmt.Scan(&name)
		m[name] = 0
		names = append(names, name)
	}
	m["Invalid"] = 0
	var num int
	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		var s string
		fmt.Scan(&s)
		if _, ok := m[s]; ok {
			m[s]++
		} else {
			m["Invalid"]++
		}
	}
	for _, v := range names {
		fmt.Printf("%s : %d\n", v, m[v])
	}
	fmt.Printf("%s : %d\n", "Invalid", m["Invalid"])
}
