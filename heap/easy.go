package heap_queue

import "container/heap"

/* --------------------- Kth Largest Element in a Stream -------------------- */

/*
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.
Implement KthLargest class:
KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.
*/

// Time Complexity O(n)
// Space Complexity O(n)
type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	instance := KthLargest{h: (*IntHeap)(&nums), k: k}
	heap.Init(instance.h)
	for len(*instance.h) > k {
		heap.Pop(instance.h)
	}
	return instance
}

func (kth *KthLargest) Add(val int) int {
	heap.Push(kth.h, val)

	if len(*kth.h) > kth.k {
		heap.Pop(kth.h)
	}

	return (*kth.h)[0]
}
