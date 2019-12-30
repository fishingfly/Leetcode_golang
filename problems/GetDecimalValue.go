package problems

type ListNode struct {
	Val int
	Next *ListNode
}
func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}
	result := 0
	for head != nil {
		result = result << 1
		if head.Val == 1 {
			result += 1
		}
		head = head.Next
	}
	return result
}
