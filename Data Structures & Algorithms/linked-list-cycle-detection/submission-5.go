/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	} else if head.Next == nil {
		return false
	}
    slow := head
	fast := head

	for slow.Next != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next == nil {
			break
		} else {
			fast = fast.Next
		}
		if slow == fast {
			return true
		}
	}

	return false
}
