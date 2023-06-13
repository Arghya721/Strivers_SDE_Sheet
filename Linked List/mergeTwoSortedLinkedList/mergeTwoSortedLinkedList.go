/* You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.

*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := &ListNode{
		Val:  0,
		Next: nil,
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		ans = list1
		ans.Next = mergeTwoLists(list1.Next, list2)
	} else {
		ans = list2
		ans.Next = mergeTwoLists(list1, list2.Next)
	}

	return ans
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	var head *ListNode
	var node1 *ListNode
	var node2 *ListNode
	var node3 *ListNode
	var node4 *ListNode
	var node5 *ListNode
	var node6 *ListNode
	var node7 *ListNode
	var node8 *ListNode
	var node9 *ListNode

	head = &ListNode{Val: 1, Next: nil}
	node1 = &ListNode{Val: 2, Next: nil}
	node2 = &ListNode{Val: 4, Next: nil}
	node3 = &ListNode{Val: 1, Next: nil}
	node4 = &ListNode{Val: 3, Next: nil}
	node5 = &ListNode{Val: 4, Next: nil}
	node6 = &ListNode{Val: 5, Next: nil}
	node7 = &ListNode{Val: 6, Next: nil}
	node8 = &ListNode{Val: 7, Next: nil}
	node9 = &ListNode{Val: 8, Next: nil}

	head.Next = node1
	node1.Next = node2

	node3.Next = node4
	node4.Next = node5

	node6.Next = node7
	node7.Next = node8
	node8.Next = node9

	printLinkedList(head)
	printLinkedList(node3)

	mergedList := mergeTwoLists(head, node3)
	printLinkedList(mergedList)
}
