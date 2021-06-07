/**
* @Author: Chicken dishes
* @Date: 2021/6/5 16:00
**/

package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {

	valArr := make([]int, 0)
	traverse(root, &valArr)
	return valArr[k-1]
}

func traverse(root *TreeNode, arr *[]int) {

	if root == nil {
		return
	}
	traverse(root.Right, arr)
	*arr = append(*arr, root.Val)
	traverse(root.Left, arr)
}
func main() {
	tree := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	c := kthLargest(&tree, 2)
	fmt.Println(c)
}
