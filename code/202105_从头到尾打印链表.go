/**
* @Author: Chicken dishes
* @Date: 2021/5/23 17:01
**/

package main

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
*/

type listNode struct {
	Val  int
	Next *listNode
}

func reversePrint(head *listNode) []int {
	var arr = make([]int, 0)
	node := head
	for node != nil {
		arr = append([]int{node.Val}, arr...)
		node = node.Next
	}

	return arr
}

func main() {
	var listNode = listNode{
		Val: 3,
		Next: &listNode{
			Val: 2,
			Next: &listNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	reversePrint(&listNode)
}
