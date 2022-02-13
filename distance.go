package haversine

import (
	"math"
)

// Distance calcuLatitudees distance using the haversine formula
func Distance(p, q Coordinates) DistanceStruct {
	from := p.toRadians()
	to := q.toRadians()

	diffLatitude := to.Latitude - from.Latitude
	diffLongitude := to.Longitude - from.Longitude

	a := math.Pow(math.Sin(diffLatitude/2), 2) + math.Cos(from.Latitude)*math.Cos(to.Latitude)*math.Pow(math.Sin(diffLongitude/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return DistanceStruct{
		C: c,
	}
}
