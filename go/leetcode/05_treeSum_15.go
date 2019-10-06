package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	if n < 3 {
		return nil
	}
	result := [][]int{}
	for i := 0; i < n; i++{
		l := i + 1
		r := n - 1
		if i != 0 && nums[i] == nums[i - 1] {
			continue
		}

		for l < r {
			target := nums[i] + nums[l] + nums[r]
			if target == 0 {
				data := []int{nums[i], nums[l], nums[r]}
				result = append(result, data)
				l += 1
				r -= 1
				for l < r && nums[l] == nums[l - 1]{
					l += 1
				}
				for l < r && nums[r] == nums[r + 1]{
					r -= 1
				}
			} else if target < 0 {
				l += 1
			} else {
				r -= 1
			}
		}

	}
	return result
}

func main() {
	nums := []int{0,-4,-1,-4,-2,-3,2}
	fmt.Println(threeSum(nums))

}