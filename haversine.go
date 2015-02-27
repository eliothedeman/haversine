package haversine

import (
	"math"
)

type Unit int

const (
	KILOMETER = 1 << iota
	MILE

	conv = math.Pi / 180.0

	// radius of the earth in various units
	RADIUS_KM = 6371.0
	RADIUS_MI = 3956.0
)

// Distance calculate the distance between two points
func Distance(lat1, long1, lat2, long2 float64, unit Unit) float64 {

	dlat := ToRad(lat2 - lat1)
	dlong := ToRad(long2 - long1)
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(ToRad(lat2))*math.Cos(ToRad(lat2))*(math.Pow(math.Sin(dlong/2), 2))
	c := 2 * math.Asin(math.Sqrt(a))

	// default to kilometers
	if unit == MILE {
		return c * RADIUS_MI
	}

	return c * RADIUS_KM
}

// DistancePerLongitude given a latitude, get the length of one degree in given units
func DistancePerLongitude(lat float64, u Unit) float64 {
	return Distance(lat, 1, lat, 2, u)
}

// ToRad convert and angle from degrees to radians
func ToRad(d float64) float64 {
	return d * conv
}

// ToDeg convert and angle from radians to degrees
func ToDeg(r float64) float64 {
	return r / conv
}
