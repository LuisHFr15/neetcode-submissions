/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	invertNodes(root)
	return root
}

func invertNodes(root *TreeNode) {
	if root == nil {
		return
	}
	tmp1, tmp2 := root.Left, root.Right
	root.Left = tmp2
	root.Right = tmp1
	invertNodes(root.Left)
	invertNodes(root.Right)
}
