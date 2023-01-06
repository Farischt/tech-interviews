package binary_search

/* --------------------------- Search a 2D Matrix --------------------------- */

// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
// - Integers in each row are sorted from left to right.
// - The first integer of each row is greater than the last integer of the previous row.

// Time complexity: O(n log m)
// Space complexity: O(1)
func SearchMatrix(matrix [][]int, target int) bool {
	for i := range matrix {
		left := 0
		right := len(matrix[i]) - 1

		if target < matrix[i][left] {
			return false
		} else if target == matrix[i][left] || target == matrix[i][right] {
			return true
		} else if target > matrix[i][right] {
			continue
		}

		for left <= right {
			middle := (left + right) / 2
			if matrix[i][middle] > target {
				right = middle - 1
			} else if matrix[i][middle] < target {
				left = middle + 1
			} else {
				return true
			}
		}
	}

	return false
}
