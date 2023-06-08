/* Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.

 

Example 1:

Input: nums = [1,3,4,2,2]
Output: 2
Example 2:

Input: nums = [3,1,3,4,2]
Output: 3
 

Constraints:

1 <= n <= 105
nums.length == n + 1
1 <= nums[i] <= n
All the integers in nums appear only once except for precisely one integer which appears two or more times.
*/

package main 

import (
	"fmt"
)

func findDuplicate(nums []int) int {
    slow := nums[0]
    fast := nums[0]

    for {
        slow = nums[slow]
        fast = nums[nums[fast]]

        if slow == fast {
			break
		}
    }

    fast = nums[0]
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }

    return slow
}

func main() {
	nums := []int{1,3,4,2,2}
	fmt.Println(findDuplicate(nums))
}