package space

import "strings"

type Planet string

func Age(ageSec float64, planet Planet) float64 {
	op := 0.00
	switch strings.ToLower(string(planet)) {
	case "mercury":
		op = 0.2408467
	case "venus":
		op = 0.61519726
	case "earth":
		op = 1.00
	case "mars":
		op = 1.8808158
	case "jupiter":
		op = 11.862615
	case "saturn":
		op = 29.447498
	case "uranus":
		op = 84.016846
	case "neptune":
		op = 164.79132
	}

	return float64((((ageSec/60)/60)/24)/365.25) / op

}
