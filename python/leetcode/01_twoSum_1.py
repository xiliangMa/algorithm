#coding=utf-8

class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        if len(nums) >= 2:
            if len(nums) == 2:
                return 0, 1
            for i in range(len(nums)):
                if target-nums[i] in nums[i+1:]:
                    return [i,nums.index(target-nums[i], i+1)]

if __name__ == '__main__':
    nums = [3, 2, 3]
    target = 6
    s = Solution()
    print "两数之和", s.twoSum(nums, target)

