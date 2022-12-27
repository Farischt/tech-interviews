package stack

/*
link: https://leetcode.com/problems/valid-parentheses/
Time Complexity O(n)
Space Complexity O(n)

Explaination: use of a stack to check if closing parenthesis has an opening parenthesis.
The order counts !
*/
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}
	correspondence := make(map[rune]rune)
	correspondence['}'] = '{'
	correspondence[']'] = '['
	correspondence[')'] = '('

	for _, char := range s {
		_, present := correspondence[char]
		// If it exists, the current char is a closing element
		if present {

			// If the stack is empty, it means the first element is a closing parenthesis so we need to return false
			if len(stack) == 0 {
				return false
			}

			// Pop from stack and compare
			lastElement := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if correspondence[char] != lastElement {
				return false
			}

			// In this case, the current char is an opening char so we need to add it to the stack
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
