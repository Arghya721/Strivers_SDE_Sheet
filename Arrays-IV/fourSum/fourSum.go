/* Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

 

Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
 

Constraints:

1 <= nums.length <= 200
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
*/

package main 

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
    var ans [][]int 

    sort.Ints(nums)
    
    for i := 0 ; i < len(nums) ; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j := i+1;j<len(nums);j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }

            k := j+1
            l := len(nums) - 1 

            for k < l {
                if nums[i]+nums[j]+nums[k]+nums[l] == target {
                    temp := []int{nums[i],nums[j],nums[k],nums[l]} 
                    ans = append(ans,temp)
                    k++
                    l--
                    for k < l && nums[k] == nums[k-1] {
                        k++
                    } 
                    for k < l && nums[l] == nums[l+1] {
                        l--
                    }
                } else if nums[i]+nums[j]+nums[k]+nums[l] > target {
                    l--
                } else {
                    k++
                }
            }
        }
    }
    return ans
}

func main() {
	nums := []int{1,0,-1,0,-2,2}
	target := 0
	fmt.Println(fourSum(nums,target))
}