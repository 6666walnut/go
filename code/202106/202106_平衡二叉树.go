/**
* @Author: Chicken dishes
* @Date: 2021/6/5 16:48
**/

package main

import "math"

/**
 * 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left)
	if l == -1 {
		return -1
	}

	r := dfs(root.Right)
	if r == -1 {
		return -1
	}

	if math.Abs(float64(l-r)) <= 1 {
		return int(math.Max(float64(l), float64(r))) + 1
	}
	return -1

}
func main() {

}
