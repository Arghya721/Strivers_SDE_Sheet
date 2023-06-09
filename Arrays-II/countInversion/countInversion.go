/* Problem Statement: Given an array of N integers, count the inversion of the array (using merge-sort).

What is an inversion of an array? Definition: for all i & j < size of array, if i < j then you have to find pair (A[i],A[j]) such that A[j] < A[i].

Example 1:

Input Format: N = 5, array[] = {1,2,3,4,5}

Result: 0

Example 2:

Input Format: N = 5, array[] = {5,4,3,2,1}

Result: 10

*/

package main 

import "fmt"

func merge(arr []int , left , mid , right int) int {
	i := left
	j := mid+1
	k := 0

	temp := make([]int, right-left+1)

	count := 0

	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			k++
			i++
		} else {
			temp[k] = arr[j]
			k++
			j++
			count += mid - i + 1
		}
	}

	for i <= mid {
		temp[k] = arr[i]
		k++
		i++
	}

	for j <= right {
		temp[k] = arr[j]
		k++
		j++
	}

	for k = 0; k < len(temp); k++ {
		arr[k+left] = temp[k]
	}

	return count
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


func countInversion(arr []int) int {
	return mergeSort(arr, 0, len(arr)-1)
}


func main() {
	arr := []int{5,4,3,2,1}
	fmt.Println(countInversion(arr))
}