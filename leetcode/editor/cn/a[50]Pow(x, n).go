package main

const IntMax = int(^uint(0) >> 1)

/*
  todo 确实有点难
  @author: aries
  @date: 2022/07/27
*/
//leetcode submit region begin(Prohibit modification and deletion)

func myPow(a float64, k int) float64 {
	if k == 0 {
		return 1
	}
	if k == IntMax {
		return myPow(1/a, -(k+1)) / a
	}
	if k < 0 {
		return myPow(1/a, -k)
	}
	if k%2 == 1 {
		return a * myPow(a, k-1)
	} else {
		sub := myPow(a, k/2)
		return sub * sub
	}
}

//leetcode submit region end(Prohibit modification and deletion)
