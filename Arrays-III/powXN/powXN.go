/* Implement pow(x, n), which calculates x raised to the power n (i.e., x^n).

Example 1:

Input: x = 2.00000, n = 10
Output: 1024.00000
Example 2:

Input: x = 2.10000, n = 3
Output: 9.26100
Example 3:

Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
 

Constraints:

-100.0 < x < 100.0
-2^31 <= n <= 2^31-1
n is an integer.
Either x is not zero or n > 0.
-10^4 <= x^n <= 10^4

*/

package main

import "fmt"

func myPow(x float64, n int) float64 {
    var ans float64 
    ans = 1.0
    nn := n

    if n < 0 {
        nn = -1 * nn
    }

    for nn > 0 {
        if nn%2 == 0 {
            x = x * x
            nn = nn/2
        } else {
            ans = ans * x
            nn = nn-1
        }
    }

    if n < 0 {
        ans = 1/ans
    }

    return ans
}

func main() {
	x := 2.00001
	n := 10
	fmt.Println(myPow(x, n))
}