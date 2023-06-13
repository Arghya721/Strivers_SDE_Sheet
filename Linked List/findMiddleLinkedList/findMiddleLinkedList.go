/* Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

Example 1:

Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:

Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
 

Constraints:

The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
*/ 

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    var slow,fast *ListNode
    slow = head
    fast = head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow
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

	head = &ListNode{Val: 1}
	node1 = &ListNode{Val: 2}
	node2 = &ListNode{Val: 3}
	node3 = &ListNode{Val: 4}
	node4 = &ListNode{Val: 5}
	node5 = &ListNode{Val: 6}
	node6 = &ListNode{Val: 7}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	printLinkedList(head)
	middle := middleNode(head)
	fmt.Println("Middle node: ", middle.Val)
}