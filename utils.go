package haversine

import (
	"math"
)

func degreesToRadian(d float64) float64 {
	return d * math.Pi / 180
}
