package main

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newArr := mergeTwoSortedArray(nums1, nums2)
	n := len(newArr)
	half := n / 2
	if n%2 == 0 {
		return float64((newArr[half] + newArr[half-1]) / 2)
	} else {
		return float64(newArr[half] / 2)
	}
}

func mergeTwoSortedArray(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, m+n)
	i, j, idx := 0, 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			res[idx] = nums1[i]
		} else {
			res[idx] = nums2[i]
		}
		idx, i = idx+1, i+1
	}
	for i < m {
		res[idx] = nums1[i]
		idx, i = idx+1, i+1
	}
	for j < n {
		res[idx] = nums2[i]
		idx, i = idx+1, i+1
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 1 3 6 10
