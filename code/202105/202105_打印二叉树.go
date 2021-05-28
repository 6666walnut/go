/**
* @Author: Chicken dishes
* @Date: 2021/5/27 20:46
**/

package main

/**
 * 从上到下打印二叉树 II
 */

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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	keys := [][]int{}
	for len(nodes) > 0 {
		layer := []int{}
		layerNode := []*TreeNode{}
		for _, v := range nodes {
			layer = append(layer, v.Val)
			if v.Left != nil {
				layerNode = append(layerNode, v.Left)
			}
			if v.Right != nil {
				layerNode = append(layerNode, v.Right)
			}
		}
		keys = append(keys, layer)
		nodes = layerNode
	}
	return keys
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	levelOrder(tree)
}
