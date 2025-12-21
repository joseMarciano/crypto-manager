package rounder

import "math"

type Float interface {
	~float32 | ~float64
}

func TwoDecimalPlaces[T Float](value T) T {
	return Truncate[T](value, 2)
}

func Truncate[T Float](value T, places int) T {
	factor := math.Pow(10, float64(places))
	return T(math.Round(float64(value)*factor) / factor)
}
