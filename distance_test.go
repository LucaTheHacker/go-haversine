package haversine

import (
	"testing"
)

type testCoordinates struct {
	p     Coordinates
	q     Coordinates
	outKm float64
}

var tests = []testCoordinates{
	{
		NewCoordinates(22.55, 43.12),  // Rio de Janeiro, Brazil
		NewCoordinates(13.45, 100.28), // Bangkok, Thailand
		6094.544408786774,
	},
	{
		NewCoordinates(20.10, 57.30), // Port Louis, Mauritius
		NewCoordinates(0.57, 100.21), // Padang, Indonesia
		5145.525771394785,
	},
	{
		NewCoordinates(51.45, 1.15),  // Oxford, United Kingdom
		NewCoordinates(41.54, 12.27), // Vatican, City Vatican City
		1389.1793118293067,
	},
	{
		NewCoordinates(22.34, 17.05), // Windhoek, Namibia
		NewCoordinates(51.56, 4.29),  // Rotterdam, Netherlands
		3429.89310043882,
	},
	{
		NewCoordinates(63.24, 56.59), // Esperanza, Argentina
		NewCoordinates(8.50, 13.14),  // Luanda, Angola
		6996.18595539861,
	},
	{
		NewCoordinates(90.00, 0.00), // North/South Poles
		NewCoordinates(48.51, 2.21), // Paris,  France
		4613.477506482742,
	},
	{
		NewCoordinates(45.04, 7.42),  // Turin, Italy
		NewCoordinates(3.09, 101.42), // Kuala Lumpur, Malaysia
		10078.111954385415,
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range tests {
		km := Distance(input.p, input.q).Kilometers()

		if input.outKm != km {
			t.Errorf("fail: want %v %v -> %v got %v",
				input.p,
				input.q,
				input.outKm,
				km,
			)
		}
	}
}
