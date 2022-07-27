package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(int(math.Pow(2, 31)+1), 1))
}

/**
 * @author: aries
 * @date:
 */
//leetcode submit region begin(Prohibit modification and deletion)
func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		return -dividend
	}
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		if dividend < 0 {
			return -divide(-dividend, divisor)
		} else {
			return -divide(dividend, -divisor)
		}
	}
	i, val, max := 0, 0, int(math.Pow(2, 31))
	for {
		val, i = val+divisor, i+1
		if val <= dividend && val+divisor > dividend {
			if i <= -max || i >= max-1 {
				return max - 1
			}
			return i
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
