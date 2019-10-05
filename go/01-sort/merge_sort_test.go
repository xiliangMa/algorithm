package _1_sort

import (
	"testing"
)

/**
 * 归并排序 merger sort 递归
 */
func Merger(left, right []int) []int {
	i, j := 0, 0
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

func Merger_Sort(arrs []int) []int {
	if len(arrs) <= 1 {
		return arrs
	}

	middle := len(arrs) / 2
	left := Merger_Sort(arrs[:middle])
	right := Merger_Sort(arrs[middle:])
	return Merger(left, right)

}

func Test_Msort(t *testing.T) {
	arrs := []int{3, 1, 6, 4, 5, 0, 8, 9, 7, 2}
	t.Log("归并排序(递归)---后：", Merger_Sort(arrs))
}
