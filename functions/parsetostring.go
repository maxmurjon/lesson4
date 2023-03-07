package functions

import "strconv"

func IntToString(a int) (result string) {
	result = strconv.Itoa(a)
	return
}
