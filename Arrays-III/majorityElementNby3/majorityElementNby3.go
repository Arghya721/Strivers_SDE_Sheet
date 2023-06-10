/* Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Example 1:

Input: nums = [3,2,3]
Output: [3]
Example 2:

Input: nums = [1]
Output: [1]
Example 3:

Input: nums = [1,2]
Output: [1,2]


Constraints:

1 <= nums.length <= 5 * 10^4
-10^9 <= nums[i] <= 10^9

*/

package main

import (
	"fmt"
)

func majorityElement(nums []int) []int {
	num1 := nums[0]
	num2 := nums[0]

	count1 := 0
	count2 := 0

	for i := 0; i < len(nums); i++ {
		if num1 == nums[i] {
			count1++
		} else if num2 == nums[i] {
			count2++
		} else if count1 == 0 {
			num1 = nums[i]
			count1++
		} else if count2 == 0 {
			num2 = nums[i]
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1 = 0
	count2 = 0

	for i := 0; i < len(nums); i++ {
		if num1 == nums[i] {
			count1++
		} else if num2 == nums[i] {
			count2++
		}
	}

	var ans []int

	if count1 > len(nums)/3 {
		ans = append(ans, num1)
	}
	if count2 > len(nums)/3 {
		ans = append(ans, num2)
	}

	return ans
}

func main() {
	nums := []int{3, 2, 2, 1, 5, 2, 3}
	fmt.Println(majorityElement(nums))
}
