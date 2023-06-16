/* Given a Linked List of size N, where every node represents a sub-linked-list and contains two pointers:
(i) a next pointer to the next node,
(ii) a bottom pointer to a linked list where this node is head.
Each of the sub-linked-list is in sorted order.
Flatten the Link List such that all the nodes appear in a single level while maintaining the sorted order.

Note: The flattened list will be printed using the bottom pointer instead of the next pointer.
For more clarity have a look at the printList() function in the driver code.



Example 1:

Input:
5 -> 10 -> 19 -> 28
|     |     |     |
7     20    22   35
|           |     |
8          50    40
|                 |
30               45
Output:  5-> 7-> 8- > 10 -> 19-> 20->
22-> 28-> 30-> 35-> 40-> 45-> 50.
Explanation:
The resultant linked lists has every
node in a single level.
(Note: | represents the bottom pointer.)


Example 2:

Input:
5 -> 10 -> 19 -> 28
|          |
7          22
|          |
8          50
|
30
Output: 5->7->8->10->19->22->28->30->50
Explanation:
The resultant linked lists has every
node in a single level.

(Note: | represents the bottom pointer.)


Your Task:
You do not need to read input or print anything. Complete the function flatten() that takes the head of the linked list as input parameter and returns the head of flattened link list.

Expected Time Complexity: O(N*N*M)
Expected Auxiliary Space: O(1)

*/

package main

import "fmt"

type Node struct {
	data   int
	next   *Node
	bottom *Node
}

func mergeTwoLinkedList(a, b *Node) *Node {
	dummy := &Node{data: 0}
	temp := dummy

	for a != nil && b != nil {
		if a.data < b.data {
			temp.bottom = a
			temp = temp.bottom
			a = a.bottom
		} else {
			temp.bottom = b
			temp = temp.bottom
			b = b.bottom
		}
	}

	for a != nil {
		temp.bottom = a
		temp = temp.bottom
		a = a.bottom
	}
	for b != nil {
		temp.bottom = b
		temp = temp.bottom
		b = b.bottom
	}

	return dummy.bottom
}

func flatten(root *Node) *Node {
	if root.next == nil {
		return root
	}

	root.next = flatten(root.next)

	root = mergeTwoLinkedList(root, root.next)

	return root
}

func main() {
	// Create and flatten the linked list
	// You can modify this part based on your requirements
	root := &Node{data: 1}
	root.next = &Node{data: 2}
	root.next.next = &Node{data: 3}
	root.next.next.next = &Node{data: 4}
	root.next.next.bottom = &Node{data: 5}
	root.next.next.bottom.bottom = &Node{data: 6}
	root.next.next.next.bottom = &Node{data: 7}
	root.next.next.next.bottom.bottom = &Node{data: 8}
	root.next.next.next.next = &Node{data: 9}
	root.next.next.next.next.bottom = &Node{data: 10}

	result := flatten(root)
	for result != nil {
		fmt.Print(result.data, "->")
		result = result.bottom
	}
}
