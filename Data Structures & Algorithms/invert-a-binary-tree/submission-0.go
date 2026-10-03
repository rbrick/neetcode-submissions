/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    
	if root == nil {
		return nil
	}
	left, right := root.Left, root.Right

	root.Left = right
	root.Right = left

	invertTree(left)
	invertTree(right)


    return root
}
