class Solution(object):
    def searchInsert(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        index = 0
        n = len(nums)
        while index < n:
            mid = index + (n - index) / 2
            if nums[mid] > target:
                n = mid
            elif nums[mid] < target:
                index = mid + 1
            else:
                return mid
        return index




if __name__ == '__main__':
    s = Solution()
    nums = [1,3,5,6]
    target = 2
    print s.searchInsert(nums, target)