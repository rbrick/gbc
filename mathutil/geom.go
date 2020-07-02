package mathutil

import "math"

func ToRadians(deg float64) float64 {
	return deg / 180.0 * math.Pi
}

func ToDegrees(rad float64) float64 {
	return rad * 180.0 / math.Pi
}
