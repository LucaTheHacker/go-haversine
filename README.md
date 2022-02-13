# go-haversine

Heavily inspired by [Umahmood's haversine](https://github.com/umahmood/haversine), go-haversine provides a nice go
interface to calculate distance between coordinates using the haversine formula.

Instead of returning two values for both Kilometers and Miles, go-haversine returns a struct that contains the
parameter "C" used to calculate the distance.  
This allows the developer to do calculation optimization and result caching, and also to use the package directly in a
function call, without using variables assignment.

## Usage

```go
package main

import (
	"log"
	
	"github.com/LucaTheHacker/go-haversine"
)

func main() {
	oxford := haversine.Coordinates{
		Latitude:  51.45,
		Longitude: 1.15,
	}
	turin := haversine.Coordinates{
		Latitude:  45.04,
		Longitude: 7.42,
	}
	
	result := haversine.Distance(oxford, turin)
	log.Printf(
		"Oxford and Turin are %.2fkm (%d meters) away.",
		result.Kilometers(), int(result.Kilometers()*1000),
	)
	
	// I'm in an hurry method
	log.Printf(
		"Oxford and Turin are %.2fkm away.",
		haversine.Distance(
			haversine.NewCoordinates(51.45, 1.15),
			haversine.NewCoordinates(45.04, 7.42),
		).Kilometers(),
	)
}

```
    