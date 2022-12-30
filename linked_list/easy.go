package linked_list

/* ---------------------- Reverse a singly linked list. --------------------- */

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList returns the reversed linked list.
// Time Complexity O(n) because we are linearly iterating over the input array.
// Space Complexity O(1) because we are not creating any new data structure.
func ReverseList(head *ListNode) *ListNode {
	var previous *ListNode
	var current *ListNode = head
	var temp *ListNode

	for current != nil {
		temp = current.Next
		current.Next = previous

		previous = current
		current = temp
	}

	return previous

}
