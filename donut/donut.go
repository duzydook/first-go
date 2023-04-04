package donut

import (
	"math"
	"strings"
)

var A, B float64 = 1, 1
var Z string = ".,-~:;=!*#$@"

func Animate() string {
	var b [1760]string = [1760]string{}
	var z [1760]float64 = [1760]float64{}
	A += 0.07
	B += 0.03

	var cA, sA, cB, sB float64 = math.Cos(A), math.Sin(A), math.Cos(B), math.Sin(B)

	for k := 0; k < 1760; k++ {
		if (k % 80) == 79 {
			b[k] = "\n"
		} else {
			b[k] = " "
		}
		z[k] = 0
	}

	for j := float64(0); j < 6.28; j += 0.07 {
		// j, theta
		var ct, st float64 = math.Cos(j), math.Sin(j)

		for i := float64(0); i < 6.28; i += 0.02 {
			// i, phi
			var cp, sp float64 = math.Cos(i), math.Sin(i)
			var h float64 = ct + 2
			var D float64 = 1 / (sp*h*sA + st*cA + 5)
			var t float64 = sp*h*cA - st*sA

			var x int = int(40 + 30*D*(cp*h*cB-t*sB))
			var y int = int(12 + 15*D*(cp*h*sB+t*cB))
			var o int = x + 80*y
			var N int = int(8 *
				((st*sA-sp*ct*cA)*cB -
					sp*ct*sA -
					st*cA -
					cp*ct*sB))
			if y < 22 && y > 0 && x > 0 && x < 79 && D > z[o] {
				z[o] = D
				if N > 0 {
					b[o] = string(Z[N])
				} else {
					b[o] = string(Z[0])
				}
			}
		}
	}

	return strings.Join(b[:], "")
}
