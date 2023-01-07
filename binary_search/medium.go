package binary_search

import "math"

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

/* ------------------ Find Minimum In Rotated Sorted Array ------------------ */

// Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:
// - [4,5,6,7,0,1,2] if it was rotated 4 times.
// - [0,1,2,4,5,6,7] if it was rotated 7 times.
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
// Given the sorted rotated array nums of unique elements, return the minimum element of this array.
// You must write an algorithm that runs in O(log n) time.

// Time complexity: O(log n)
// Space complexity: O(1)
func FindMinimumInRotatedSortedArray(nums []int) int {
	res := math.MaxInt
	left := 0
	right := len(nums) - 1

	for left <= right {
		if nums[left] < nums[right] {
			res = int(math.Min(float64(res), float64(nums[left])))
			break
		}

		middle := (left + right) / 2
		res = int(math.Min(float64(res), float64(nums[middle])))

		if nums[middle] >= nums[left] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return res
}

/* ------------------ Search In RotatedSortedArray ------------------ */

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
// You are given a target value to search. If found in the array return its index, otherwise return -1.
// You may assume no duplicate exists in the array.
// Your algorithm's runtime complexity must be in the order of O(log n).

// Time complexity: O(log n)
// Space complexity: O(1)
func SearchInRotatedSortedArray(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		}

		// Left side is sorted
		if nums[left] <= nums[middle] {
			// Check if the target is in its range
			if nums[left] <= target && target <= nums[middle] {
				right = middle - 1
			} else {
				// Go to the other range
				left = middle + 1
			}
		} else {
			if nums[middle] <= target && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return -1
}
