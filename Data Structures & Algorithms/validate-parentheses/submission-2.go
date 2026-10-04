func isValid(s string) bool {
	stack := list.New()

	isClosing := func(r rune) bool {
		return strings.ContainsRune("])}", r)
	}

	closingFor := func(r rune) rune {
		switch r {
		case '(':
			return ')'
		case '[':
			return ']'
		case '{':
			return '}'

		default:
			return 0x0
		}
	}

	if isClosing(rune(s[0])) {
		return false
	}

	// ( | [ | {
	stack.PushBack(rune(s[0]))

	for i := 1; i < len(s); i++ {
		curr := rune(s[i])

		if isClosing(curr) {

			prior := stack.Back()

			if prior == nil {
				return false
			}

			if curr != closingFor(prior.Value.(rune)) {
				return false
			}

			stack.Remove(prior)
			continue
		} else {
			stack.PushBack(curr)
		}
	}

	return stack.Len() == 0
}