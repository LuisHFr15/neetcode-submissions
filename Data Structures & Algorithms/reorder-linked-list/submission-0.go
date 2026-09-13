/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
    l1 := head
	l2 := head.Next

	for l2 != nil && l2.Next != nil {
		l1 = l1.Next
		l2 = l2.Next.Next
	}

	l2 = l1.Next
	l1.Next = nil
	var prev *ListNode
	for l2 != nil {
		aux := l2.Next
		l2.Next = prev
		prev = l2
		l2 = aux
	}

	l1 = head
	for prev != nil {
		tmp1, tmp2 := l1.Next, prev.Next
        l1.Next = prev
        prev.Next = tmp1
        l1, prev = tmp1, tmp2
	}
}
