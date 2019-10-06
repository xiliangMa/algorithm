package mian

import (
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m -= 1
		} else {
			nums1[m+n-1] = nums2[n-1]
			n -= 1
		}
	}
	if m == 0 {
		for i := 0; i <= n-1; i++ {
			nums1[i] = nums2[i]
		}
	}
	return nums1
}

func Test_Merger(t *testing.T) {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	t.Log(merge(nums1, m, nums2, n))


	for i := 0; i <= n-1; i++ {
		nums1[i] = nums2[i]
	}



}
