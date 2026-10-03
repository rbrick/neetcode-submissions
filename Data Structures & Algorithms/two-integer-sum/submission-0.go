func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

	for i, n := range nums {
		diff := target - n

		if j, ok := seen[n]; ok {
			return []int{j, i}
		}

		seen[diff] = i 
	}

	
	return []int{0, 0}
}
