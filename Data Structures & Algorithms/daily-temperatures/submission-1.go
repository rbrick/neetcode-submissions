func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			currTemp := temperatures[i]
			top := stack[len(stack)-1]
			prevIdx, prevTemp := top, temperatures[top]

			for currTemp > prevTemp {
				results[prevIdx] = i - prevIdx

				stack = stack[:len(stack)-1] // pop()
				if len(stack) == 0 {
					break
				}

				top := stack[len(stack)-1]
				prevIdx, prevTemp = top, temperatures[top]
			}

			stack = append(stack, i) // push()
		}
	}

	return results
}