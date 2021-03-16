package gobase

import (
	"fmt"
	"testing"
)

func swap(x, y string) (string, string) {
	return y, x
}

func TestSwap(*testing.T) {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
