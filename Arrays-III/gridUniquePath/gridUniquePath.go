/* There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.

Example 1:

Input: m = 3, n = 7
Output: 28
Example 2:

Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
 

Constraints:
1 <= m, n <= 100

*/

package main 

import (
	"fmt"
)

// func countPath(i int, j int, m int,n int, dp [][]int) int {
//     if i == m-1 || j == n-1 {
//         return 1
//     } 
//     if i >= m || j >= n {
//         return 0
//     }

//     if dp[i][j] != -1 {
//         return dp[i][j]
//     }

//     return countPath(i+1,j,m,n, dp) + countPath(i,j+1,m,n,dp)
// }

// func uniquePaths(m int, n int) int {

//     dp := make([][]int,m+1)
//     for i := range dp {
//         dp[i] = make([]int,n+1)
//     }

//     for i := 0; i < m+1; i++ {
//         for j := 0; j < n+1; j++ {
//             dp[i][j] = -1
//         }
//     }

//     return countPath(0,0,m,n, dp);
// }

func uniquePaths(m int, n int) int {
	N := m + n - 2
	r := m - 1
	res := 1.0

	for i := 1; i <= r; i++ {
		res = res * float64(N-r+i) / float64(i)
	}

	return int(res)
}


func main() {
    fmt.Println(uniquePaths(51,9))
}