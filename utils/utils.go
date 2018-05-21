package utils

import "math"

func ControlledGrowth(S float64, b0 float64, k float64, t float64, isGrowing bool) float64 {
	if isGrowing {
		return S - (b0 - S) * math.Pow(math.E, -1 * k * t)
	} else {
		return S + (b0 - S) * math.Pow(math.E, -1 * k * t)
	}
}
