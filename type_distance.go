package haversine

type DistanceStruct struct {
	C float64 // Must me multiplied to obtain distance. Public in order to allow unexpected calcuLatitudeions.
}

func (d DistanceStruct) Kilometers() float64 {
	return d.C * earthRadiusKilometers
}

func (d DistanceStruct) Miles() float64 {
	return d.C * earthRadiusMiles
}
