/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countDepth(root *TreeNode, count int) int {
	if root == nil {
		return count
	}

	dR := countDepth(root.Right, count+1)
	dL := countDepth(root.Left, count+1)

	if dR > dL {
		return dR
	}

	return dL

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

    rightDepth, leftDepth := countDepth(root.Right, 1), countDepth(root.Left, 1)

	if rightDepth > leftDepth {
		return rightDepth
	}

	return leftDepth
}
