package leetcode

import "testing"

/**
 * 合并两个有序数组
 * 归并排序处理
 *	输入:
 *	nums1 = [1,2,3,0,0,0], m = 3
 *	nums2 = [2,5,6],       n = 3
 *	输出: [1,2,2,3,5,6]
 */
func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	i, j := 0, 0
	left := nums1[:m]
	right := nums2
	var result []int
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i += 1
		} else {
			result = append(result, right[j])
			j += 1
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}


func Test_Merge(t *testing.T){
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	t.Log(Merge(nums1, m, nums2, n))
}