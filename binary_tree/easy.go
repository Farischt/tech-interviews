package binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* ----------------------- Invert Binary Tree ----------------------- */

// Time Complexity O(n)
// Space Complexity O(n)
func InvertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	temp := root.Right
	root.Right = root.Left
	root.Left = temp

	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}

/* ----------------------- Maximum Depth of Binary Tree ----------------------- */

// Time Complexity O(n)
// Space Complexity O(n)
func MaxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	return int(math.Max(float64(1+leftDepth), float64(1+rightDepth)))
}

/* ----------------------- Is same tree ----------------------- */

// Time Complexity O(n)
// Space Complexity O(n)
func IsSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !IsSameTree(p.Left, q.Left) {
		return false
	}

	if !IsSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
