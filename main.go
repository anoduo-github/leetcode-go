package main

import "fmt"

const (
	z = "zz"
	k
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(cap(nums))
}
