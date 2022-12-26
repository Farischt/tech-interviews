package arrays_hash

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

/* All the exercices are taken from https://neetcode.io/practice . */

/* Exercise: ---------------------- Contains duplicates interview --------------------- */

// ContainsDuplicate returns true if the input array contains duplicated values.
// Time Complexity O(n) because we are linearly iterating over the input array.
// Space Complexity O(n) because we create a new hashmap.
// Difficulty: Easy
func ContainsDuplicate(nums []int) bool {
	visitedValues := make(map[int]bool)

	for _, el := range nums {
		_, ok := visitedValues[el]
		if ok {
			return true
		} else {
			visitedValues[el] = true
		}
	}
	return false
}

/* Exercise: -------------------------- Is anagram interview -------------------------- */

// IsAnagram returns true if the word t is the anagram of the word s.
// Time Complexity O(n), with n the length of the word, because we are linearly looping over the word.
// Space Complexity O(k), with k the number of unique letters in the word. We are creating two hash map so, the time complexity is k + k = 2k = O(k).
// Difficulty: Easy
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		// If lengths are different it cannot be an anagram
		return false
	} else if s == t {
		// If the words are the same return true
		return true
	}

	// Storing the letters count of word s
	letters := make(map[byte]uint)
	// Storing the letters count of word t
	supposedAnagramLetters := make(map[byte]uint)

	// Since s & t have the same length, we can iterate over one loop
	for idx := range s {
		letters[s[idx]] += 1
		supposedAnagramLetters[t[idx]] += 1
	}

	return reflect.DeepEqual(letters, supposedAnagramLetters)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// IsAnagramBis returns true if the word t is the anagram of the word s.
// Time Complexity depends on the sorting algorithm used : worst case O(n^2) average case O(n * log(n)).
// Space Complexity also depends on the sorting algorithm used : it can ben O(1) or O(n).
// Difficulty: Easy
// In this case the time and space complexity is discussable.
func IsAnagramBis(s string, t string) bool {
	return SortString(s) == SortString(t)
}

/* Exercise: --------------------------------- Two sum --------------------------------- */

// Two sum return an array of the two index were the sum of the corresponding elements equals the target value.
// Time Complexity O(n^2).
// Space Complexity O(1). Since we create an array of two values at worst space complexity is 1 + 1 = 2 = O(1)
// Difficulty: Easy
func TwoSum(nums []int, target int) []int {
	for left := 0; left < len(nums)-1; left++ {
		for right := 1; right < len(nums); right++ {
			if nums[left]+nums[right] == target {
				return []int{left, right}
			}
		}
	}
	return []int{}
}

// Time Complexity O(n).
// Space Complexity O(n).
// Difficulty: Easy
// Using a map of the values storign their indexes.
func TwoSumBis(nums []int, target int) []int {
	mp := make(map[int]int) // key: number, val: index in nums

	for i, num := range nums {
		if idx, ok := mp[target-num]; ok {
			return []int{idx, i}
		}
		mp[num] = i
	}
	return []int{}
}

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
