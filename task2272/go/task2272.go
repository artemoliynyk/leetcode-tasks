package task2272

import (
	"fmt"
	"strings"
)

func uniqString(s string) string {
	uString := ""

	for _, c := range s {
		if !strings.ContainsRune(uString, c) {
			uString += string(c)
		}
	}

	return uString
}

func largestVariance(s string) int {
	var variance int = 0
	chars := uniqString(s)

	if len(chars) <= 1 || len(chars) == len(s) {
		return 0
	}

	for i1, c1 := range chars {
		for i2, c2 := range chars {

			if i1 == i2 {
				continue
			}

			fmt.Printf("Checking pair: %v, %v\n", string(c1), string(c2))

			val := 0
			has_c1, skip_c1 := false, false
			for _, char := range s {
				modifier := 0

				if char == c1 {
					has_c1 = true

					modifier = -1
					
					// TODO: implement ignore duplicating chars 
					if skip_c1 {
						modifier = 0
					}
				}

				if char == c2 {
					modifier = 1
				}

				val += modifier

				fmt.Printf("\tChar: %v, mod: %v, val: %v (has_c1/skip_c1: %v/%v)\n", string(char), modifier, val, has_c1, skip_c1)

				if has_c1 && variance <= val {
					variance = val
				}
			}


			fmt.Printf("Cycle variance: %v\n---------\n", variance)
		}
	}

	return variance
}