package leetcode

import (
	"testing"
)

/**
 * 只出现一次的数字
 */
func SingleNumber(nums []int) int {
	num_map := make(map[int]int)
	for _, value := range nums {
		if _, ok := num_map[value]; ok {
			num_map[value] += 1
		} else {
			num_map[value] = 1
		}
	}

	for i, value := range num_map {
		if value == 1 {
			return i
		}
	}
	return 0
}

func Test_Single_Number(t *testing.T) {
	nums := []int{2,2,1}
	t.Log(SingleNumber(nums))
}
