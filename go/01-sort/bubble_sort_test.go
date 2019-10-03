package _1_sort

import (
	"testing"
)

/**
 * 冒泡排序 Bubble Sort
 */
func Test_BSort(t *testing.T) {
	arrs := []int{3, 1, 6, 4, 5, 0, 8, 9, 7, 2}
	length := len(arrs)

	t.Log("冒泡排序---前：", arrs)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if arrs[i] > arrs[j] {
				arrs[i], arrs[j] = arrs[j], arrs[i]
			}
		}
	}
	t.Log("冒泡排序---后：", arrs)
}

/**
 * 冒泡排序递归函数
 */
func Bsort(arrs []int) []int {
	for i := 0; i < len(arrs)-1; i++ {
		if arrs[i] > arrs[i+1] {
			temp := arrs[i]
			arrs[i] = arrs[i+1]
			arrs[i+1] = temp
			Bsort(arrs)
		}
	}
	return arrs
}

/**
 * 冒泡排序(递归) Bubble Sort
 */
func Test_Bsort2(t *testing.T) {
	arrs := []int{3, 1, 6, 4, 5, 0, 8, 9, 7, 2}
	Bsort(arrs)
	t.Log("冒泡排序(递归)---后：", arrs)
}
