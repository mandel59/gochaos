package mandelbrot

import (
	"math/cmplx"
)

// Calc tests whether c is in the Mandelbrot set.
// return 0 if c is in the set,
// otherwise return i (> 0) where i is number of trials
// when absolute value of z exceed by 2.
func Calc(c complex128, limit int) (int, complex128) {
	z := c
	for i := 1; i <= limit ; i++ {
		if real(z * cmplx.Conj(z)) > 4 {
			return i, z
		}
		z *= z
		z += c
	}
	return 0, z
}

