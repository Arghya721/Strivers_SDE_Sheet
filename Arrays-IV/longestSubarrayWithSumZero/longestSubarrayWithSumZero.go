/* Given an array having both positive and negative integers. The task is to compute the length of the largest subarray with sum 0.

Example 1:

Input:
N = 8
A[] = {15,-2,2,-8,1,7,10,23}
Output: 5
Explanation: The largest subarray with
sum 0 will be -2 2 -8 1 7.
Your Task:
You just have to complete the function maxLen() which takes two arguments an array A and n, where n is the size of the array A and returns the length of the largest subarray with 0 sum.

Expected Time Complexity: O(N).
Expected Auxiliary Space: O(N).

*/

package main

import (
	"fmt"
)

func LongestSubsetWithZeroSum(arr []int) int {
	umap := make(map[int]int)

	sum := 0
	maxLen := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum == 0 {
			maxLen = i + 1
		} else {
			if val, ok := umap[sum]; ok {
				maxLen = max(maxLen, i-val)
			} else {
				umap[sum] = i
			}
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{4, 2, -3, 1, 6}
	result := LongestSubsetWithZeroSum(arr)
	fmt.Println(result)
}
