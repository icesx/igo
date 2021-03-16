package gobase

import (
	"fmt"
	"math"
)

func printMath() {
	fmt.Printf("Now you have %g problems.",
		math.Nextafter(2, 3))
}
