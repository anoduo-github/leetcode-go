package primary

//ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * LC合并两个有序链表
 * start
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	arr := make([]int, 0)
	for l1 != nil {
		arr = append(arr, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		arr = append(arr, l2.Val)
		l2 = l2.Next
	}
	res := quickSort(arr)
	if len(res) < 1 {
		return nil
	}
	head := &ListNode{}
	temp := head
	for i, v := range res {
		temp.Val = v
		if i < len(res)-1 {
			l := &ListNode{}
			temp.Next = l
			temp = temp.Next
		}
	}
	return head
}

func quickSort(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}
	low := make([]int, 0)
	high := make([]int, 0)
	mid := make([]int, 0)
	mid = append(mid, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[0] {
			high = append(high, arr[i])
		} else if arr[i] < arr[0] {
			low = append(low, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low = quickSort(low)
	high = quickSort(high)
	res := append(append(low, mid...), high...)
	return res
}

/**
 * LC合并两个有序链表
 * end
 */

//LC回文链表
func isPalindrome(head *ListNode) bool {
	sum := make([]int, 0)
	for head != nil {
		sum = append(sum, head.Val)
		head = head.Next
	}
	if len(sum) < 1 {
		return false
	}
	l := len(sum)
	for i := 0; i < l/2; i++ {
		if sum[i] != sum[l-i-1] {
			return false
		}
	}
	return true
}

//LC环形链表
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]int)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = 1
		head = head.Next
	}
	return false
}
