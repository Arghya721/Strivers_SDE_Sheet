/* Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
Only one valid answer exists.

*/

package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	var ans []int
	for i := 0; i < len(nums); i++ {
		if index, ok := mp[target-nums[i]]; ok {
			ans = append(ans, index)
			ans = append(ans, i)
			return ans
		}
		mp[nums[i]] = i
	}
	return ans
}

func pairSum(arr []int, s int) [][]int {
	mp := make(map[int]int)
	var ans [][]int
	for i := 0; i < len(arr); i++ {
		if val, ok := mp[s-arr[i]]; ok {
			output := make([]int, 0)
			for val > 0 {
				output = append(output, min(arr[i], s-arr[i]))
				output = append(output, max(arr[i], s-arr[i]))
				ans = append(ans, output)
				val--
			}
		}
		mp[arr[i]]++
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i][0] != ans[j][0] {
			return ans[i][0] < ans[j][0]
		}
		return ans[i][1] < ans[j][1]
	})
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	fmt.Println(twoSum(nums, target))
	fmt.Println(pairSum(nums, target))
}
