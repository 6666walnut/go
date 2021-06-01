/**
* @Author: Chicken dishes
* @Date: 2021/5/26 20:42
**/

package main

/**
 * 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var mergeList = &ListNode{}
	node := mergeList
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return mergeList.Next
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
	var listNode1 = ListNode{
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
	mergeTwoLists(&listNode, &listNode1)
}
