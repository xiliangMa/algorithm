#coding=utf-8

class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()
        if len(nums) < 3:
            return []
        data = []
        flag = False
        for a in nums:
            temp1 = nums[:]
            temp1.remove(a)
            for b in temp1:
                temp2 = nums[:]
                temp2.remove(a)
                temp2.remove(b)
                if 0 - a - b in temp2:
                    if flag == False:
                        result = []
                        result.append(a)
                        result.append(b)
                        result.append(0 - a - b)
                        result.sort()
                        if result not in data:
                            data.append(result)
                    if a == b == 0:
                        flag = True
        return data

    def threeSum1(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()
        data = []
        length = len(nums)
        if length < 3:
            return []
        for i in range(length):
            left = i + 1
            right = length - 1
            if i != 0 and nums[i] == nums[i - 1]:
                continue

            while left < right:
                target = nums[i] + nums[left] + nums[right]
                if target == 0:
                    result = [nums[i], nums[left], nums[right]]
                    data.append(result)
                    left += 1
                    right -= 1
                    while left < right and nums[left] == nums[left - 1]:
                        left += 1
                    while left < right and nums[right] == nums[right + 1]:
                        right -= 1
                elif target < 0:
                    left += 1
                elif target > 0:
                    right -= 1
        return data





if __name__ == '__main__':
    s = Solution()
    nums = [-1, 0, 1, 2, -1, -4]
    print "三数之和", s.threeSum1(nums)