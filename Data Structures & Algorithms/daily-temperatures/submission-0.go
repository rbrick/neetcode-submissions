func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	stack := make([][]int, 0)

	for i := 0; i < len(temperatures); i++ {
		currTemp := temperatures[i]
		if len(stack) == 0 {
			stack = append(stack, []int{i, currTemp})
		} else {
			top := stack[len(stack)-1]
			prevIdx, prevTemp := top[0], top[1]

			for currTemp > prevTemp {
				results[prevIdx] = i - prevIdx

				stack = stack[:len(stack)-1] // pop()
				if len(stack) == 0 {
					break
				}

				top := stack[len(stack)-1]
				prevIdx, prevTemp = top[0], top[1]
			}
			
			stack = append(stack, []int{i, currTemp}) // push()

		}
	}

	return results
}