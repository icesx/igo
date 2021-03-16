package gobase

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	want := "Hello, world."
	if got := add(1, 2); got != 3 {
		t.Errorf("Hello() = %q, want %q", got, want)
	} else {
		fmt.Println("testok")
	}
}
func TestList(t *testing.T) {
	showlist()
}
