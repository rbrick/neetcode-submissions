func maxArea(heights []int) int {
	left, right := 0, len(heights)-1
	area := 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for left < right {
		a, b := heights[left], heights[right]
		w, h := right-left, min(a, b)
		ar := w * h

		if ar > area {

			// log.Printf("%d * %d = %d. %d > %d\n", w, h, ar, ar, area)
			area = ar
		}

		if b >= a {
			left++
		} else {
           right--
		}	
		
	}

	return area
}
