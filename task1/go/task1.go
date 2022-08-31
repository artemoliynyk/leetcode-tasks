package task1

func twoSum(nums []int, target int) []int {
	for i1, n1 := range nums {
		for i2, n2 := range nums[i1+1:] {
			
			if n1+n2 == target {
				return []int{i1, i2+i1+1}
			}
		}
	}

	return []int{-1, -1}
}
