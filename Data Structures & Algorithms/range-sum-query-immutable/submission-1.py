class NumArray:

    def __init__(self, nums: List[int]):
        self.prefix_sums = []
        currSum = 0

        for n in nums:
            currSum += n
            self.prefix_sums.append(currSum) 


    def sumRange(self, left: int, right: int) -> int:
        return self.prefix_sums[right]  - (self.prefix_sums[left - 1] if left > 0 else 0)


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)