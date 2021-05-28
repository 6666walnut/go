/**
* @Author: Chicken dishes
* @Date: 2021/5/26 21:02
**/

package main

/**
 * 请完成一个函数，输入一个二叉树，该函数输出它的镜像。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func main() {

	tree := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 7,
		},
	}
	mirrorTree(&tree)
}
