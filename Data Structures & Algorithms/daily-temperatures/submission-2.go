func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i := 0; i < len(temperatures); i++ {

		// what this is saying is
		// len(stack)
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			results[top] = i - top
			stack = stack[:len(stack)-1] // pop()
		}

		stack = append(stack, i) // push()
	}

	return results
}