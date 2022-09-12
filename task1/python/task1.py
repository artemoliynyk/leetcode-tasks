#!python3

from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        visited = {}

        for i, v in enumerate(nums):
            if (target - v) in visited:
                return [i, visited[target - v]]
            visited[v] = i
        
        return [-1, -1]