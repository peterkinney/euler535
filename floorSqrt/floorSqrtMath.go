package floorSqrt

import (
	"math"
)

func FloorSqrt(xx uint64) uint64 {
    zz, oldz := float64(1.5), float64(0)
    fx := float64(xx)
    for math.Abs(zz - oldz) > 1 {
        oldz = zz
        zz = zz - (zz * zz - fx) / (2 * zz)
    }
    floor := uint64(math.Floor(zz))
    square := floor * floor
    if xx >= square && xx < square + 2 * floor + 1 {
        return floor
    }
    if xx < square && xx >= square - (2 * floor - 1) {
        return floor - 1
    }
    if xx >= square + 2 * floor + 1 && xx < square + 4 * floor + 4 {
        return floor + 1
    }
    //fmt.Printf("\n\n\nquick floor failed, reverting to math.Sqrt\nx: %d est: %d \n\n", xx, zz)
    //return uint64(math.Floor(math.Sqrt(float64(xx))))
	return uint64(1)
}

func CachedFloorSqrt() func (uint64) uint64 {
    var lastSqrt, lastSquare uint64
    return func (xx uint64) uint64 {
        if xx >= lastSquare && xx < lastSquare + 2 * lastSqrt + 1 {
            return lastSqrt
        } else if xx >= lastSquare + 2 * lastSqrt + 1 && xx < lastSquare + 4 * lastSqrt + 4 {
            lastSqrt++
            lastSquare = lastSqrt * lastSqrt
            return lastSqrt
        } else {
            lastSqrt = FloorSqrt(xx)
            lastSquare = lastSqrt * lastSqrt
            return lastSqrt
        }
    }
}

func SumFSq(nn uint64) uint64 {
    if nn == 0 {
        return 0
    }
    highestSquare := FloorSqrt(nn) - 1
    base := (highestSquare * (4 * highestSquare + 5) * (highestSquare + 1))/6
    remainder := nn - ((highestSquare + 1) * (highestSquare + 1) - 1)
    return base + remainder * (highestSquare + 1)
}

func SumFSqRange(start, length uint64) uint64 {
	if length == 0 {
		return 0
	}
	if length == 1 {
		return FloorSqrt(start)
	}
    return SumFSq(start + length - 1) - SumFSq(start - 1)
}