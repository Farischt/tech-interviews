package two_pointers

import "math"

/* ----------------------- Two Sum II Array is sorted ----------------------- */

// Time Complexity O(n)
// Space Complexity O(1)
func TwoSum(numbers []int, target int) []int {

	left := 0
	right := len(numbers) - 1

	for left < right {

		if numbers[left]+numbers[right] > target {
			// We need to move the right pointer
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}

	}

	return []int{left + 1, right + 1}
}

/* ----------------------- Container With Most Water ----------------------- */
// Time Complexity O(n)
// Space Complexity O(1)
func MaxArea(height []int) int {
	maximumArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		area := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		maximumArea = int(math.Max(float64(maximumArea), float64(area)))

		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}

	return maximumArea
}
