package haversine

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func (c Coordinates) toRadians() Coordinates {
	return Coordinates{
		Latitude:  degreesToRadian(c.Latitude),
		Longitude: degreesToRadian(c.Longitude),
	}
}

// NewCoordinates returns a Coordinates struct based on parameters passed
func NewCoordinates(latitude, longitude float64) Coordinates {
	return Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	}
}
