/* You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	ansNode := ListNode{
		Val:  0,
		Next: nil,
	}
	temp := &ansNode
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			sum := l1.Val + l2.Val + carry
			if sum >= 10 {
				sum = sum - 10
				carry = 1
			} else {
				carry = 0
			}
			temp.Next = &ListNode{
				Val:  sum,
				Next: nil,
			}
			temp = temp.Next
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil && l2 != nil {
			sum := l2.Val + carry
			if sum >= 10 {
				sum = sum - 10
				carry = 1
			} else {
				carry = 0
			}
			temp.Next = &ListNode{
				Val:  sum,
				Next: nil,
			}
			temp = temp.Next
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			sum := l1.Val + carry
			if sum >= 10 {
				sum = sum - 10
				carry = 1
			} else {
				carry = 0
			}
			temp.Next = &ListNode{
				Val:  sum,
				Next: nil,
			}
			temp = temp.Next
			l1 = l1.Next
		}
	}

	if carry == 1 {
		temp.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return ansNode.Next
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	var l1 *ListNode = &ListNode{
		Val:  2,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	l1.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}

	var l2 *ListNode = &ListNode{
		Val:  5,
		Next: nil,
	}

	l2.Next = &ListNode{
		Val:  6,
		Next: nil,
	}

	l2.Next.Next = &ListNode{
		Val:  9,
		Next: nil,
	}

	printLinkedList(l1)
	printLinkedList(l2)

	ans := addTwoNumbers(l1, l2)

	printLinkedList(ans)

}
