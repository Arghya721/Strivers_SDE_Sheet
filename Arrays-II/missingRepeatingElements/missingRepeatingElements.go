/* You are given a read only array of n integers from 1 to n.

Each integer appears exactly once except A which appears twice and B which is missing.

Return A and B.

Note: Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Note that in your output A should precede B.

Example:

Input:[3 1 2 5 3]

Output:[3, 4]

A = 3, B = 4
*/

package main

import (
	"fmt"
)

func repeatedNumber(A []int) []int {
	size := len(A)
	sr := int64(size * (size + 1) / 2)
	pr := int64(size * (size + 1) * (2*size + 1) / 6)
	sa, pa := int64(0), int64(0)

	for i := 0; i < size; i++ {
		sa += int64(A[i])
		pa += int64(A[i]) * int64(A[i])
	}

	mis := ((sr - sa) + (pr-pa)/(sr-sa)) / 2
	dup := mis - (sr - sa)

	ans := []int{int(dup), int(mis)}
	return ans
}

func main() {
	A := []int{3, 1, 2, 5, 3}
	fmt.Println(repeatedNumber(A))
}
