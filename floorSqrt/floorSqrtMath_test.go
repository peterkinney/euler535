package floorSqrt

import (
	"math"
	"testing"
)

func TestFloorSqrt(tt *testing.T) {
	for ii := uint64(0); ii < 300; ii++ {
		want := uint64(math.Floor(math.Sqrt(float64(ii))))
		got := FloorSqrt(ii)
		if want != got {
			tt.Errorf("FloorSqrt(%d) Failed: want %d, got %d",ii,want,got)
		}
	}
	for ii := uint64(1000000); ii <= 2000000; ii++ {
		want := uint64(math.Floor(math.Sqrt(float64(ii))))
		got := FloorSqrt(ii)
		if want != got {
			tt.Errorf("FloorSqrt(%d) Failed: want %d, got %d",ii,want,got)
		}
	}
}

