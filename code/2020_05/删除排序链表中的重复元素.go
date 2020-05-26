/**
* @Author: Chicken dishes
* @Date: 2020/5/26 20:43
 */
package main

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node.Next = node.Next
		}
	}
	return head
}

func main() {
	var listNode ListNode = ListNode{
		Val:  1,
		Next: nil,
	}

	deleteDuplicates(&listNode)
}
