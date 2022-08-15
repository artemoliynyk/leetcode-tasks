package task2272

func largestVariance(s string) int {
	stringLen := len(s)

	if stringLen < 2 {
		return 0
	}

	maxVar := 0;

	for ln := 3; ln <= stringLen; ln++ {
		for offset := 0; offset < stringLen; offset++ {
			part := offset + ln

			if part > stringLen {
				break
			}

			substr := s[offset : offset+ln]

			max := CalculateLeghts(substr)

			if maxVar < max {
				maxVar = max
			}
		}
	}

	return maxVar
}

func CalculateLeghts(s string) (int) {

	if len(s) < 3 {
		return 0
	}

	charLens := make(map[string]int)

	for i := 0; i < len(s); i++ {
		curr := s[i : i+1]

		length, exist := charLens[curr]

		if !exist {
			length = 0
		}

		charLens[curr] = length + 1
	}

	if len(charLens) < 2 {
		return 0
	}

	min, max := -1, -1

	for _, val := range charLens {
		if min < 0 || val < min {
			min = val
		}

		if max < 0 || val > max {
			max = val
		}
	}

	if max > min {
		return max - min
	}

	return 0
}