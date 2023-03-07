package functions

import (
	
	"strings"
)

func RmStrings(str string) string  {
	// k := "123456 Hello world"
	s := strings.Split(str, "Hello world");
	strins := strings.Join(s, " ");
	
     return strins
	
}

