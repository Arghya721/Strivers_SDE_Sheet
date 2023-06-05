/* 73. Set Matrix Zeroes
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
You must do it in place.

Example 1:

Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:

Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]


Constraints:

m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-2^31 <= matrix[i][j] <= 2^31 - 1
*/

package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	isRow := false
	isCol := false

	// checking the first row for zero
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			isRow = true
			break
		}
	}

	// checking the first column for zero
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			isCol = true
			break
		}
	}

	// mapping the first row and column according to the matrix
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// making column 0 according to the first row
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	// making row 0 according to first column
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}

	// make the first column 0 if isRow = true
	if isRow {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	// make the first row 0 if isCol = true
	if isCol {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	// print the matrix before setting zeroes
	fmt.Println("Matrix before setting zeroes: ")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	// set zeroes
	setZeroes(matrix)

	// print the matrix after setting zeroes
	fmt.Println("Matrix after setting zeroes: ")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
