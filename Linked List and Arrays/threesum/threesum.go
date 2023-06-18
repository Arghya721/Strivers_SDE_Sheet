/* Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

 

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
 

Constraints:

3 <= nums.length <= 3000
-10^5 <= nums[i] <= 10^5

*/

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}

	// Sort the input array
	sort.Ints(nums)

	// Use a set to store unique triplets
	set := make(map[[3]int]bool)

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				// Found a triplet with sum 0, add it to the set
				triplet := [3]int{nums[i], nums[j], nums[k]}
				set[triplet] = true

				// Move the pointers to find the next unique pairs
				p1 := nums[j]
				p2 := nums[k]
				for p1 == nums[j] && j < k {
					j++
				}
				for p2 == nums[k] && j < k {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	// Convert the set to a 2D slice
	for triplet := range set {
		result = append(result, []int{triplet[0], triplet[1], triplet[2]})
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}