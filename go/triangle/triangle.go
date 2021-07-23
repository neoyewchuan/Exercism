package triangle

import "math"

type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides finds the type of triangle from the given sides measurement.
func KindFromSides(a, b, c float64) Kind {

	var k Kind = NaT // not a triangle
	// validate the given number, it must be a number, must be greater than 0,
	// and must not be a infinity number
	if _validate(a) && _validate(b) && _validate(c) {
		// all sides are equal
		if a == b && b == c {
			k = Equ // equilateral
		} else if a+b >= c && a+c >= b && b+c >= a {
			// check if this is a valid triangle
			// he sum of the lengths of any two sides
			// must be greater than or equal to the length of the third side

			// any two given sides are equal
			if a == b || a == c || b == c {
				k = Iso // isosceles
			} else {
				k = Sca // scalene
			}
		}
	}
	return k
}

func _validate(f float64) bool {
	return f > 0 && !math.IsInf(f, -1) && !math.IsInf(f, 1) && !math.IsInf(f, 0) && !math.IsNaN(f)
}
