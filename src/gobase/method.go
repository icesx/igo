package gobase

import (
	"fmt"
)

type myint int

func (i *myint) method() int {
	*i = *i * *i
	return 0
}
func use_method() {
	var i myint = 2
	i.method()
	fmt.Println("i=", i)
}
