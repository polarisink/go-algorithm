package main

func main() {
	reverse(12345)
}

//leetcode submit region begin(Prohibit modification and deletion)

const num = 214748364

func reverse(x int) int {
	res := 0
	for x != 0 {
		temp := x % 10
		if res > num || (res == num && temp > 7) {
			return 0
		}
		if res < -num || (res == -num && temp < -8) {
			return 0
		}
		res, x = res*10+temp, x/10
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
