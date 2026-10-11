func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		// target - diff
		// 7 - 4 = 3
		diff := target - nums[i]
		
		// so IF we came across 3 in the array
		// we know we have a pair
		if j, ok := seen[diff]; ok {
			// J would be the smallest as this linear
			return []int{j, i}
		}

		// else we haven't seen our pair yet. 
		// so we can set seen[diff] to index i. which is what we will be searching for.
		seen[nums[i]] = i
	}

	return []int{}
}
