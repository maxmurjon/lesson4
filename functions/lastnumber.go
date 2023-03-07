package functions

import (
	"strconv"
)

func LastNumber(n string) int{

	var last_digit, _ = strconv.Atoi(string(n))

	return last_digit % 10
}