
func evalRPN(tokens []string) int {
	currStack := []int{}
	asNumeric := func(s string) (int, bool) {
		i, err := strconv.Atoi(s)

		if err != nil {
			return 0, false
		}

		return i, true
	}

	pop := func() int {
		// lifo
		elem := currStack[len(currStack)-1]
		currStack = currStack[:len(currStack)-1]
		return elem
	}

	push := func(i int) {
		currStack = append(currStack, i)
	}

	// '+', '-', '*', and '/'

	operations := map[string]func(a, b int) int{
		"+": func(a, b int) int {
			return a + b
		},
		"-": func(a, b int) int {
			return a - b
		},
		"/": func(a, b int) int {
			return a / b
		},
		"*": func(a, b int) int {
			return a * b
		},
	}

	for _, token := range tokens {

		if number, isNumber := asNumeric(token); isNumber {
			push(number)
		} else {
			// ["1","2","+","3","*","4","-"]

			// "4","13","5","/","+"
			// stack: [4, 13, 5]
			// 13/5 = 2
			// stack [4, 2]
			// add
			// stack [6]
			op := operations[token]
			// get top two elements (binops)
			a, b := pop(), pop()
			res := op(b, a)

			push(res)
		}

	}

	return pop()
}
