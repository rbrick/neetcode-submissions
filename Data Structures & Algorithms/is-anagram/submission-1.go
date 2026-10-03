func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}	

	freq1, freq2 := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(s); i++ {
		freq1[s[i]]++
		freq2[t[i]]++
	}

	for _, c := range t {
		if freq1[byte(c)] != freq2[byte(c)] {
			return false
		}
	}
	return true
}
