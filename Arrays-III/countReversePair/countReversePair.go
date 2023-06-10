/* Given an integer array nums, return the number of reverse pairs in the array.

A reverse pair is a pair (i, j) where:

0 <= i < j < nums.length and
nums[i] > 2 * nums[j].
 

Example 1:

Input: nums = [1,3,2,3,1]
Output: 2
Explanation: The reverse pairs are:
(1, 4) --> nums[1] = 3, nums[4] = 1, 3 > 2 * 1
(3, 4) --> nums[3] = 3, nums[4] = 1, 3 > 2 * 1
Example 2:

Input: nums = [2,4,3,5,1]
Output: 3
Explanation: The reverse pairs are:
(1, 4) --> nums[1] = 4, nums[4] = 1, 4 > 2 * 1
(2, 4) --> nums[2] = 3, nums[4] = 1, 3 > 2 * 1
(3, 4) --> nums[3] = 5, nums[4] = 1, 5 > 2 * 1
 

Constraints:

1 <= nums.length <= 5 * 104
-2^31 <= nums[i] <= 2^31 - 1

*/

package main 

import (
	"fmt"
)

func merge(nums []int, low int, mid int, high int) int {
	total := 0
	j := mid + 1

	for i := low; i <= mid; i++ {
		for j <= high && int64(nums[i]) > 2*int64(nums[j]) {
			j++
		}
		total += (j - (mid + 1))
	}

	t := make([]int, 0)
	left, right := low, mid+1

	for left <= mid && right <= high {
		if nums[left] <= nums[right] {
			t = append(t, nums[left])
			left++
		} else {
			t = append(t, nums[right])
			right++
		}
	}

	for left <= mid {
		t = append(t, nums[left])
		left++
	}
	for right <= high {
		t = append(t, nums[right])
		right++
	}

	for i := low; i <= high; i++ {
		nums[i] = t[i-low]
	}

	return total
}

func mergeSort(arr []int , left , right int) int {
	count := 0;
	if left < right {
		mid := (left + right)/2
		count += mergeSort(arr, left, mid)
		count += mergeSort(arr, mid+1, right)
		count += merge(arr,left , mid , right)
	}
	return count
}

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func main() {
	fmt.Println(reversePairs([]int{1,3,2,3,1}))
	fmt.Println(reversePairs([]int{2,4,3,5,1}))
}