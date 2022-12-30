package arrays_hash

import "fmt"

/* Exercise: --------------------------------- Group anagram --------------------------------- */

// GroupAnagrams returns a 2D list of grouped anagrams.
// Time Complexity O(n * m) with n the lenght of the array and m the average lenght of each word
// Space Complexity O(n) with n the number of element in the array
// Difficulty: Medium
// Using a map of the values storign their indexes.
func GroupAnagrams(strs []string) [][]string {
	var outputResult [][]string = [][]string{}
	groupedAnagrams := make(map[[26]int][]string)

	for _, word := range strs {
		// Create the count result of the currentword
		count := [26]int{}
		// For each character of the word i will update the count array
		for _, char := range word {
			count[int(char)-int('a')] = count[int(char)-int('a')] + 1
		}

		groupedAnagrams[count] = append(groupedAnagrams[count], word)

	}

	for _, element := range groupedAnagrams {
		outputResult = append(outputResult, element)
	}

	return outputResult
}

// TopKMostFrequent returns the k most frequents nums in an array.
// Use an hashmap to count each occurences.
// Map the hashmap and store each value in a 2D array as follow
// Index -> the number of occurence
// Value -> The element
// Time Complexity O(n). The loop at line 158 is constant time loop.
// Space Complexity O(n)
// Difficulty: Medium
func TopKMostFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	scores := make([][]int, len(nums)+1)
	outputResult := []int{}

	for _, el := range nums {
		count[el]++
	}

	for number, occurences := range count {
		scores[occurences] = append(scores[occurences], number)
	}

	for i := len(scores) - 1; i > 0; i-- {
		for _, el := range scores[i] {
			outputResult = append(outputResult, el)

			if len(outputResult) == k {
				return outputResult
			}
		}
	}

	return outputResult
}

// Usage of prefix and postfix values.
// Could be done in place directyl in order to save space complexity.
// Time Complexity O(n)
// Space Complexity O(n)
// Difficulty: Medium
func ProductExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))

	// prefix is the multiplication of all the numbers that come before i.
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefix[i] = nums[i]
		} else {
			prefix[i] = prefix[i-1] * nums[i]
		}
	}

	// postfix is the multiplication of all the numbers that come after i.
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			postfix[i] = nums[i]
		} else {
			postfix[i] = postfix[i+1] * nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefixValue := 1
			answer[i] = prefixValue * postfix[i+1]
		} else if i == len(nums)-1 {
			postfixValue := 1
			answer[i] = prefix[i-1] * postfixValue
		} else {
			answer[i] = prefix[i-1] * postfix[i+1]
		}
	}

	return answer
}

// Doing it this way we cannot save space complexity.
func ProductExceptSelfBis(nums []int) []int {

	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	output := make([]int, len(nums))

	sum := 1
	for i := range nums {
		if i == 0 {
			prefix[i] = 1
		} else {
			sum = sum * nums[i-1]
			prefix[i] = sum
		}
	}

	sum = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			suffix[i] = sum
		} else {
			sum = sum * nums[i+1]
			suffix[i] = sum
		}
	}

	for i := range nums {
		output[i] = prefix[i] * suffix[i]
	}

	return output
}

// Time Complexity O(n)
// Space Complexity O(n)
// Difficulty: Medium
func IsValidSudoku(board [][]byte) bool {
	// Init
	rowCount := make(map[int]map[int]int)
	colCount := make(map[int]map[int]int)
	boxCount := make(map[[2]int]map[int]int)

	for row := 0; row < 9; row++ {

		// Init row nested map
		rowCount[row] = make(map[int]int)
		for col := 0; col < 9; col++ {

			// Init col nested map
			colCount[col] = make(map[int]int)

			// Init box nested map
			if col < 3 && row < 3 {
				boxCount[[2]int{row, col}] = make(map[int]int)
			}

			// In char is a dot we go for next iteration
			if board[row][col] != 46 {
				continue
			}

			// Increment the count of visited number
			rowCount[row][int(board[row][col])]++
			colCount[col][int(board[row][col])]++
			boxCount[[2]int{row / 3, col / 3}][int(board[row][col])]++

			// Check if a number is already visited and return false if true
			if rowCount[row][int(board[row][col])] > 1 || colCount[col][int(board[row][col])] > 1 || boxCount[[2]int{row / 3, col / 3}][int(board[row][col])] > 1 {
				fmt.Printf("row=%d col=%d \n", row, col)
				return false
			}

		}
	}
	return true
}
