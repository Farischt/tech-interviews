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

// Time Complexity O(n) because we are linearly iterating over the linkedlist.
// Space Complexity O(1) because we are not creating any new data structure.
func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	currentNode := head
	fasterNode := head.Next

	for fasterNode != nil && fasterNode.Next != nil {
		if currentNode == fasterNode {
			return true
		}
		currentNode = currentNode.Next
		fasterNode = fasterNode.Next.Next
	}

	return false
}

// Time Complexity O(n) because we are linearly iterating over the linkedlist.
// Space Complexity O(n) because we are creating a map.
func HasCycleBis(head *ListNode) bool {
	if head == nil {
		return false
	}

	visitedNodes := make(map[ListNode]int)
	currentNode := head

	for currentNode.Next != nil {

		visitedNodes[*currentNode] += 1
		if visitedNodes[*currentNode] > 1 {
			return true
		}
		currentNode = currentNode.Next
	}

	return false
}
