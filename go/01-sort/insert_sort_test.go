package _1_sort

import "testing"

/**
 * 插入排序 Insert sort
 */
func Test_ISort(t *testing.T) {
	arrs := []int{3, 1, 6, 4, 5, 0, 8, 9, 7, 2}
	for i := 0; i < len(arrs)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arrs[j-1] > arrs[j] {
				arrs[j-1], arrs[j] = arrs[j], arrs[j-1]
			}
		}
		t.Log("插入排	序---后：", arrs)
	}

}
