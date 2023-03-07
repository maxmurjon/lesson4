package main

import (
	"app/funcs/functions"
	"fmt"
)

func main() {
	var son int
	fmt.Scanln(&son)
	functions.IntToString(son)
}
