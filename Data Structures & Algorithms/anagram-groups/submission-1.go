
func groupAnagrams(strs []string) [][]string {
	score := func(s string) [26]byte {
		var table [26]byte
		for _, c := range s {

			table[c-'a']++
		}

		return table
	}

	// score => index
	m := make(map[[26]byte]int)
	res := make([][]string, 0)

	for _, v := range strs {
		s := score(v)
		idx, ok := m[s]

		if !ok {
			m[s] = len(res)
			res = append(res, []string{v})
		} else {
			res[idx] = append(res[idx], v)
		}
	}

	return res
}