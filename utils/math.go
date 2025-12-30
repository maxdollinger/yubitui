package utils

func Clamp(n int, min int, max int) int {
	if n < min {
		return min
	}

	if n > max {
		return max
	}

	return n
}
