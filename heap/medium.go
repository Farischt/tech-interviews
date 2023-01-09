package heap_queue

import "container/heap"

/* --------------------- Kth Closest Point To Origin -------------------- */

// Time: O(nlogk)
// Space: O(k)

type Point struct {
	distance   int
	coordinate []int
}

type MaxHeapPoint []Point

func (h MaxHeapPoint) Less(i, j int) bool {
	return h[i].distance > h[j].distance
}

func (h MaxHeapPoint) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeapPoint) Len() int {
	return len(h)
}

func (h *MaxHeapPoint) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeapPoint) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func ComputeDistance(x int, y int) int {
	origin := []int{0, 0}

	a := (origin[0] - x) * (origin[0] - x)
	b := (origin[1] - y) * (origin[1] - y)

	return a + b
}

// Use a max heap of len k

func KClosest(points [][]int, k int) [][]int {
	maxHeap := MaxHeapPoint{}
	res := [][]int{}

	for _, point := range points {
		distance := ComputeDistance(point[0], point[1])
		heap.Push(&maxHeap, Point{
			distance:   distance,
			coordinate: point,
		})
	}

	for len(maxHeap) > k {
		heap.Pop(&maxHeap)
	}

	for _, point := range maxHeap {
		res = append(res, point.coordinate)
	}

	return res
}
