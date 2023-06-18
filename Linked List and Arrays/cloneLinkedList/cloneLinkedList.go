/* A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

val: an integer representing Node.val
random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
Your code will only be given the head of the original linked list.

Example 1:

Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

Example 2:

Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]

Example 3:

Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]


Constraints:

0 <= n <= 1000
-10^4 <= Node.val <= 10^4
Node.random is null or is pointing to some node in the linked list.

*/

package main

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Step 1: Create a copy of each node and insert it between the original nodes
	temp := head
	for temp != nil {
		newNode := &Node{Val: temp.Val, Next: nil}
		newNode.Next = temp.Next
		temp.Next = newNode
		temp = temp.Next.Next
	}

	// Step 2: Update the random pointers of the copied nodes
	itr := head
	for itr != nil {
		if itr.Random != nil {
			itr.Next.Random = itr.Random.Next
		}
		itr = itr.Next.Next
	}

	// Step 3: Separate the original and copied nodes into two separate lists
	dummy := &Node{}
	itr = head
	temp = dummy
	var fast *Node
	for itr != nil {
		fast = itr.Next.Next
		temp.Next = itr.Next
		itr.Next = fast
		temp = temp.Next
		itr = fast
	}

	return dummy.Next
}

func printList(head *Node) {
	if head == nil {
		fmt.Println("Empty list")
		return
	}

	current := head
	for current != nil {
		fmt.Printf("(%d: %v, %v) ", current.Val, current.Next, current.Random)
		current = current.Next
	}
	fmt.Println()
}


func main() {

	var head *Node
	node1 := &Node{Val: 1, Next: nil}
	node2 := &Node{Val: 2, Next: nil}
	node3 := &Node{Val: 3, Next: nil}
	node4 := &Node{Val: 4, Next: nil}

	head = node1
	head.Next = node2
	head.Next.Next = node3
	head.Next.Next.Next = node4

	head.Random = node4
	head.Next.Random = node1
	head.Next.Next.Random = nil
	head.Next.Next.Next.Random = node2

	fmt.Println("Original Linked List")
	printList(head)

	fmt.Println("Cloned Linked List")
	clonedList := copyRandomList(head)
	printList(clonedList)
}
