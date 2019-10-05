package leetcode

import "testing"

/**
 * 求众数
 */

func Majority_Element(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	num_map := make(map[int]int)
	for _, value := range nums {
		if _, ok := num_map[value]; ok {
			num_map[value] += 1
			if num_map[value] > len(nums) / 2 {
				return value
			}
		} else {
			num_map[value] = 1
		}
	}

	return 0
}

func Test_Majority_Element(t *testing.T){
	nums := []int{3, 2, 3}
	t.Log(Majority_Element(nums))
}