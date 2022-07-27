package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(13531))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	ss, size := []rune(s), len(s)
	for i := 0; i <= (size-1)/2; i++ {
		if ss[i] != ss[size-1-i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
