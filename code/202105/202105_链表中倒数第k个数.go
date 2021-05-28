/**
* @Author: Chicken dishes
* @Date: 2021/5/26 19:40
**/

package main

/**
 * 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
 * 例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getKthFromEnd(head *ListNode, k int) *ListNode {
	node := head
	var num = 0
	for node != nil {
		node = node.Next
		num++
	}
	for i := 0; i < num-k; i++ {
		head = head.Next
	}
	return head
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
	getKthFromEnd(&listNode, 4)
}
