/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := new(ListNode)
	last := head

	for list1 != nil || list2 != nil {
		aux := new(ListNode)
		if list1 == nil {
			aux.Val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			aux.Val = list1.Val
			list1 = list1.Next
		} else if list1.Val <= list2.Val {
			aux.Val = list1.Val
			list1 = list1.Next
		} else {
			aux.Val = list2.Val
			list2 = list2.Next
		}
		last.Next = aux
		last = aux
	}
	return head.Next
}