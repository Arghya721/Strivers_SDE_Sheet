/* Given a string input of length n, find the length of the longest substring without repeating characters i.e return a substring that does not have any repeating characters.
Substring is the continuous sub-part of the string formed by removing zero or more characters from both ends.

Constraints:
 1<= n <=10^5

Time Limit: 1 sec
Sample Input 1:
 abcabcbb
Sample Output1:
 3
Explanation For Sample Input 1:
Substring "abc" has no repeating character with the length of 3.

Sample Input 2:
 aaaa
Sample Output 2:
1

*/

package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func uniqueSubstrings(input string) int {
	umap := make(map[rune]int)

	i := 0
	j := 0
	count := 0

	for j < len(input) {
		umap[rune(input[j])]++

		for umap[rune(input[j])] > 1 {
			umap[rune(input[i])]--
			i++
		}

		count = max(count, j-i+1)
		j++
	}

	return count
}

func main() {
	input := "abcabcbb"
	fmt.Println(uniqueSubstrings(input))
}
