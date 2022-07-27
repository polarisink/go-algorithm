package main

func main() {
	removeElement([]int{1, 2, 3, 3, 5, 6}, 3)
}

/**
 * @author: aries
 * @date:
 */
//leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	fast, slow, length := 0, 0, len(nums)
	for fast < length {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow = slow + 1
		}
		fast = fast + 1
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)
