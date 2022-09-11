package task1

func twoSum(nums []int, target int) []int {
	// v2
	seen := make(map[int]int)

	for i, v := range nums {
		if seenIdx, exist := seen[target-v]; exist {
			return []int{i, seenIdx}
		}

		seen[v] = i
	}

	return []int{0, 0}
}
