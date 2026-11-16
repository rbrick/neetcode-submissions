
type Counter struct {
	charCount map[rune]int
}

func (c *Counter) Enough(s string) bool {

	wordCount := Count(s)

	for _, r := range s {
		if c.charCount[r] < wordCount.charCount[r] {
			return false
		}
	}

	return true
}


func Count(s string) (counter *Counter) {
	counter = &Counter{
		charCount: make(map[rune]int),
	}

	for _, r := range s {
		counter.charCount[r]++
	}

	return
}




func countCharacters(words []string, chars string) int {
    charCounter := Count(chars)
	allLengths := 0

	for _, word := range words {
		if charCounter.Enough(word) {
			allLengths += len(word)
		}
	}


	return allLengths
}