package cbmath

import (
	//#cgo LDFLAGS: -lm
	//#include <math.h>
	"C"
	"fmt"
)

// PrintSqrt - C binding
func PrintSqrt(n int) {
	fmt.Println(C.sqrt(C.double(n)))
}
