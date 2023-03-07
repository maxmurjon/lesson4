package main

import (
	"app/funcs/functions"
	"fmt"
)

func main() {
	var son int
	fmt.Scanln(&son)
	str := functions.IntToString(son)
	fmt.Println(functions.Text(son))
	functions.EvenDigitNum(str)
}
