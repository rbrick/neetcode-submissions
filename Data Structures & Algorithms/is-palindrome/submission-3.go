func isValid(a rune) bool {
	valid := (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
	return valid
}

func isPalindrome(s string) bool {

	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if isValid(r) {
			return r
		}

		return -1
	}, s)

	left, right := 0, len(s)-1

	for left < right {
		a, b := s[left], s[right]

		if a != b {
			return false
		}

		left++
		right--
	}

	return true
}
