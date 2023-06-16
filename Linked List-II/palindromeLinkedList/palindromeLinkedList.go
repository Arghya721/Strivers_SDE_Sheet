/* Given the head of a singly linked list, return true if it is a
palindrome or false otherwise.

Example 1:
Input: head = [1,2,2,1]
Output: true
Example 2:

Input: head = [1,2]
Output: false


Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	var next *ListNode
	prev = nil
	next = nil

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev

}

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = reverse(slow.Next)
	slow = slow.Next
	dummy := head

	for slow != nil {
		if slow.Val != dummy.Val {
			return false
		}
		slow = slow.Next
		dummy = dummy.Next
	}

	return true
}

func main() {
	var l1 *ListNode = &ListNode{
		Val:  1,
		Next: nil,
	}
	var l2 *ListNode = &ListNode{
		Val:  2,
		Next: nil,
	}
    var l3 *ListNode = &ListNode{
        Val:  2,
        Next: nil,
    }
    var l4 *ListNode = &ListNode{
        Val:  1,
        Next: nil,
    }

	l1.Next = l2
    l2.Next = l3
    l3.Next = l4

	fmt.Println(isPalindrome(l1))
}
