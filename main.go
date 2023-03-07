package main

import (
	"fmt"
	"app/funcs/functions"
)

func main() {
	var son int
	fmt.Scanln(&son)
	str:=functions.IntToString(son)
	functions.EvenDigitNum(str)
}
