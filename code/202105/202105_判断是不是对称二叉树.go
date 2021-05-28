/**
* @Author: Chicken dishes
* @Date: 2021/5/26 21:15
**/

package main

/**
 * 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsIsSymmetric(root.Left, root.Right)
}

func dfsIsSymmetric(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if (l == nil && r != nil) || (l != nil && r == nil) || (l.Val != r.Val) {
		return false
	}
	return dfsIsSymmetric(l.Left, r.Right) && dfsIsSymmetric(r.Left, l.Right)
}
