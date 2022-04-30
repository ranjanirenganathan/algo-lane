package strings

import "fmt"

/*Check for balanced parentheses in linear time using constant space  Related To: Strings
Problem
Example:

((()())(())) should be accepted but ())() should be rejected.​
Answer
Initialize a counter at zero.
Whenever you see (, increase it by one, and whenever you see ),
decrease it by one. Accept if the counter was always non negative and ended up at zero.
Since counter is always between −1 and n, only O(log n) bits (note it is not O(1)) are needed to store it.
*/

func Isbalanced(s string) string {
	right := 1
	left := 1
	var result string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == `(` {
			left++
		} else if string(s[i]) == `)` {
			right++
		} else {
			fmt.Printf("Invalid input provided")
			break
		}
		if left == right {
			result = "Accepted"
		} else {
			result = "Rejected"
		}
	}
	return result
}
