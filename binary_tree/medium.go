package binary_tree

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
