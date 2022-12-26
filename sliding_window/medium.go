package sliding_window

import (
	"math"
	"reflect"
)

// LenghtOfLonguestSubstringWithoutRepetition returns the lenght of the longuest substring in a string without character repetition.
// Technique used is sliding window.
// Time Complexity O(n)
// Space Complexity O(n)
// Difficulty: Medium
func LenghtOfLonguestSubstringWithoutRepetition(s string) int {
	var res int
	visitedChar := make(map[byte]int)

	left := 0
	right := 0

	for right < len(s) {
		rightCharacter := s[right]
		visitedChar[rightCharacter]++

		for visitedChar[rightCharacter] > 1 {
			leftCharacter := s[left]
			visitedChar[leftCharacter]--
			left++
		}

		res = int(math.Max(float64(res), float64(right-left+1)))
		right++
	}

	return res
}

/* -------------------------- Character Replacement ------------------------- */
// Time Complexity O(n)
// Space Complexity O(n)
// Difficulty: Medium
func getMostFrequentChar(f map[byte]int) byte {
	count := 0
	var currentByte byte
	for key, value := range f {
		if value > count {
			count = value
			currentByte = key
		}
	}

	return currentByte
}

func CharacterReplacement(s string, k int) int {
	charCount := make(map[byte]int)
	left := 0
	res := 0

	for right := range s {
		charCount[s[right]]++
		mostFrequentChar := getMostFrequentChar(charCount)
		windowLen := right - left + 1

		if (windowLen - charCount[mostFrequentChar]) > k {
			// it is invalid so we need to move the left pointer
			charCount[s[left]]--
			left++
			windowLen--
		}

		if windowLen > res {
			res = windowLen
		}
	}

	return res
}

// Time Complexity O(n)
// Space Complexity O(n)
// Difficulty: Medium
func CheckInclusion(s1 string, s2 string) bool {

	s1Count := make(map[string]int)
	s2Count := make(map[string]int)

	for _, char := range s1 {
		s1Count[string(char)]++
	}

	maxLen := len(s1)
	left := 0

	for right := range s2 {
		s2Count[string(s2[right])]++

		if right-left+1 > maxLen {
			s2Count[string(s2[left])]--

			if s2Count[string(s2[left])] == 0 {
				delete(s2Count, string(s2[left]))
			}

			left++
		}

		if reflect.DeepEqual(s1Count, s2Count) {
			return true
		}
	}
	return false
}
