func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)

	for _, n := range nums {
		seen[n] = true
	}

	return len(seen) < len(nums)
}
