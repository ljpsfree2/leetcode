package tasks

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var node *ListNode
	var length int = 0
	var step, stop int
	var result []*ListNode

	node = head
	for node.Next {
		length = length + 1
		node = node.Next
	}

	step = length / k
	stop = length % k
	result = append(result, head)
	for node.Next {
		if j < stop {
			for i := 0; i < step+1; i++ {
				node = node.Next
			}
			j = j + 1
		} else {
			for i := 0; i < step; i++ {
				node = node.Next
			}
			j = j + 1
		}
	}

}
