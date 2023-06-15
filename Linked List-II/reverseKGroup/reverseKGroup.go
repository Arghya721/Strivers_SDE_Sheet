/* Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

Example 1:

Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]

Example 2:
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
 

Constraints:

The number of nodes in the list is n.
1 <= k <= n <= 5000
0 <= Node.val <= 1000

*/ 

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func countNode(head *ListNode) int {
    count := 0

    for head !=nil {
        count++
        head=head.Next
    }
    return count
}


func reverseKGroup(head *ListNode, k int) *ListNode {
    count := countNode(head)

    dummyNode := &ListNode{0, head}


    var curr, prev, next *ListNode
    prev = dummyNode

    for count >= k {
        curr = prev.Next
        next = curr.Next 

        for i := 1; i<k ; i++ {
            curr.Next = next.Next
            next.Next = prev.Next
            prev.Next = next
            next = curr.Next
        }

        prev = curr
        count -= k
    }

    return dummyNode.Next
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	var head *ListNode = &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}

	printLinkedList(head)

	head = reverseKGroup(head, 2)

	printLinkedList(head)

}