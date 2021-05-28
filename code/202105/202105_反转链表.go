/**
* @Author: Chicken dishes
* @Date: 2021/5/26 20:00
**/

package main

/**
 * 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	node := head
	var prev *ListNode
	for node != nil {
		temp := node.Next
		node.Next = prev
		prev = node
		node = temp
	}
	return prev

}

func main() {
	var listNode = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}
	reverseList(&listNode)
}
