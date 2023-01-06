package binary_search

import "fmt"

/* ------------------------------ Binary search ----------------------------- */

// Binary search is a search algorithm that finds the position of a target value within a sorted array.
// Binary search compares the target value to the middle element of the array.
// If they are not equal, the half in which the target cannot lie is eliminated and the search continues on the remaining half,
// again taking the middle element to compare to the target value, and repeating this until the target value is found.
// If the search ends with the remaining half being empty, the target is not in the array.
// Time complexity: O(log n)
// Space complexity: O(1)
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (right + left) / 2
		fmt.Println(middle)
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}

	return -1
}
