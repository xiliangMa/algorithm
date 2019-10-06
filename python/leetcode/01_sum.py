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

if __name__ == '__main__':
    nums = [3, 2, 3]
    target = 6
    s = Solution()
    print "两数之和", s.twoSum(nums, target)

    nums = [12,0,3,-14,5,-11,11,-5,-2,-1,6,-7,-10,1,4,1,1,9,-3,6,-15,0,6,1,6,-12,3,7,11,-6,-8,0,9,3,-7,-1,7,-10,1,13,-4,-7,-9,-7,9,3,1,-13,-3,13,8,-11,-9,-8,-3,4,-13,7,-11,5,-14,-4,-9,10,6,-9,-6,-9,-12,11,-11,-9,11,-5,0,-3,13,-14,-1,-13,7,-7,14,5,0,-4,-6,-6,-11,-2,14,-10,2,12,8,-7,-11,-13,-9,14,5,-5,-9,1,-2,6,5,-12,-1,-10,-9,-9,-10,12,11]
    print "三数之和", s.threeSum(nums)

