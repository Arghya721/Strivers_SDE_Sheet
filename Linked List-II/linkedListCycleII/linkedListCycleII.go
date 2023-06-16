/* Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

 

Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:

Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.

Example 3:

Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
 

Constraints:

The number of the nodes in the list is in the range [0, 104].
-10^5 <= Node.val <= 10^5
pos is -1 or a valid index in the linked-list.

*/ 

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

    if head == nil {
        return head
    }

    fast := head
    slow := head 
    start := head

    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if fast == slow {
            for slow != start {
                slow = slow.Next
                start = start.Next
            }
            return slow
        }
    }
    return nil
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

	l1.Next = l2
	l2.Next = l1

	fmt.Println(detectCycle(l1).Val)
}