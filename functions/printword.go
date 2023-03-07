package functions

import "strconv"

func Text(s int) (res string) {
	res = strconv.Itoa(s)
	return res + "Hello World"
}
