/* Kadane's Algorithm (Maximum Subarray Problem)
Given an integer array nums, find the
subarray
 with the largest sum, and return its sum.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.


Constraints:

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4

*/

package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	sum := 0
	ans := math.MinInt

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if sum < 0 {
			sum = 0
		}

		if sum > ans {
			ans = sum
		}
	}

	return ans
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
