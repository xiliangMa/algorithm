package _1_sort

import (
	"testing"
)

/**
 * 快速排序(1) quick sort
 */
func Qsort(arrs []int) []int {
	if len(arrs) >= 2 {
		pivot := arrs[len(arrs)-1]
		var left, right []int
		arrs = arrs[:len(arrs)-1]
		for _, arr := range arrs {
			if arr <= pivot {
				left = append(left, arr)
			} else {
				right = append(right, arr)
			}
		}
		return append(append(Qsort(left), pivot), Qsort(right)...)

	} else {
		return arrs
	}

}

/**
 * 快速排序(1) quick sort
 * 避免额外的内存开销 去除前面left、right 数组
 */
func Qsort2(arrs []int, start, end int) []int {
	if start >= end {
		return arrs
	}
	pivod := arrs[end]
	left := start
	right := end - 1
	for left <= right {
		for left <= right && arrs[left] < pivod {
			left += 1
		}


		for left <= right && arrs[right] > pivod {
			right -= 1
		}

		if left >= right {
			break
		}
		arrs[left], arrs[right] = arrs[right], arrs[left]
	}

	arrs[left], arrs[end] = arrs[end], arrs[left]
	Qsort2(arrs, start, left-1)
	Qsort2(arrs, left+1, end)
	return arrs
}

func Test_Qsort(t *testing.T) {
	arrs := []int{3, 1, 6, 4, 5, 0, 8, 9, 7, 2}
	t.Log("快速排序(1)---后：", Qsort(arrs))

	t.Log("快速排序(2)---后：", Qsort2(arrs, 0, 9))
}
