package two_pointers

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
