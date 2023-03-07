package functions

import "strconv"

func EvenDigitNum(s string) int {
	sumj := 0
	for i := 0; i < len(s); i++ {
		if i%2 != 0 {
			j, _ := strconv.Atoi(string(s[i]))
			sumj += j
		}
	}
	return sumj
}
