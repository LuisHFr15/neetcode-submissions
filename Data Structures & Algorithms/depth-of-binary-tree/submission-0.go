/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	return deepSearch(root, 0)
}

func deepSearch(root *TreeNode, count int) int {
	if root == nil {
		return count
	}
	count += 1
	return max(deepSearch(root.Left, count), deepSearch(root.Right, count))
}
