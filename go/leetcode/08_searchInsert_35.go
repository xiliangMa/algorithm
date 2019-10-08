package leetcode

func searchInsert(nums []int, target int) int {
	index := 0
	n := len(nums)

	for index < n {
		middle := index + (n - index) / 2
		if target < nums[middle] {
			n = middle
		} else if target > nums[middle] {
			index = middle + 1
		} else {
			return middle
		}
	}
	return index
}
