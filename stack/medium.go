package stack

import (
	"fmt"
	"math"
	"strconv"
)

/*
link https://leetcode.com/problems/min-stack/description/
Time Complexity O(1)
Space Complexity
WARNING: this MinStack only works with the test cases of leet code
As you can see some edge cases are not handle in the best way
*/

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
type MinStack struct {
	storage            []int
	recordsOfMinValues []int
}

func Constructor() MinStack {
	return MinStack{
		storage:            []int{},
		recordsOfMinValues: []int{math.MaxInt},
	}
}

func (ms *MinStack) Push(val int) {
	ms.storage = append(ms.storage, val)

	if val <= ms.recordsOfMinValues[len(ms.recordsOfMinValues)-1] {
		ms.recordsOfMinValues = append(ms.recordsOfMinValues, val)
	}
}

func (ms *MinStack) Pop() {
	if len(ms.storage) == 0 {
		fmt.Println("stack is empty, cannot pop")
		return
	}

	topElement := ms.Top()
	if topElement == ms.recordsOfMinValues[len(ms.recordsOfMinValues)-1] {
		ms.recordsOfMinValues = ms.recordsOfMinValues[:len(ms.recordsOfMinValues)-1]
	}

	ms.storage = ms.storage[:len(ms.storage)-1]

}

func (ms *MinStack) Top() int {
	if len(ms.storage) == 0 {
		fmt.Println("stack is empty, cannot get top")
		return math.MaxInt
	}
	return ms.storage[len(ms.storage)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.recordsOfMinValues[len(ms.recordsOfMinValues)-1]
}

/* -------------------- Evaluate Reverse Polish Notation -------------------- */
/*
link: https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
Time Complexity O(n)
Space Complexity O(n)
*/

func EvalRPN(tokens []string) int {
	stack := []int{}

	for _, element := range tokens {

		if !isElementOperator(element) {
			number, _ := strconv.Atoi(element)
			stack = append(stack, number)
		} else {
			// In this case the element is an operator
			// So we need to pop the 2 top elements of the stack
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Compute them (beware to divide in the good order)
			res := compute(left, element, right)
			stack = append(stack, res)
			// Add the result to the top of the stack
		}

	}

	return stack[len(stack)-1]
}

func isElementOperator(element string) bool {
	return element == "+" || element == "-" || element == "*" || element == "/"
}

func compute(left int, operator string, right int) int {
	if operator == "+" {
		return left + right
	} else if operator == "-" {
		return left - right
	} else if operator == "*" {
		return left * right
	} else {
		return left / right
	}
}

/* ---------------------------- Daily temperature --------------------------- */

type record struct {
	temperature int
	day         int
}

func DailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []record{}

	for idx, temperature := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1].temperature < temperature {
			// pop from the stack
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top.day] = idx - top.day
		}

		stack = append(stack, record{
			temperature: temperature,
			day:         idx,
		})
	}

	return result
}
