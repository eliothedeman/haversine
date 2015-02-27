package haversine

import "testing"

// Expected distances are ints to help guard against rounding errors on differnt bit systems.
// I assume that any error small enough to be masked by this is acceptable.
var (
	// test map of form:
	// expected_value : [lat1,long1,lat2,long2]
	distanceTests = map[int][]float64{
		1334: []float64{50, -1, 34, 12},
		6135: []float64{-12.4, 1.2325235, 70, 89.0},
	}

	dplTests = map[int]float64{
		64: 20.0,
		2:  88.12134,
	}
)

func TestDistance(t *testing.T) {
	for k, v := range distanceTests {
		if d := Distance(v[0], v[1], v[2], v[3], MILE); int(d) != k {
			t.Errorf("distance: expected: %d got: %d from %v", k, int(d), v)
		}
	}
}

func TestDistancePerLongitude(t *testing.T) {
	for k, v := range dplTests {
		if d := DistancePerLongitude(v, MILE); int(d) != k {
			t.Errorf("distance: expected: %d got: %d from %v", k, int(d), v)
		}
	}
}

func TestToRad(t *testing.T) {
	r := ToRad(1)
	if r != 0.017453292519943295 {
		t.Fail()
	}
}

func TestToDeg(t *testing.T) {
	d := ToDeg(1)
	if d != 57.29577951308232 {
		t.Fail()
	}
}
