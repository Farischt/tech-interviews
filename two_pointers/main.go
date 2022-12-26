package two_pointers

import (
	"regexp"
	"strings"
)

/* Exercise: ------------------------------ Is palindrome ----------------------------- */

// IsPalindrome returns true if a string is a palindrome. Non alphanumeric characters are skipped.
// Time Complexity O(n^2)
// Space Complexity O(n)
// Difficulty: Easy
func IsPalindrome(s string) bool {
	right := len(s) - 1
	left := 0

	var validLetter = regexp.MustCompile(`[a-zA-Z0-9]`)

	for left < right {
		for left < right && !validLetter.MatchString(string(s[left])) {
			left++
		}

		for left < right && !validLetter.MatchString(string(s[right])) {
			right--
		}

		l := strings.ToLower(string(s[left]))
		r := strings.ToLower(string(s[right]))
		if l != r {
			return false
		}

		right--
		left++
	}

	return true
}
