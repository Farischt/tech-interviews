package binary_tree

import "math"

/* ----------------------- Binary Tree Level Order Traversal ----------------------- */

// Time Complexity O(n)
// Space Complexity O(n)
// Technique: BFS
func LevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	res := [][]int{}

	for len(queue) > 0 {
		queueLen := len(queue)
		level := []int{}

		for i := 0; i < queueLen; i++ {
			first := queue[0]
			queue = queue[1:]

			if first.Left != nil {
				queue = append(queue, first.Left)
			}

			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			level = append(level, first.Val)
		}

		res = append(res, level)
	}

	return res
}

/* ----------------------- Binary Tree Right Side View ----------------------- */

// Time Complexity O(n)
// Space Complexity O(n)
// Technique: BFS and keep track of the last element in the level
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		queueLen := len(queue)
		level := []int{}

		for i := 0; i < queueLen; i++ {
			firstElement := queue[0]
			level = append(level, firstElement.Val)
			queue = queue[1:]

			if firstElement.Left != nil {
				queue = append(queue, firstElement.Left)
			}

			if firstElement.Right != nil {
				queue = append(queue, firstElement.Right)
			}
		}

		rightMostElement := level[len(level)-1]
		res = append(res, rightMostElement)
	}

	return res
}

/* ----------------------- Valid Binary Search Tree ----------------------- */

// Time Complexity O(n)
// Space Complexity O(n)
func ValidNode(node *TreeNode, left int, right int) bool {
	if node == nil {
		return true
	}

	if !(node.Val < right && node.Val > left) {
		return false
	}

	return ValidNode(node.Left, left, node.Val) && ValidNode(node.Right, node.Val, right)
}

func IsValidBST(root *TreeNode) bool {
	return ValidNode(root, math.MinInt, math.MaxInt)
}

/* ----------------------- Count good nodes ----------------------- */

// Time Complexity O(n)
// Space Complexity O(n)
func GoodNodes(root *TreeNode) int {
	return GoodNodesUtil(root, root.Val)
}

func GoodNodesUtil(root *TreeNode, parent int) int {
	if root == nil {
		return 0
	}

	res := 1
	max := root.Val

	if parent > root.Val {
		res = 0
		max = parent
	}

	res += GoodNodesUtil(root.Left, max)
	res += GoodNodesUtil(root.Right, max)

	return res
}

/* ---------------------- Kth Smallest Element in a BST --------------------- */

// Time Complexity O(n)
// Space Complexity O(n)
func KthSmallest(root *TreeNode, k int) int {

	stack := []*TreeNode{}
	currentNode := root
	n := 0

	for {
		for currentNode != nil {
			stack = append(stack, currentNode)
			currentNode = currentNode.Left
		}

		currentNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n++

		if n == k {
			return currentNode.Val
		}

		currentNode = currentNode.Right
	}
}
