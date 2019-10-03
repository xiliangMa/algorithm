package _1_sort

import (
	"testing"
)

/**
 * 快速排序 quick sort
 */
func Qsort(arrs []int) []int{
	if len(arrs) >= 2 {
		pivot := arrs[len(arrs)-1]
		var left, right []int
		arrs = arrs[:len(arrs) -1]
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

func Test_Qsort(t *testing.T){
	arrs := []int{5, 3, 1, 2, 4}
	t.Log("快速排序---后：", Qsort(arrs))
}
