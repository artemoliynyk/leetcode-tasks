from logging import debug
import sys
from typing import List, NamedTuple
import unittest
import os
from task1 import Solution


class TestValue(NamedTuple):
    nums: List[int]
    target: int
    expected: List[int]


testData = []
# List[int]{2, 7, 11, 15}, 9, List[int]{0, 1}

testData.append(TestValue([2, 7, 11, 15], 9, [0, 1]))
testData.append(TestValue([3, 2, 4], 6, [1, 2]))
testData.append(TestValue([3, 3], 6, [0, 1]))


class SolutionTest(unittest.TestCase):
    def test_twoSum(self):
        for testSet in testData:
            actual = Solution.twoSum(self, testSet.nums, testSet.target)
            print("actual" + actual)


if __name__ == '__main__':
    unittest.main()
