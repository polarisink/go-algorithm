package main

/**
  二分法
  @author: aries
  @date:
*/
//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	low, high := 0, size-1
	for low <= high {
		mid := (high + low) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low

}

//紫色
//leetcode submit region end(Prohibit modification and deletion)
