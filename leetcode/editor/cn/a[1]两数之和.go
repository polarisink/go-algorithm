/**
 * @author: lqs
 * @date: 2022/1/6
 */
package main

func main() {
	arr := twoSum([]int{1, 3, 7, 11}, 18)
	for i := range arr {
		println(arr[i])
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	count := len(nums)
	var arr []int
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if nums[i]+nums[j] == target {
				arr = []int{i, j}
				break
			}
		}
	}

	return arr
}

//leetcode submit region end(Prohibit modification and deletion)
