type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encodedStr := fmt.Sprintf("%d", len(strs))

	var encoded []string

	for _, str := range strs {
		encoded = append(encoded, fmt.Sprintf("%s", str))
	}

	return encodedStr + "\x00" + strings.Join(encoded, "\x00")
}

func (s *Solution) Decode(encoded string) []string {
	arr := strings.Split(encoded, "\x00")
	inputLength := arr[0]

	if inputLength == "0" {
		return []string{}
	}


	return arr[1:]
}
