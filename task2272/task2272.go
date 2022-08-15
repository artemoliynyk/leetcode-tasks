package task2272

import (
	"errors"
)

func largestVariance(s string) int {
	stringLen := len(s)

	if stringLen < 2 {
		return 0
	}

	maxVar := 0;

	for ln := 2; ln <= stringLen; ln++ {
		for offset := 0; offset < stringLen; offset++ {
			part := offset + ln

			if part > stringLen {
				break
			}

			substr := s[offset : offset+ln]

			minC, maxC, err := CalculateLeghts(substr)
			// fmt.Printf("%d/%d\n", minC, maxC)

			if err ==  nil && maxVar < maxC - minC {
				maxVar = maxC - minC
			}
		}
	}

	return maxVar
}

func CalculateLeghts(s string) (min int, max int, e error) {

	min = -1
	max = -1

	if len(s) < 2 {
		return 0, 0, errors.New("wrong string length")
	}

	// fmt.Println(s)

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
		return 1, 1, errors.New("all chars are the same")
	}

	// fmt.Println(charLens)

	for _, val := range charLens {
		if min < 0 || val < min {
			min = val
		}

		if max < 0 || val > max {
			max = val
		}
	}

	return min, max, nil
}

/*
func CalculateLeghtsSeq(s string) (min int, max int, e error) {

	min = 1
	max = 1

	if len(s) < 2 {
		return min, max, errors.New("wrong string length")
	}

	fmt.Println(s)

	charLens := make(map[string]int);
	var prev string = s[0:1]
	charLens[prev] = 1

	for i := 1; i < len(s); i++ {
		curr := s[i : i+1]

		length, exist := charLens[curr]

		if !exist {
			length = 1
		}

		if prev == curr {
			length++
		}

		charLens[curr] = length

		fmt.Printf("\tPREV: %s, CURR: %s\n", prev, curr)

		prev = curr
	}

	fmt.Println(charLens)

	if len(charLens) < 2 {
		return min, max, errors.New("all chars are the same")
	}

	for _, val := range charLens {
		fmt.Println(val)

		if val < min {
			min = val
		}

		if val > max {
			min = val
		}
	}

	return min, max, nil
}
*/
