/* Given a binary array nums, return the maximum number of consecutive 1's in the array.

 

Example 1:

Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 2
 

Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
*/

package main

import "fmt"

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxConsecutiveOnes(nums []int) int {
    count := 0
    currCount := 0
    for i := 0;i<len(nums);i++ {
        if nums[i] == 1 {
            currCount++
        }
        if nums[i] == 0 {
            currCount = 0
        }
        count = max(count,currCount)
    }
    return count
}

func main(){
	nums := []int{1,1,0,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}