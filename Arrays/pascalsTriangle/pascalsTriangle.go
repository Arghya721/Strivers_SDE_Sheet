/*
	Given an integer numRows, return the first numRows of Pascal's triangle.

Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:

Input: numRows = 1
Output: [[1]]

Constraints:

1 <= numRows <= 30
*/
package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	var ans [][]int

	ans = make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)

		ans[i][0] = 1
		ans[i][i] = 1

		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans
}

func main() {
	pascalTriangle := generate(5)

	fmt.Println(pascalTriangle)
}
