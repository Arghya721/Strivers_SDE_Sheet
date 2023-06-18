/* Given the head of a linked list, rotate the list to the right by k places.

Example 1:

Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]

Example 2:

Input: head = [0,1,2], k = 4
Output: [2,0,1]
 

Constraints:

The number of nodes in the list is in the range [0, 500].
-100 <= Node.val <= 100
0 <= k <= 2 * 10^9

*/ 

package main 

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

    if head == nil || k == 0  || head.Next == nil {
        return head
    }

    temp := head
    count := 1

    // calculating length
    for temp.Next != nil {
        count++
        temp = temp.Next
    }

    // link the last node to first
    temp.Next = head
    k = k%count
    k = count - k

    for k != 0 {
        temp = temp.Next
        k--
    }

    head = temp.Next

    temp.Next = nil

    return head
    
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
		Val:  3,
		Next: nil,
	}
	var l4 *ListNode = &ListNode{
		Val:  4,
		Next: nil,
	}
	var l5 *ListNode = &ListNode{
		Val:  5,
		Next: nil,
	}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	k := 2

	head := rotateRight(l1, k)

	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}

}

