/* Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.



Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9


Constraints:

0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9

*/

package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	sorted := make(map[int]bool)
	for _, num := range nums {
		sorted[num] = true
	}

	if len(sorted) == 0 {
		return 0
	}

	ans := 1
	curr := 1
	keys := make([]int, 0, len(sorted))
	for k := range sorted {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	prev := keys[0]

	for _, num := range keys {
		if num == prev+1 {
			curr++
		} else if num != prev {
			curr = 1
		}
		prev = num
		ans = max(ans, curr)
	}
	return ans
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
