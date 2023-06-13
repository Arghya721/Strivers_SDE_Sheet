/*
	Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:

Input: head = [1,2]
Output: [2,1]
Example 3:

Input: head = []
Output: []

Constraints:

The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var next *ListNode
	next = nil
	var prev *ListNode
	prev = nil

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	head = prev
	return head
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{Val: 1, Next: nil}
	head.Next = &ListNode{Val: 2, Next: nil}
	head.Next.Next = &ListNode{Val: 3, Next: nil}
	head.Next.Next.Next = &ListNode{Val: 4, Next: nil}
	head.Next.Next.Next.Next = &ListNode{Val: 5, Next: nil}

	fmt.Println("Original Linked List")
	printLinkedList(head)

	head = reverseList(head)

	fmt.Println("Reversed Linked List")
	printLinkedList(head)
}
