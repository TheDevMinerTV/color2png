package main

import "golang.org/x/exp/constraints"

func clamp[T constraints.Ordered](v, min, max T) T {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}
