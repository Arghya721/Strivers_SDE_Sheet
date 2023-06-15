/* Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.

Example 1:

Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'
Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
- Note that the intersected node's value is not 1 because the nodes with value 1 in A and B (2nd node in A and 3rd node in B) are different node references. In other words, they point to two different locations in memory, while the nodes with value 8 in A and B (3rd node in A and 4th node in B) point to the same location in memory.

Example 2:

Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Intersected at '2'
Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [1,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.

Example 3:

Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: No intersection
Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.


Constraints:

The number of nodes of listA is in the m.
The number of nodes of listB is in the n.
1 <= m, n <= 3 * 10^4
1 <= Node.val <= 10^5
0 <= skipA < m
0 <= skipB < n
intersectVal is 0 if listA and listB do not intersect.
intersectVal == listA[skipA] == listB[skipB] if listA and listB intersect.

*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ptr1 := headA
	ptr2 := headB

	if headA == nil || headB == nil {
		return nil
	}

	for ptr1 != ptr2 {

		if ptr1 == nil {
			ptr1 = headB
		} else {
			ptr1 = ptr1.Next
		}

		if ptr2 == nil {
			ptr2 = headA
		} else {
			ptr2 = ptr2.Next
		}
	}

	return ptr1

}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}


func main() {
	// Create the intersected nodes
	intersectionNode := &ListNode{
		Val:  8,
		Next: nil,
	}

	// Create the first linked list
	var l1 ListNode
	l1.Val = 4
	l1.Next = &ListNode{
		Val:  1,
		Next: intersectionNode,
	}
	l1.Next.Next.Next = &ListNode{
		Val:  7,
		Next: nil,
	}

	// Create the second linked list
	var l2 ListNode
	l2.Val = 5
	l2.Next = &ListNode{
		Val:  6,
		Next: intersectionNode,
	}

	// printLinkedList(&l1)

	fmt.Println(getIntersectionNode(&l1, &l2).Val)
}
