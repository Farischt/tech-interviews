package arrays_hash

import (
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
