package _1_sort

import "testing"

/***
 * 选择排序 select sort
 */
func Test_Test_Ssort(t *testing.T) {
	arrs := []int{3, 1, 6, 4, 5, 0, 8, 9, 7, 2}
	t.Log("选择排序---前：", arrs)

	for i := 0; i < len(arrs)-1; i++ {
		min := arrs[i]
		for j := i + 1; j < len(arrs); j++ {
			if min > arrs[j] {
				min = arrs[j]
				arrs[i], arrs[j] = arrs[j], arrs[i]
			}
		}
	}

	t.Log("选择排序---后：", arrs)
}

/**
 * 递归函数体
 */

func SelectSort(arrs []int) []int {
	for i := 0; i < len(arrs)-1; i++ {
		min := arrs[i]
		if min > arrs[i+1] {
			min = arrs[i+1]
			arrs[i], arrs[i+1] = arrs[i+1], arrs[i]
			SelectSort(arrs)
		}
	}
	return arrs
}

func Test_Ssort2(t *testing.T) {
	arrs := []int{3, 1, 6, 4, 5, 0, 8, 9, 7, 2}
	t.Log("选择排序(递归)---前：", arrs)

	SelectSort(arrs)

	t.Log("选择排序(递归）---后：", arrs)
}
